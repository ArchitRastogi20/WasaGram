package database

import (
	"strings"
)

// Database function that gets a user's nickname
func (db *appdbimpl) GetNickname(user User) (string, error) {

	var nickname string

	err := db.c.QueryRow(`SELECT nickname FROM User WHERE UserID = ?`, user.IdUser).Scan(&nickname)
	if err != nil {
		// Error during the execution of the query
		return "", err
	}
	return nickname, nil
}

// Database function that updates an existing user in the User table
func (db *appdbimpl) UpdateUser(user User, updatedUser UserPropertiesUpdate) error {
	// Check if there are any fields to update
	// if !hasFieldsToUpdate(updatedUser) {
	// 	return errors.New("no fields provided for update")
	// }

	// Generate the SQL query based on the provided fields
	query := "UPDATE User SET "
	var params []interface{}

	if updatedUser.Username != "" {
		query += "Username = ?, "
		params = append(params, updatedUser.Username)
	}
	if updatedUser.Firstname != "" {
		query += "Firstname = ?, "
		params = append(params, updatedUser.Firstname)
	}
	if updatedUser.Lastname != "" {
		query += "Lastname = ?, "
		params = append(params, updatedUser.Lastname)
	}
	if updatedUser.Email != "" {
		query += "Email = ?, "
		params = append(params, updatedUser.Email)
	}

	// Trim the trailing comma and add the WHERE clause
	query = strings.TrimSuffix(query, ", ") + " WHERE UserID = ?"

	// Add the UserID to the parameters
	params = append(params, user.IdUser)

	// Execute the SQL query
	_, err := db.c.Exec(query, params...)
	if err != nil {
		// Error during the execution of the query
		return err
	}

	return nil
}
