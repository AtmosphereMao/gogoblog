package router

import (
	"github.com/gorilla/mux"
	"myblog/app/http/controllers"
	"myblog/app/http/controllers/articles"
	"myblog/app/http/controllers/auth"
	_info "myblog/app/http/controllers/info"
	"myblog/app/http/middlewares"
	"net/http"
)

func RegisterRouter(r *mux.Router){
	// Index
	ic := new(controllers.IndexController)
	r.HandleFunc("/",	ic.Index).Methods("GET").Name("home")

	// Auth
	register := new(auth.RegisterController)
	r.HandleFunc("/register", register.Index).Methods("GET").Name("auth.register")
	r.HandleFunc("/register", register.Register).Methods("POST").Name("auth.register")
	login := new(auth.LoginController)
	r.HandleFunc("/login", login.Index).Methods("GET").Name("auth.login")
	r.HandleFunc("/login", login.Login).Methods("POST").Name("auth.login")
	r.HandleFunc("/logout", login.Logout).Methods("POST").Name("auth.logout")
	forgetPassword := new(auth.ForgetPasswordController)
	r.HandleFunc("/forget", forgetPassword.Index).Methods("GET").Name("auth.forget")
	r.HandleFunc("/forget", forgetPassword.Find).Methods("POST").Name("auth.forget")
	r.HandleFunc("/forget/reset", forgetPassword.Reset).Methods("GET").Name("auth.forget.reset")
	r.HandleFunc("/forget/reset", forgetPassword.ResetStore).Methods("POST").Name("auth.forget.reset")

	// Info
	info := new(_info.InfoController)
	r.HandleFunc("/info/{id:[0-9]+}", middlewares.Authenticate(info.Index)).Methods("GET").Name("info")
	r.HandleFunc("/info/{id:[0-9]+}", middlewares.Authenticate(info.Edit)).Methods("POST").Name("info")

	// Articles
	ac := new(articles.ArticlesController)
	r.HandleFunc("/articles", ac.Index).Methods("GET").Name("articles.index")

	// 静态资源public
	r.PathPrefix("/css/").Handler(http.FileServer(http.Dir("./public")))
	r.PathPrefix("/js/").Handler(http.FileServer(http.Dir("./public")))

	// Session服务
	r.Use(middlewares.StartSession)
}