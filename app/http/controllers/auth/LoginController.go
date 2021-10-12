package auth

import (
	"fmt"
	"myblog/app/http/service/auth"
	"myblog/app/models/user"
	"myblog/lib/flash"
	"myblog/lib/view"
	"net/http"
)

type LoginController struct{
}

func (*LoginController)Index(w http.ResponseWriter, r *http.Request){
	view.Render(w, view.D{}, "auth.login")
}

func (*LoginController)Login(w http.ResponseWriter, r *http.Request){
	_user := user.User{
		Email:  			r.PostFormValue("email"),
		Password: 			r.PostFormValue("password"),
	}
	errs := auth.LoginStore(_user)
	fmt.Println(errs)
	if len(errs) > 0 {
		if _,flag := errs["Status"];flag {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "登录失败，请联系管理员")
		}else{
			view.Render(w, view.D{
				"Errors": errs,
				"User": _user,
			}, "auth.login")
		}
	}else{
		flash.Success("欢迎回来！")
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func (*LoginController)Logout(w http.ResponseWriter, r *http.Request){
	auth.Logout()
	flash.Success("注销成功!")
	http.Redirect(w, r, "/", http.StatusFound)
}