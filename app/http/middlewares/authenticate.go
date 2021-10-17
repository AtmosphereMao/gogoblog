package middlewares

import (
	"myblog/app/http/service/auth"
	"myblog/core"
	"myblog/lib/flash"
	"net/http"
)

func Authenticate(next http.HandlerFunc)http.HandlerFunc{
	return http.HandlerFunc(
		func (w http.ResponseWriter, r *http.Request){
			if !auth.Check(){
				flash.Warning("用户未登录")
				http.Redirect(w, r, core.Name2URL("auth.login"), http.StatusFound)
				return
			}
			next(w, r)
		})
}