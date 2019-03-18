package common

import (
	. "chitchat_mvc/app/utils"
	"chitchat_mvc/app/models"	
	
	"net/http"
	"github.com/julienschmidt/httprouter"
)

// GET /err?msg=
// shows the error message page
func Err(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	vals := r.URL.Query()
	_, err :=  Session(w, r)
	if err != nil {
		 GenerateHTML(w, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		 GenerateHTML(w, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}


func Index(w http.ResponseWriter, r *http.Request ,_ httprouter.Params) {
	threads, err := models.Threads()
	if err != nil {
		 Error_message(w, r, "Cannot get threads")
	} else {
		_, err :=  Session(w, r)
		if err != nil {
			 GenerateHTML(w, threads, "layout", "public.navbar", "index")
		} else {
			 GenerateHTML(w, threads, "layout", "private.navbar", "index")
		}
	}
}
