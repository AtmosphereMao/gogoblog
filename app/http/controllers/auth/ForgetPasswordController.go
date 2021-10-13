package auth

import (
	"myblog/app/http/service/auth"
	"myblog/app/models/user"
	"myblog/lib/flash"
	"myblog/lib/view"
	"net/http"
)

type ForgetPasswordController struct {
}

func (*ForgetPasswordController)Index(w http.ResponseWriter, r *http.Request){
	view.Render(w, view.D{},"auth.forget")
}

func (*ForgetPasswordController)Find(w http.ResponseWriter, r *http.Request){
	var _user user.User
	_user.Email = r.PostFormValue("email")
	errs := auth.RequestFindPassword(_user)
	if len(errs) > 0 {
		if _, flag := errs["Status"]; flag {
			flash.Danger(errs["messages"][0])
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			view.Render(w, view.D{
				"Errors": errs,
				"User":   _user,
			}, "auth.forget")
		}
	}
}

func (*ForgetPasswordController)Reset(w http.ResponseWriter, r *http.Request){

}
