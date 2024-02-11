package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	// Login route + create user
	rt.router.POST("/session", rt.wrap(rt.sessionHandler))

	// Remove user route
	rt.router.DELETE("/users/:userID", rt.wrap(rt.deleteuser))

	// Update user rouete
	rt.router.POST("/users/:userID", rt.wrap(rt.UpdateUserDetails))

	// Followers endpoint
	rt.router.POST("/users/:userID/follow", rt.wrap(rt.putFollow))
	rt.router.DELETE("/users/:userID/follow", rt.wrap(rt.deleteFollow))

	// Ban endpoint
	rt.router.POST("/users/:userID/ban", rt.wrap(rt.putBan))
	rt.router.DELETE("/users/:userID/ban", rt.wrap(rt.deleteBan))

	// Upload image endpoint
	rt.router.POST("/users/:userID/image", rt.wrap(rt.UploadImage))

	// Delete image endpoint
	rt.router.DELETE("/users/:userID/image/:imageID", rt.wrap(rt.deletePhoto))

	// Comments endpoint
	rt.router.POST("/users/:userID/image/:imageID/comment", rt.wrap(rt.postComment))
	rt.router.DELETE("/users/:userID/image/:imageID/comment", rt.wrap(rt.deleteComment))

	// Likes endpoint
	rt.router.POST("/like/users/:userID/image/:imageID", rt.wrap(rt.putLike))
	rt.router.DELETE("/like/users/:userID/image/:imageID", rt.wrap(rt.deleteLike))

	// User profile endpoint
	rt.router.GET("/users/:userID/profile", rt.wrap(rt.getUserProfileNew))

	// User profile endpoint
	rt.router.GET("/users/:userID/stream", rt.wrap(rt.getUserStream))

	// Util function for images
	rt.router.GET("/users/:userID/image/:imageID/comment/stream", rt.wrap(rt.getAllComments))

	// util function checks if like exsits or not
	rt.router.GET("/users/:userID/image/:imageID/like/stream", rt.wrap(rt.LikeCheck))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
