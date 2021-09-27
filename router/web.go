package router

import (
	"github.com/gorilla/mux"
	"myblog/app/http/controllers"
	"myblog/app/http/controllers/articles"
	"myblog/app/http/controllers/auth"
	"net/http"
)

func RegisterRouter(r *mux.Router){
	// Index
	ic := new(controllers.IndexController)
	r.HandleFunc("/",	ic.Index).Methods("GET").Name("home")

	// Auth
	register := new(auth.RegisterController)
	r.HandleFunc("/register", register.Index).Methods("GET").Name("auth.register")
	login := new(auth.LoginController)
	r.HandleFunc("/login", login.Index).Methods("GET").Name("auth.login")


	// Articles
	ac := new(articles.ArticlesController)
	r.HandleFunc("/articles", ac.Index).Methods("GET").Name("articles.index")

	// 静态资源public
	r.PathPrefix("/css/").Handler(http.FileServer(http.Dir("./public")))
	r.PathPrefix("/js/").Handler(http.FileServer(http.Dir("./public")))

}