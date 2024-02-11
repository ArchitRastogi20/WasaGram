package api

import (
	"net/http"
	"wasa/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that removes a user from the banned list of another
func (rt *_router) deleteBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	bearerToken := extractBearer(r.Header.Get("Authorization"))
	pathId := ps.ByName("userID")

	// Get the value of "BannedUserID" from the query parameters
	bannedUserID := r.URL.Query().Get("BannedUserID")

	// Check the user's identity for the operation
	valid := validateRequestingUser(pathId, bearerToken)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// Users can't ban themselves so this action shouldn't be possible.
	// In order to avoid making any useless operation, terminate the execution of the function.
	if bannedUserID == bearerToken {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Remove the follower in the db via db function
	err := rt.db.UnbanUser(
		User{IdUser: pathId}.ToDatabase(),
		User{IdUser: bannedUserID}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("remove-ban: error executing delete query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Respond with 204 http status
	w.WriteHeader(http.StatusNoContent)
}
