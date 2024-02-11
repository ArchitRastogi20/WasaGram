package api

import (
	"encoding/json"
	"io"

	"net/http"
	"wasa/service/api/reqcontext"

	// "wasa/service/database"

	"github.com/julienschmidt/httprouter"
	// "github.com/sirupsen/logrus"
)

func (rt *_router) UploadImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	// Extract the user ID from the Authorization header
	userID := extractBearer(r.Header.Get("Authorization"))
	if userID == "" {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Error("upload-image: userID not provided in Authorization header")
		return
	}

	// Parse the form data to get the file
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("upload-image: failed to parse image")
		return
	}

	// Get the file from the form data
	file, header, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("upload-image: file not provided")
		return
	}
	defer file.Close()

	// Read file data
	_, err = io.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("upload-image: internal server error")
		return
	}

	// Convert UploadImageRequest to database struct
	uploadRequest := UploadImageRequest{File: header}
	dbRequest := uploadRequest.ToDatabase()

	// Call the database function to handle file upload
	err = rt.db.UploadImage(User{IdUser: userID}.ToDatabase(), dbRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("upload-image: internal server error")
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Image uploaded successfully"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("upload-image: can't create response JSON")
	}
}
