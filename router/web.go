package router

import (
	"github.com/gorilla/mux"
	"myblog/app/http/controllers"
)

func RegisterRouter(r *mux.Router){
	ic := new(controllers.IndexController)
	r.HandleFunc("/",	ic.Index).Methods("GET").Name("home")
}