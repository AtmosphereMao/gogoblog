package core

import(
"github.com/gorilla/mux"
"myblog/lib/config"
"myblog/lib/helper"
"net/http"
)

var route *mux.Router

func SetRoute(r *mux.Router){
	route = r
}

func Name2URL(routeName string, pairs ...string) string{
	url, err := route.Get(routeName).URL(pairs...)
	if err != nil{
		helper.LogError(err)
		return ""
	}
	return helper.ToString(config.Env("app.url")) + url.String()
}

func GetRouteVariable(parameter string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameter]
}
