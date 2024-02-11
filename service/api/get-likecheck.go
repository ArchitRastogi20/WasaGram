package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"wasa/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) LikeCheck(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Extracting bearer token from Authorization header
	bearerAuth := extractBearer(r.Header.Get("Authorization"))

	// Parsing photo ID from the path
	photoID, err := strconv.Atoi(ps.ByName("imageID"))
	if err != nil {
		ctx.Logger.WithError(err).Error("checkIfLikeExists: error converting photoID to int")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Validating user identity for the operation
	valid := validateRequestingUser(ps.ByName("userID"), bearerAuth)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// Calling the database function to check if the like exists
	likeExists, err := rt.db.CheckIfLikeExists(
		User{IdUser: bearerAuth}.ToDatabase(),
		PhotoId{IdPhoto: int64(photoID)}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("checkIfLikeExists: error from the database function")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Responding with a boolean indicating whether the like exists
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]bool{"likeExists": likeExists})
}
