package info

import (
	"myblog/lib/view"
	"net/http"
)

type InfoController struct{
}

func (*InfoController)Index(w http.ResponseWriter, r *http.Request){
	view.Render(w, view.D{}, "info.info")
}
