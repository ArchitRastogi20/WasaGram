package database

import (
	"fmt"
)

// // Database function that retrieves the list of photos of a user (only if the requesting user is not banned by that user)
// func (db *appdbimpl) GetPhotosList(requestingUser User, targetUser User) ([]Photo, error) { // requestinUser User,

// 	rows, err := db.c.Query("SELECT * FROM Image WHERE UserID = ? ORDER BY date DESC", targetUser.IdUser)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// Wait for the function to finish before closing rows
// 	defer func() { _ = rows.Close() }()

// 	// Read all the photos in the resulset
// 	var photos []Photo
// 	for rows.Next() {
// 		var photo Photo
// 		err = rows.Scan(&photo.PhotoId, &photo.Owner, &photo.Date)
// 		if err != nil {
// 			return nil, err
// 		}

// 		comments, err := db.GetCompleteCommentsList(requestingUser, targetUser, PhotoId{IdPhoto: int64(photo.PhotoId)}) // Old: GetCommentsLen
// 		if err != nil {
// 			return nil, err
// 		}
// 		photo.Comments = comments

// 		likes, err := db.GetLikesList(requestingUser, targetUser, PhotoId{IdPhoto: int64(photo.PhotoId)}) // Old: GetLikesLen
// 		if err != nil {
// 			return nil, err
// 		}
// 		photo.Likes = likes

// 		photos = append(photos, photo)
// 	}

// 	if rows.Err() != nil {
// 		return nil, err
// 	}

// 	return photos, nil
// }

// for user profile

func (db *appdbimpl) GetPhotosUserProfile(requestingUser User) (UserProfile, error) {
	var userProfile UserProfile

	// Query the database to get user details along with follower and following counts
	err := db.c.QueryRow(`
		SELECT
			U.UserID, U.Username, U.Firstname, U.Lastname, U.Email,
			(Select COUNT(DISTINCT Follow.Followers)
			From Follow
			Where U.UserID = Follow.Followlisrowner ) AS FollowersCount,
			(Select COUNT(DISTINCT Follow.Followlisrowner)
			From Follow
			Where U.UserID = Follow.Followers) AS FollowingCount
		FROM User U
		WHERE U.UserID = ?
		GROUP BY U.UserID, U.Username, U.Firstname, U.Lastname, U.Email`, requestingUser.IdUser).
		Scan(&userProfile.UserID, &userProfile.Username, &userProfile.Firstname, &userProfile.Lastname, &userProfile.Email,
			&userProfile.FollowersCount, &userProfile.FollowingCount)
	if err != nil {
		// log.Println("Error retrieving user profile:", err)
		return UserProfile{}, fmt.Errorf("failed to retrieve user details: %w", err)
	}

	// Query the database to get user's photos with likes and comments
	rows, err := db.c.Query(`
		SELECT I.ImageID, I.Imagedata, I.Time,
			COUNT(L.LikeID) AS LikeCount,
			COUNT(C.CommentID) AS CommentCount
		FROM Image I
		LEFT JOIN Like L ON I.ImageID = L.ImageID
		LEFT JOIN Comment C ON I.ImageID = C.ImageID
		WHERE I.UserID = ?
		GROUP BY I.ImageID, I.Imagedata, I.Time
		ORDER BY I.Time DESC`, requestingUser.IdUser)
	if err != nil {
		// log.Println("Error retrieving user photos:", err)
		return UserProfile{}, fmt.Errorf("failed to retrieve user photos: %w", err)
	}
	defer rows.Close()

	// Iterate through the result set and populate the photos array
	for rows.Next() {
		var userPhoto UserPhoto
		err := rows.Scan(&userPhoto.ImageID, &userPhoto.PhotoData, &userPhoto.Timestamp, &userPhoto.LikeCount, &userPhoto.CommentCount)
		if err != nil {
			// log.Println("Error scanning user photo:", err)
			return UserProfile{}, fmt.Errorf("failed to iterate: %w", err)
		}
		userProfile.Photos = append(userProfile.Photos, userPhoto)
	}
	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		// log.Println("Error during iteration through rows:", err)
		return UserProfile{}, fmt.Errorf("error during iteration: %w", err)
	}

	return userProfile, nil
}

// Database function that removes a photo from the database
func (db *appdbimpl) DeletePhoto(owner User, p PhotoId) error {

	_, err := db.c.Exec("DELETE FROM Image WHERE UserID = ? AND ImageID = ?",
		owner.IdUser, p.IdPhoto)
	if err != nil {
		// Error during the execution of the query
		return err
	}

	return nil
}

// [Util] Database function that checks if a photo exists
func (db *appdbimpl) CheckPhotoExistence(targetPhoto PhotoId) (bool, error) {

	var rows int
	err := db.c.QueryRow("SELECT COUNT(*) FROM Image WHERE (UserID = ?)", targetPhoto.IdPhoto).Scan(&rows)
	if err != nil {
		return false, err
	}

	if rows == 0 {
		return false, nil
	}
	return true, nil

}

// Function that retrieves the user stream
func (db *appdbimpl) GetUserStreamNew(requestingUser User) (UserStream, error) {

	var userStream UserStream

	// Query the database to get images from people the user follows
	rows, err := db.c.Query(`
		SELECT I.ImageID, I.Imagedata, I.Time,
			COUNT(L.LikeID) AS LikeCount,
			COUNT(C.CommentID) AS CommentCount
		FROM Image I
		LEFT JOIN Like L ON I.ImageID = L.ImageID
		LEFT JOIN Comment C ON I.ImageID = C.ImageID
		WHERE I.UserID IN (
			SELECT Followlisrowner
			FROM Follow
			WHERE Followers = ?
		)
		GROUP BY I.ImageID, I.Imagedata, I.Time
		ORDER BY I.Time DESC`, requestingUser.IdUser)
	if err != nil {
		// log.Println("Error retrieving user stream images:", err)
		return nil, err
	}
	defer rows.Close()

	// Iterate through the result set and populate the userStream array
	for rows.Next() {
		var userPhoto UserPhoto
		err := rows.Scan(&userPhoto.ImageID, &userPhoto.PhotoData, &userPhoto.Timestamp, &userPhoto.LikeCount, &userPhoto.CommentCount)
		if err != nil {
			// log.Println("Error scanning user stream image:", err)
			return nil, err
		}
		userStream = append(userStream, userPhoto)
	}
	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		// log.Println("Error during iteration through rows:", err)
		return nil, err
	}

	return userStream, nil
}
