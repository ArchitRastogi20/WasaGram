package api

import (
	"mime/multipart"
	// "time"
	"wasa/service/database"
)

// Error messages
const INTERNAL_ERROR_MSG = "internal server error"
const PNG_ERROR_MSG = "file is not a png format"
const JPG_ERROR_MSG = "file is not a jpg format"
const IMG_FORMAT_ERROR_MSG = "images must be jpeg or png"
const INVALID_JSON_ERROR_MSG = "invalid json format"
const INVALID_IDENTIFIER_ERROR_MSG = "identifier must be a string between 3 and 16 characters"

// JSON Error Structure
type JSONErrorMsg struct {
	Message string `json:"message"` // Error messages
}

// // Photo structure for the APIs
// type Photo struct {
// 	Comments []database.CompleteComment `json:"comments"` // Number of comments of a photo
// 	Likes    []database.CompleteUser    `json:"likes"`    // Number of likes of a photo
// 	Owner    string                     `json:"owner"`    // Unique id of the owner
// 	PhotoId  int                        `json:"photo_id"` // Unique id of the photo
// 	Date     time.Time                  `json:"date"`     // Date in which the photo was uploaded
// }

// User structure for the APIs
type User struct {
	IdUser string `json:"userID"` // User's unique id
}

type UserPropertiesUpdate struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

// Struct for uploading an image
type UploadImageRequest struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

// PhotoId structure for the APIs
type PhotoId struct {
	IdPhoto int64 `json:"photo_id"` // Photo unique id
}

// Nickname structure for the APIs
type Nickname struct {
	Nickname string `json:"nickname"` // Nickname of a user
}

// Comment structure for the APIs
type Comment struct {
	Comment string `json:"comment"` // Comment content
}

// CommentId structure for the APIs
type CommentId struct {
	IdComment int64 `json:"comment_id"` // Identifier of a comment
}

// CompleteComment structure for the APIs
type CompleteComment struct {
	IdComment int64  `json:"comment_id"` // Identifier of a comment
	IdPhoto   int64  `json:"photo_id"`   // Photo unique id
	IdUser    string `json:"userID"`     // User's unique id
	Nickname  string `json:"nickname"`   // Nickname of a user
	Comment   string `json:"comment"`    // Comment content
}

// // CompleteProfile structure for the APIs
// type CompleteProfile struct {
// 	Name      string           `json:"userID"`
// 	Nickname  string           `json:"nickname"`
// 	Followers []database.User  `json:"followers"`
// 	Following []database.User  `json:"following"`
// 	Posts     []database.Photo `json:"posts"`
// }

// for user profile

type UserPhoto struct {
	ImageID      int    `json:"imageID"`
	PhotoData    []byte `json:"photo_data"`
	Timestamp    string `json:"timestamp"`
	LikeCount    int    `json:"likeCount"`
	CommentCount int    `json:"commentCount"`
}

// for full user

type UserProfile struct {
	UserID         int                  `json:"userID"`
	Username       string               `json:"username"`
	Firstname      *string              `json:"firstname"`
	Lastname       *string              `json:"lastname"`
	Email          *string              `json:"email"`
	FollowersCount int                  `json:"followersCount"`
	FollowingCount int                  `json:"followingCount"`
	Photos         []database.UserPhoto `json:"photos"`
}

type UserStream []database.UserPhoto

// func (upf UserStream) ToDatabase() database.UserProfile {
// 	return database.UserStream{

// 	}
// }

type Comment_stream struct {
	CommentID   int    `json:"commentID"`
	CommentText string `json:"commentText"`
	UserID      int    `json:"userID"`
}

func (cs Comment_stream) ToDatabase() database.Comment_stream {
	return database.Comment_stream{
		CommentID:   cs.CommentID,
		CommentText: cs.CommentText,
		UserID:      cs.UserID,
	}
}

func (upf UserProfile) ToDatabase() database.UserProfile {
	return database.UserProfile{
		UserID:         upf.UserID,
		Username:       upf.Username,
		Firstname:      upf.Firstname,
		Lastname:       upf.Lastname,
		Email:          upf.Email,
		FollowersCount: upf.FollowersCount,
		FollowingCount: upf.FollowingCount,
		Photos:         upf.Photos,
	}
}

// Converts a User from the api package to a User of the database package
func (up UserPhoto) ToDatabase() database.UserPhoto {
	return database.UserPhoto{
		ImageID:      up.ImageID,
		PhotoData:    up.PhotoData,
		Timestamp:    up.Timestamp,
		LikeCount:    up.LikeCount,
		CommentCount: up.CommentCount,
	}
}

// Converts a User from the api package to a User of the database package
func (u User) ToDatabase() database.User {
	return database.User{
		IdUser: u.IdUser,
	}
}

// ToDatabase converts a UserPropertiesUpdate from the API package to a UserPropertiesUpdate of the database package
func (upu UserPropertiesUpdate) ToDatabase() database.UserPropertiesUpdate {
	return database.UserPropertiesUpdate{
		Username:  upu.Username,
		Firstname: upu.Firstname,
		Lastname:  upu.Lastname,
		Email:     upu.Email,
	}
}

// ToDatabase converts an UploadImageRequest from the API package to a database.UploadImageRequest
func (uir UploadImageRequest) ToDatabase() database.UploadImageRequest {
	return database.UploadImageRequest{
		File: uir.File,
	}
}

// // Converts a Photo from the api package to a Photo of the database package
// func (p Photo) ToDatabase() database.Photo {
// 	return database.Photo{
// 		Comments: p.Comments,
// 		Likes:    p.Likes,
// 		Owner:    p.Owner,
// 		PhotoId:  p.PhotoId,
// 		Date:     p.Date,
// 	}
// }

// Converts a PhotoId from the api package to a PhotoId of the database package
func (p PhotoId) ToDatabase() database.PhotoId {
	return database.PhotoId{
		IdPhoto: p.IdPhoto,
	}
}

// Converts a PhotoId from the api package to a PhotoId of the database package
func (n Nickname) ToDatabase() database.Nickname {
	return database.Nickname{
		Nickname: n.Nickname,
	}
}

// Converts a Comment from the api package to a Comment of the database package
func (c Comment) ToDatabase() database.Comment {
	return database.Comment{
		Comment: c.Comment,
	}
}

// Converts a CommentId from the api package to a CommentId of the database package
func (c CommentId) ToDatabase() database.CommentId {
	return database.CommentId{
		IdComment: c.IdComment,
	}
}

// Converts a CompleteComment from the api package to a CompleteComment of the database package
func (cc CompleteComment) ToDatabase() database.CompleteComment {
	return database.CompleteComment{
		IdComment: cc.IdComment,
		IdPhoto:   cc.IdPhoto,
		IdUser:    cc.IdUser,
		Comment:   cc.Comment,
	}
}
