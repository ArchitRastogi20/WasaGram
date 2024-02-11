package api

import (
	"encoding/json"
	"net/http"
	"wasa/service/api/reqcontext"
	"wasa/service/database"

	// "wasa/service/database"

	"github.com/julienschmidt/httprouter"
)

// // Function that retrieves all the necessary info of a profile
// func (rt *_router) getUserStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
// 	requestingUserId := extractBearer(r.Header.Get("Authorization"))
// 	requestedUser := ps.ByName("id") // userID

// 	userExists, err := rt.db.CheckUser(User{IdUser: requestedUser}.ToDatabase())
// 	if err != nil {
// 		ctx.Logger.WithError(err).Error("getUserStream/db.CheckUser: error executing query")
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	if !userExists {
// 		w.WriteHeader(http.StatusNoContent)
// 		return
// 	}

// 	photos, err := rt.db.GetPhotosList(User{IdUser: requestingUserId}.ToDatabase(), User{IdUser: requestedUser}.ToDatabase())
// 	if err != nil {
// 		ctx.Logger.WithError(err).Error("getUserStream/db.GetPhotosList: error executing query")
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	nickname, err := rt.db.GetNickname(User{IdUser: requestedUser}.ToDatabase())
// 	if err != nil {
// 		ctx.Logger.WithError(err).Error("getUserStream/db.GetNickname: error executing query")
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	_ = json.NewEncoder(w).Encode(CompleteProfile{
// 		Name:     requestedUser,
// 		Nickname: nickname,
// 		Posts:    photos,
// 	})
// }

// Function that retrieves all the necessary info of a profile
// Function that retrieves all the necessary info of a profile
func (rt *_router) getUserStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	requestingUserID := extractBearer(r.Header.Get("Authorization"))

	// Call the database function to get the user stream
	userStream, err := rt.db.GetUserStreamNew(database.User{IdUser: requestingUserID})
	if err != nil {
		ctx.Logger.WithError(err).Error("getUserProfile/getUserStream: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Respond with the user stream
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(userStream)
}
