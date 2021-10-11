package auth

import (
	"myblog/app/http/service/auth"
	"myblog/lib/flash"
	"myblog/lib/view"
	"net/http"
)

type LoginController struct{
}

func (*LoginController)Index(w http.ResponseWriter, r *http.Request){
	view.Render(w, view.D{}, "auth.login")
}
func (*LoginController)Logout(w http.ResponseWriter, r *http.Request){
	auth.Logout()
	flash.Success("注销成功!")
	http.Redirect(w, r, "/", http.StatusFound)
}