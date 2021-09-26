package controllers

import (
	"fmt"
	"net/http"
)

type IndexController struct {
}

func (*IndexController) Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "<h1>Hello World</h1>")
}