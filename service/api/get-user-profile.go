package api

import (
	"encoding/json"
	"net/http"
	"wasa/service/api/reqcontext"

	// "wasa/service/database"

	"github.com/julienschmidt/httprouter"
)

// Function that retrieves all the necessary info of a profile
func (rt *_router) getUserProfileNew(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	requestingUserId := extractBearer(r.Header.Get("Authorization"))
	// requestedUser := ps.ByName("id")

	// Retrieve user profile information using the new GetPhotosUserProfile function
	userProfile, err := rt.db.GetPhotosUserProfile(User{IdUser: requestingUserId}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("getUserProfile/db.GetPhotosUserProfile: error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(userProfile)
}

// // Function that retrieves all the necessary info of a profile
// func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
// 	requestingUserId := extractBearer(r.Header.Get("Authorization"))
// 	requestedUser := ps.ByName("id")

// 	// Check if the requesting user is banned by the requested profile owner
// 	userBanned, err := rt.db.BannedUserCheck(User{IdUser: requestingUserId}.ToDatabase(),
// 		User{IdUser: requestedUser}.ToDatabase())
// 	if err != nil {
// 		ctx.Logger.WithError(err).Error("getUserProfile/db.BannedUserCheck/userBanned: error executing query")
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	if userBanned {
// 		w.WriteHeader(http.StatusForbidden)
// 		return
// 	}

// 	// Check if the requested profile was banned by the requesting user. If true, respond with partial content
// 	requestedProfileBanned, err := rt.db.BannedUserCheck(User{IdUser: requestedUser}.ToDatabase(),
// 		User{IdUser: requestingUserId}.ToDatabase())
// 	if err != nil {
// 		ctx.Logger.WithError(err).Error("getUserProfile/db.BannedUserCheck/requestedProfileBanned: error executing query")
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	if requestedProfileBanned {
// 		w.WriteHeader(http.StatusPartialContent)
// 		return
// 	}

// 	userExists, err := rt.db.CheckUser(User{IdUser: requestedUser}.ToDatabase())
// 	if err != nil {
// 		ctx.Logger.WithError(err).Error("getUserProfile/db.CheckUser: error executing query")
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	if !userExists {
// 		w.WriteHeader(http.StatusNoContent)
// 		return
// 	}

// 	followers, err := rt.db.GetFollowers(User{IdUser: requestedUser}.ToDatabase())
// 	if err != nil {
// 		ctx.Logger.WithError(err).Error("getUserProfile/db.GetFollowers: error executing query")
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	following, err := rt.db.GetFollowing(User{IdUser: requestedUser}.ToDatabase())
// 	if err != nil {
// 		ctx.Logger.WithError(err).Error("getUserProfile/db.GetFollowing: error executing query")
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	// here we will give the userID
// 	photos, err := rt.db.GetPhotosList(User{IdUser: requestingUserId}.ToDatabase(), User{IdUser: requestedUser}.ToDatabase())
// 	if err != nil {
// 		ctx.Logger.WithError(err).Error("getUserProfile/db.GetPhotosList: error executing query")
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	nickname, err := rt.db.GetNickname(User{IdUser: requestedUser}.ToDatabase())
// 	if err != nil {
// 		ctx.Logger.WithError(err).Error("getUserProfile/db.GetNickname: error executing query")
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	_ = json.NewEncoder(w).Encode(CompleteProfile{
// 		Name:      requestedUser,
// 		Nickname:  nickname,
// 		Followers: followers,
// 		Following: following,
// 		Posts:     photos,
// 	})
// }
