package main

import (
	"net/http"
	"time"
	_ "github.com/julienschmidt/httprouter"
	. "chitchat_mvc/app/utils"
	"chitchat_mvc/app/routers"
)



func main() {
	
	Info("===", GetCurrentDirectory(),Config.DbDriverName)
	P("ChitChat",  Version(), "started at",  Config.Address)

	// Initialize a router as usual
	handler := routers.New()

	// starting up the server
	server := &http.Server{
		Addr:            Config.Address,
		Handler:        handler,
		ReadTimeout:    time.Duration(  Config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(  Config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
	
}