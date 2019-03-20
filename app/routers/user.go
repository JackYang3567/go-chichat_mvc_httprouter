package routers

import (
	"net/http"
	. "chitchat_mvc/app/utils"
	"github.com/julienschmidt/httprouter"
	"chitchat_mvc/app/controllers/user"
	"chitchat_mvc/app/controllers/common"
	"chitchat_mvc/app/controllers/thread"
)
// New .
func New() http.Handler {
	router := httprouter.New()

	// handle static assets
	router.ServeFiles("/static/*filepath", http.Dir(Config.Static))

	// index
	router.GET("/", common.Index)
	// error
	router.GET("/err", common.Err)

	// defined in controllers/user/user.go
	router.GET("/login", user.Login)
	router.GET("/logout", user.Logout)
	router.GET("/signup", user.Signup)
	router.POST("/signup_account", user.SignupAccount)
	router.POST("/authenticate", user.Authenticate)

	// defined in controllers/thread/thread.go
	router.GET("/thread/new", thread.NewThread)
	router.POST("/thread/create", thread.CreateThread)
	router.POST("/thread/post", thread.PostThread)
	router.GET("/thread/read/:id", thread.ReadThread)

	

	return router
}