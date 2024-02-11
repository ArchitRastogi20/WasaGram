package database

// Database function that retrieves the list of followers of a user
func (db *appdbimpl) GetFollowers(requestinUser User) ([]User, error) {

	rows, err := db.c.Query("SELECT follower FROM Follow WHERE Followlisrowner = ?", requestinUser.IdUser)
	if err != nil {
		return nil, err
	}
	// Wait for the function to finish before closing rows.
	defer func() { _ = rows.Close() }()

	// Read all the users in the resulset (users that follow the requesting user)
	var followers []User
	for rows.Next() {
		var folower User
		err = rows.Scan(&folower.IdUser)
		if err != nil {
			return nil, err
		}
		followers = append(followers, folower)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return followers, nil
}

// Database function that retrieves the list of users followed by the user
func (db *appdbimpl) GetFollowing(requestinUser User) ([]User, error) {

	rows, err := db.c.Query("SELECT Followlisrowner FROM Follow WHERE follower = ?", requestinUser.IdUser)
	if err != nil {
		return nil, err
	}
	// Wait for the function to finish before closing rows.
	defer func() { _ = rows.Close() }()

	// Read all the users in the resulset (users followed by the requesting user)
	var following []User
	for rows.Next() {
		var folowed User
		err = rows.Scan(&folowed.IdUser)
		if err != nil {
			return nil, err
		}
		following = append(following, folowed)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return following, nil
}

// Add a follower (followlistowner is the person that is being followed)
func (db *appdbimpl) FollowUser(follower User, followed User) error {

	_, err := db.c.Exec("INSERT INTO Follow (Followlisrowner, Followers) VALUES (?, ?)",
		followed.IdUser, follower.IdUser)
	if err != nil {
		return err
	}

	return nil
}

// Database function that removes a follower from a user
func (db *appdbimpl) UnfollowUser(follower User, followed User) error {

	_, err := db.c.Exec("DELETE FROM Follow WHERE Followlisrowner=? AND Followers=?",
		followed.IdUser, follower.IdUser)
	if err != nil {
		return err
	}

	return nil
}
