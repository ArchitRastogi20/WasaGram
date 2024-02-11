package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"wasa/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that adds a comment to a photo and sends a response containing the unique id of the created comment
func (rt *_router) postComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("Content-Type", "application/json")
	// photoOwnerId := ps.ByName("id") // Commented out since it's not being used
	requestingUserId := extractBearer(r.Header.Get("Authorization"))

	if isNotLogged(requestingUserId) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Removed the banned user check section

	// Copy body content (comment sent by user) into comment (struct)
	var comment Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("post-comment/Decode: failed to decode request body json")
		return
	}

	// Check if the comment has a valid length (<=200)
	if len(comment.Comment) > 200 {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("post-comment: comment longer than 200 characters")
		return
	}

	// Convert the photo identifier from string to int64
	imageID, err := strconv.ParseInt(ps.ByName("imageID"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("post-comment/ParseInt: failed to convert imageID to int64")
		return
	}

	// Function call to db for comment creation
	commentId, err := rt.db.CommentPhoto(
		PhotoId{IdPhoto: imageID}.ToDatabase(),
		User{IdUser: requestingUserId}.ToDatabase(),
		comment.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("post-comment/db.CommentPhoto: failed to execute query for insertion")
		return
	}

	w.WriteHeader(http.StatusCreated)

	// The response body will contain the unique id of the comment
	err = json.NewEncoder(w).Encode(CommentId{IdComment: commentId})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("post-comment/Encode: failed to encode response")
		return
	}
}
