package database

// Database fuction that allows a user (banner) to ban another one (banned)
func (db *appdbimpl) BanUser(banner User, banned User) error {

	_, err := db.c.Exec("INSERT INTO Ban (Banlistowner, Banned) VALUES (?, ?)", banner.IdUser, banned.IdUser)
	if err != nil {
		return err
	}

	return nil
}

// // Database fuction that removes a user (banned) from the banned list of another one (banner)
// func (db *appdbimpl) UnbanUser(banner User, banned User) error {

// 	_, err := db.c.Exec("DELETE FROM Ban WHERE Banlistowner=? AND Banned=?", banner.IdUser, banned.IdUser)

// 	if err != nil {
// 		return err
// 	}

//		return nil
//	}
//
// Database function that removes a user (banned) from the banned list of another one (banner)
func (db *appdbimpl) UnbanUser(banner User, banned User) error {
	sqlQuery := "DELETE FROM Ban WHERE Banlistowner=? AND Banned=?"

	// Print the SQL query and parameters for debugging
	// debug statement
	// fmt.Printf("Executing SQL query: %s with parameters %s, %s\n", sqlQuery, banner.IdUser, banned.IdUser)

	_, err := db.c.Exec(sqlQuery, banner.IdUser, banned.IdUser)
	if err != nil {
		return err
	}

	return nil
}

// [Util] Database fuction that checks if the requesting user was banned by another 'user'. Returns 'true' if is banned, 'false' otherwise
func (db *appdbimpl) BannedUserCheck(requestingUser User, targetUser User) (bool, error) {

	var cnt int
	err := db.c.QueryRow("SELECT COUNT(*) FROM Ban WHERE Banned = ? AND Banlistowner = ?",
		requestingUser.IdUser, targetUser.IdUser).Scan(&cnt)

	if err != nil {
		// Count always returns a row thanks to COUNT(*), so this situation should not happen
		return true, err
	}

	// If the counter is 1 then the user was banned
	if cnt == 1 {
		return true, nil
	}
	return false, nil
}
