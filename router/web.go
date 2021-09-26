package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterRouter(r *mux.Router){
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "<h1>Hello Word<h1>")
	})
}