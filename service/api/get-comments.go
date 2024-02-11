package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	// "http"
	"wasa/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getAllComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Extract userID from the Authorization header
	requestingUserID := extractBearer(r.Header.Get("Authorization"))

	// check auth
	if !validIdentifier(requestingUserID) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Retrieve the photo ID directly as an int from the path
	photoId, err := strconv.Atoi(ps.ByName("imageID"))
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-delete/Atoi: error converting photoId to int")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Call the database function to get all comments for the image
	comments, err := rt.db.GetAllComments(PhotoId{IdPhoto: int64(photoId)}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("allComments: failed to get comments from database")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"error": fmt.Errorf("failed to get comments from database: %w", err)})
		return
	}

	// Respond with the comments
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(comments)
}
