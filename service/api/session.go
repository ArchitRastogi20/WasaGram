package api

import (
	"encoding/json"
	"io"
	"net/http"

	"wasa/service/api/reqcontext"

	// dbpkg "wasa/service/database"
	// "fmt"

	"github.com/julienschmidt/httprouter"
)

// func (rt *_router) sessionHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var nick Nickname
// 	var user User
// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		// The body was not a parseable JSON, reject it
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Println("Error decoding user:", err)
// 		return
// 	} else if !validIdentifier(user.IdUser) {
// 		// Here we checked the user identifier and we discovered that it's not valid
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Println("Invalid user identifier:", user.IdUser)
// 		return
// 	}
// 	// Debug statement
// 	fmt.Printf("Decoded User: %+v\n", user)
// 	// Create the user in the database using LoginRequest
// 	err = rt.db.CreateUser(
// 		User{IdUser: user.IdUser}.ToDatabase(),
// 		nick.ToDatabase())
// 	if err != nil {
// 		// Handle errors
// 		w.WriteHeader(http.StatusInternalServerError)
// 		ctx.Logger.WithError(err).Error("session: can't create user")
// 		return
// 	}

// }

func (rt *_router) sessionHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")
	var nick Nickname
	var user User

	// Read the request body and store it in a buffer
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Decode user information
	err = json.Unmarshal(body, &user)
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !validIdentifier(user.IdUser) {
		// Here we checked the user identifier and discovered that it's not valid
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// // Debug statement
	// fmt.Printf("Decoded User: %+v\n", user)

	// Decode nickname information from the stored buffer
	err = json.Unmarshal(body, &nick)
	if err != nil {
		// Handle the error when decoding Nickname
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// // Debug statement
	// fmt.Printf("Decoded Nickname: %+v\n", nick)

	// Create the user in the database using LoginRequest
	err = rt.db.CreateUser(
		User{IdUser: user.IdUser}.ToDatabase(),
		nick.ToDatabase())
	if err != nil {
		// Handle errors
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("session: can't create user")
		return
	}
}
