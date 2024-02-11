package api

import (
	"net/http"
	"strconv"
	"wasa/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that deletes a photo (this includes comments and likes)
func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	bearerAuth := extractBearer(r.Header.Get("Authorization"))

	// Retrieve the photo ID directly as an int from the path
	photoId, err := strconv.Atoi(ps.ByName("imageID"))
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-delete/Atoi: error converting photoId to int")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Check the user's identity for the operation
	valid := validateRequestingUser(ps.ByName("userID"), bearerAuth)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// Call to the db function to remove the photo
	err = rt.db.DeletePhoto(
		User{IdUser: bearerAuth}.ToDatabase(),
		PhotoId{IdPhoto: int64(photoId)}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-delete/RemovePhoto: error coming from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Respond with 204 http status
	w.WriteHeader(http.StatusNoContent)
}
