package database

import "fmt"

// GetAllComments retrieves all comments for a given image from the database.
func (db *appdbimpl) GetAllComments(p PhotoId) ([]Comment_stream, error) {
	rows, err := db.c.Query(`
		SELECT CommentID, Commenttext, UserID
		FROM Comment
		WHERE ImageID = ?`, p.IdPhoto)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve comments for image: %w", err)
	}
	defer rows.Close()

	var comments []Comment_stream

	for rows.Next() {
		var comment Comment_stream
		err := rows.Scan(&comment.CommentID, &comment.CommentText, &comment.UserID)
		if err != nil {
			return nil, fmt.Errorf("failed to scan comment: %w", err)
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during iteration through rows: %w", err)
	}

	return comments, nil
}
