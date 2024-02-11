package database

import (
	"fmt"
	"io"
	"time"
)

// // Database function that allows a user (banner) to ban another one (banned)
// func (db *appdbimpl) UploadImage(user User, data UploadImageRequest) error {

// 	// Execute the SQL statement to insert the new image directly
// 	_, err := db.c.Exec(`
// 		INSERT INTO Image (UserID, Time, Imagedata)
// 		VALUES (?, ?, ?)
// 	`, user, time.Now(), data)

// 	if err != nil {
// 		return fmt.Errorf("error inserting image into the database: %w", err)
// 	}

//		// fmt.Println("Image uploaded successfully.")
//		// fmt.Println(&data)
//		return nil
//	}
func (db *appdbimpl) UploadImage(user User, data UploadImageRequest) error {
	// Get the user ID from the User struct
	userID := user.IdUser

	// Open the uploaded file
	file, err := data.File.Open()
	if err != nil {
		return fmt.Errorf("error opening uploaded file: %w", err)
	}
	defer file.Close()

	// Read the file content
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading file content: %w", err)
	}

	// Execute the SQL statement to insert the new image
	_, err = db.c.Exec(`
		INSERT INTO Image (UserID, Time, Imagedata) 
		VALUES (?, ?, ?)
	`, userID, time.Now(), fileContent)

	if err != nil {
		return fmt.Errorf("error inserting image into the database: %w", err)
	}

	return nil
}
