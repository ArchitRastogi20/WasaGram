package api

import (
	"net/http"
	"wasa/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that adds a user to banned list of another
func (rt *_router) putBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// pathId := ps.ByName("userID")
	BannedUserID := r.URL.Query().Get("BannedUserID") // User ID to be banned from the path
	banlistownerID := extractBearer(r.Header.Get("Authorization"))

	// Check the user's identity for the operation (only owner of the account can add a banned user to that account list)
	valid := validateRequestingUser(banlistownerID, banlistownerID)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// Check if the user is trying to ban themself
	if banlistownerID == BannedUserID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Add the new banned user in the db via db function
	err := rt.db.BanUser(
		User{IdUser: banlistownerID}.ToDatabase(),
		User{IdUser: BannedUserID}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("put-ban/db.BanUser: error executing insert query")

		// Something  didn't work internally
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Ban implies removing the follow
	err = rt.db.UnfollowUser(
		User{IdUser: banlistownerID}.ToDatabase(),
		User{IdUser: BannedUserID}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("put-ban/db.UnfollowUser1: error executing insert query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// The banned user will not follow the user anymore or else will have the banner in his home
	err = rt.db.UnfollowUser(
		User{IdUser: BannedUserID}.ToDatabase(),
		User{IdUser: banlistownerID}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("put-ban/db.UnfollowUser2: error executing insert query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Respond with 204 http status
	w.WriteHeader(http.StatusNoContent)
}
