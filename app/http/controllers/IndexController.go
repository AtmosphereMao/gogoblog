package controllers

import (
	"myblog/lib/view"
	"net/http"
)

type IndexController struct {
}

func (*IndexController) Index(w http.ResponseWriter, r *http.Request){
	view.Render(w, view.D{
	}, "index")
}