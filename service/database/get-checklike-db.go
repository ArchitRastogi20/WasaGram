package database

import "fmt"

func (db *appdbimpl) CheckIfLikeExists(user User, photoID PhotoId) (bool, error) {
	// Check if the like exists
	var exists bool
	err := db.c.QueryRow("SELECT EXISTS (SELECT 1 FROM Like WHERE Personwholikes = ? AND ImageID = ?)", user.IdUser, photoID.IdPhoto).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking if like exists: %w", err)
	}

	return exists, nil
}
