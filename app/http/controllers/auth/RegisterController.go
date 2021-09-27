package auth

import (
	"myblog/lib/view"
	"net/http"
)

type RegisterController struct {
}

func (*RegisterController)Index(w http.ResponseWriter, r *http.Request){
	view.Render(w,view.D{}, "auth.register")
}
