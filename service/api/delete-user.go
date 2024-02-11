package api

import (
	// "encoding/json"

	"net/http"

	"wasa/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// func (rt *_router) deleteuser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
// 	w.Header().Set("Content-Type", "application/json")

// 	// Extract the user ID from the URL path
// 	userID := ps.ByName("id")

// 	// Check if the extracted user ID is valid
// 	if !validIdentifier(userID) {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

//		// Delete the user from the DB
//		err := rt.db.DeleteUser(User{IdUser: userID}.ToDatabase())
//		if err != nil {
//			w.WriteHeader(http.StatusOK)
//			err = json.NewEncoder(w).Encode(User{IdUser: userID})
//			if err != nil {
//				w.WriteHeader(http.StatusInternalServerError)
//				ctx.Logger.WithError(err).Error("session: can't create response JSON")
//			}
//			return
//		}
//	}
func (rt *_router) deleteuser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	// Extract the user ID from the URL path
	userID := ps.ByName("userID") // Change from "id" to "userID"

	// Check if the extracted user ID is valid
	if !validIdentifier(userID) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// debug
	// fmt.Printf("Decoded User: %+v\n", userID)

	// Delete the user from the DB
	err := rt.db.DeleteUser(User{IdUser: userID}.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("session: failed to delete user")
		return
	}

	// Set the status code to indicate successful deletion with no content
	w.WriteHeader(http.StatusNoContent)
}
