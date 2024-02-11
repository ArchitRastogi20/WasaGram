package database

// // Database function that retrieves the list of users that liked a photo
// func (db *appdbimpl) GetLikesList(requestingUser User, requestedUser User, photo PhotoId) ([]CompleteUser, error) {

// 	rows, err := db.c.Query("SELECT id_user FROM likes WHERE id_photo = ? AND id_user NOT IN (SELECT banned FROM banned_users WHERE banner = ? OR banner = ?) "+
// 		"AND id_user NOT IN (SELECT banner FROM banned_users WHERE banned = ?)",
// 		photo.IdPhoto, requestingUser.IdUser, requestedUser.IdUser, requestingUser.IdUser)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// Wait for the function to finish before closing rows.
// 	defer func() { _ = rows.Close() }()

// 	// Read all the users in the resulset (users that liked the photo that didn't ban the requesting user).
// 	var likes []CompleteUser
// 	for rows.Next() {
// 		var user CompleteUser
// 		err = rows.Scan(&user.IdUser)
// 		if err != nil {
// 			return nil, err
// 		}

// 		// Get the nickname of the user that liked the photo
// 		nickname, err := db.GetNickname(User{IdUser: user.IdUser})
// 		if err != nil {
// 			return nil, err
// 		}
// 		user.Nickname = nickname

// 		likes = append(likes, user)
// 	}

// 	if rows.Err() != nil {
// 		return nil, err
// 	}

// 	return likes, nil
// }

// Database function that adds a like of a user to a photo
func (db *appdbimpl) LikePhoto(p PhotoId, u User) error {
	// Retrieve the Personwhoimageis (user to whom the image belongs) from the Image table
	var personWhoImageIs string
	err := db.c.QueryRow("SELECT UserID FROM Image WHERE ImageID = ?", p.IdPhoto).Scan(&personWhoImageIs)
	if err != nil {
		return err
	}

	// Execute the SQL statement to insert the new like.
	_, err = db.c.Exec("INSERT INTO Like (ImageID, Personwholikes, Personwhoimageis) VALUES (?, ?, ?)", p.IdPhoto, u.IdUser, personWhoImageIs)
	if err != nil {
		return err
	}

	return nil
}

// Database function that removes a like of a user from a photo
func (db *appdbimpl) UnlikePhoto(p PhotoId, u User) error {

	_, err := db.c.Exec("DELETE FROM Like WHERE ImageID = ? AND Personwholikes = ?", p.IdPhoto, u.IdUser)
	if err != nil {
		return err
	}

	return nil
}
