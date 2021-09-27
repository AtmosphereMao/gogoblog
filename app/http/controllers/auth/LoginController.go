package auth

import (
	"myblog/lib/view"
	"net/http"
)

type LoginController struct{
}

func (*LoginController)Index(w http.ResponseWriter, r *http.Request){
	view.Render(w, view.D{}, "auth.login")
}