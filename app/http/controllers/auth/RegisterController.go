package auth

import (
	"fmt"
	"myblog/app/http/service/auth"
	"myblog/app/models/user"
	"myblog/lib/flash"
	"myblog/lib/view"
	"net/http"
)

type RegisterController struct {
}

func (*RegisterController)Index(w http.ResponseWriter, r *http.Request){
	view.Render(w,view.D{}, "auth.register")
}

func (*RegisterController)Register(w http.ResponseWriter, r *http.Request){
	_user := user.User{
		Name: 				r.PostFormValue("name"),
		Email:  			r.PostFormValue("email"),
		Password: 			r.PostFormValue("password"),
		PasswordConfirm: 	r.PostFormValue("password_confirm"),
	}
	errs := auth.RegisterStore(_user)
	if len(errs) > 0{
		if _,flag := errs["Status"];flag {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "创建用户失败，请联系管理员")
		}else{
			view.Render(w, view.D{
				"Errors": errs,
				"User": _user,
			}, "auth.register")
		}
	}else{
		flash.Success("注册成功!")
		auth.Login(_user)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}