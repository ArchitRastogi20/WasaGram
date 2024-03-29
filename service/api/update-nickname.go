package api

import (
	"encoding/json"
	"net/http"
	"wasa/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// Function that updates a user's nickname
func (rt *_router) UpdateUserDetails(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Extract userID from the URL path
	userID := ps.ByName("userID")

	// Check the user's identity for the operation
	valid := validateRequestingUser(userID, extractBearer(r.Header.Get("Authorization")))
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// Get the new nickname from the request body
	var userUpdate UserPropertiesUpdate
	err := json.NewDecoder(r.Body).Decode(&userUpdate)
	if err != nil {
		ctx.Logger.WithError(err).Error("update-user: error decoding json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Modify the username with the db function
	err = rt.db.UpdateUser(
		User{IdUser: userID}.ToDatabase(),
		userUpdate.ToDatabase(),
	)
	if err != nil {
		ctx.Logger.WithError(err).Error("update-user: error executing update query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Respond with 204 http status
	w.WriteHeader(http.StatusNoContent)
}
