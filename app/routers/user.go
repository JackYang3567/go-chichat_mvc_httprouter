package routers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"chitchat_mvc/app/controllers/user"
	"chitchat_mvc/app/controllers/common"
	"chitchat_mvc/app/controllers/thread"
)
// New .
func New() http.Handler {
	router := httprouter.New()
	
	router.GET("/", common.Index)
	router.GET("/err", common.Err)

	router.GET("/login", user.Login)
	router.GET("/logout", user.Logout)
	router.GET("/signup", user.Signup)
	router.POST("/signup_account", user.SignupAccount)
	router.POST("/authenticate", user.Authenticate)

	router.GET("/thread/new", thread.NewThread)
	router.POST("/thread/create", thread.CreateThread)
	router.POST("/thread/post", thread.PostThread)
	router.GET("/thread/read/:id", thread.ReadThread)


	return router
}