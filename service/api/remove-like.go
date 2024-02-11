package api

import (
	"net/http"
	"strconv"
	"wasa/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that removes a like from a photo
func (rt *_router) deleteLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	requestingUserId := extractBearer(r.Header.Get("Authorization"))

	// Check if the user is logged
	if isNotLogged(requestingUserId) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	photoIdInt, err := strconv.ParseInt(ps.ByName("imageID"), 10, 64)
	if err != nil {
		ctx.Logger.WithError(err).Error("remove-like/ParseInt: error converting photo_id to int64")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Insert the like in the db via db function
	err = rt.db.UnlikePhoto(
		PhotoId{IdPhoto: photoIdInt}.ToDatabase(),
		User{IdUser: requestingUserId}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("remove-like/db.UnlikePhoto: error executing insert query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Respond with 204 http status
	w.WriteHeader(http.StatusNoContent)
}
