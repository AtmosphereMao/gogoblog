package router

import (
	"github.com/gorilla/mux"
	"myblog/app/http/controllers"
	"net/http"
)

func RegisterRouter(r *mux.Router){
	// Index
	ic := new(controllers.IndexController)
	r.HandleFunc("/",	ic.Index).Methods("GET").Name("home")

	// 静态资源public
	r.PathPrefix("/css/").Handler(http.FileServer(http.Dir("./public")))
	r.PathPrefix("/js/").Handler(http.FileServer(http.Dir("./public")))
}