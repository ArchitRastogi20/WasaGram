package api

import (
	"net/http"
	"strconv"
	"wasa/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// // Function that removes a comment from a photo
// func (rt *_router) deleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

// 	w.Header().Set("Content-Type", "application/json")
// 	requestingUserId := extractBearer(r.Header.Get("Authorization"))

// 	// Check if the user isn't logged
// 	if isNotLogged(requestingUserId) {
// 		w.WriteHeader(http.StatusForbidden)
// 		return
// 	}

// 	// Check if the requesting user wasn't banned by the photo owner
// 	banned, err := rt.db.BannedUserCheck(
// 		User{IdUser: requestingUserId}.ToDatabase(),
// 		User{IdUser: ps.ByName("id")}.ToDatabase())
// 	if err != nil {
// 		ctx.Logger.WithError(err).Error("post-comment/db.BannedUserCheck: error executing query")
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	if banned {
// 		// User was banned by owner, can't post the comment
// 		w.WriteHeader(http.StatusForbidden)
// 		return
// 	}

// 	// Convert the photo identifier from string to int64
// 	photo_id_64, err := strconv.ParseInt(ps.ByName("imageID"), 10, 64)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		ctx.Logger.WithError(err).Error("post-comment: failed convert photo_id to int64")
// 		return
// 	}

// 	// Convert the comment identifier from string to int64
// 	comment_id_64, err := strconv.ParseInt(ps.ByName("comment_id"), 10, 64)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		ctx.Logger.WithError(err).Error("post-comment: failed convert photo_id to int64")
// 		return
// 	}

// 	// The comment of a user x is being removed by the author of the post
// 	if ps.ByName("id") == requestingUserId {

// 		err = rt.db.UncommentPhotoAuthor(
// 			PhotoId{IdPhoto: photo_id_64}.ToDatabase(),
// 			CommentId{IdComment: comment_id_64}.ToDatabase())
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			ctx.Logger.WithError(err).Error("post-comment: failed to execute query for insertion")
// 			return
// 		}

// 		w.WriteHeader(http.StatusNoContent)
// 		return
// 	}

// 	// Function call to db for comment removal (only authors can remove their comments)
// 	err = rt.db.UncommentPhoto(
// 		PhotoId{IdPhoto: photo_id_64}.ToDatabase(),
// 		User{IdUser: requestingUserId}.ToDatabase(),
// 		CommentId{IdComment: comment_id_64}.ToDatabase())
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		ctx.Logger.WithError(err).Error("post-comment: failed to execute query for insertion")
// 		return
// 	}

// 	w.WriteHeader(http.StatusNoContent)
// }

// taking commentID from query
// Function that removes a comment from a photo
func (rt *_router) deleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("Content-Type", "application/json")
	requestingUserId := extractBearer(r.Header.Get("Authorization"))

	// Check if the user isn't logged
	if isNotLogged(requestingUserId) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Convert the photo identifier from string to int64
	photoID, err := strconv.ParseInt(ps.ByName("imageID"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("delete-comment: failed to convert imageID to int64")
		return
	}

	// Convert the comment identifier from string to int64
	commentID, err := strconv.ParseInt(r.URL.Query().Get("CommentID"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("delete-comment: failed to convert CommentID to int64")
		return
	}

	// The comment of a user x is being removed by the author of the post
	if ps.ByName("id") == requestingUserId {
		err = rt.db.UncommentPhotoAuthor(
			PhotoId{IdPhoto: photoID}.ToDatabase(),
			CommentId{IdComment: commentID}.ToDatabase())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ctx.Logger.WithError(err).Error("delete-comment: failed to execute query for removal")
			return
		}

		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Function call to db for comment removal (only authors can remove their comments)
	err = rt.db.UncommentPhoto(
		PhotoId{IdPhoto: photoID}.ToDatabase(),
		User{IdUser: requestingUserId}.ToDatabase(),
		CommentId{IdComment: commentID}.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("delete-comment: failed to execute query for removal")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
