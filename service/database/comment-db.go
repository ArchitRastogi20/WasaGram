package database

// // Database function that retrieves the list of comments of a photo (minus the comments from users that banned the requesting user)
// func (db *appdbimpl) GetCompleteCommentsList(requestingUser User, requestedUser User, photo PhotoId) ([]CompleteComment, error) {

// 	rows, err := db.c.Query("SELECT * FROM comments WHERE id_photo = ? AND id_user NOT IN (SELECT banned FROM banned_users WHERE banner = ? OR banner = ?) "+
// 		"AND id_user NOT IN (SELECT banner FROM banned_users WHERE banned = ?)",
// 		photo.IdPhoto, requestingUser.IdUser, requestedUser.IdUser, requestingUser.IdUser)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Wait for the function to finish before closing rows
// 	defer func() { _ = rows.Close() }()

// 	// Read all the comments in the resulset (comments of the photo with authors that didn't ban the requesting user).
// 	var comments []CompleteComment
// 	for rows.Next() {
// 		var comment CompleteComment
// 		err = rows.Scan(&comment.IdComment, &comment.IdPhoto, &comment.IdUser, &comment.Comment)
// 		if err != nil {
// 			return nil, err
// 		}

// 		// Get the nickname of the user that commented
// 		nickname, err := db.GetNickname(User{IdUser: comment.IdUser})
// 		if err != nil {
// 			return nil, err
// 		}
// 		comment.Nickname = nickname

// 		comments = append(comments, comment)
// 	}

// 	if rows.Err() != nil {
// 		return nil, err
// 	}

// 	return comments, nil
// }

// Database function that adds a comment of a user to a photo
func (db *appdbimpl) CommentPhoto(p PhotoId, u User, c Comment) (int64, error) {

	res, err := db.c.Exec("INSERT INTO Comment (ImageID, Commenttext, UserID) VALUES (?, ?, ?)",
		p.IdPhoto, c.Comment, u.IdUser)
	if err != nil {
		// Error executing query
		return -1, err
	}

	commentId, err := res.LastInsertId()
	if err != nil {
		// Error getting id returned by last db operation (commentId)
		return -1, err
	}

	return commentId, nil
}

// Database function that removes a comment of a user from a photo
func (db *appdbimpl) UncommentPhoto(p PhotoId, u User, c CommentId) error {

	_, err := db.c.Exec("DELETE FROM Comment WHERE (UserID = ? AND ImageID = ? AND CommentID = ?)",
		u.IdUser, p.IdPhoto, c.IdComment)
	if err != nil {
		return err
	}

	return nil
}

// Database function that removes a comment of a user from a photo ( by post author)
func (db *appdbimpl) UncommentPhotoAuthor(p PhotoId, c CommentId) error {

	_, err := db.c.Exec("DELETE FROM Comment WHERE (ImageID = ? AND CommentID = ?)",
		p.IdPhoto, c.IdComment)
	if err != nil {
		return err
	}

	return nil
}
