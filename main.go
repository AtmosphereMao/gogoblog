package main
import (
	"myblog/app/http/middlewares"
	"myblog/config"
	"myblog/databases"
	C "myblog/lib/config"
	"myblog/lib/helper"
	"myblog/router"
	"net/http"
)

func init(){
	config.Initialize()
}

func main(){	
	databases.InitDB()
	route := router.InitRoute()
	http.ListenAndServe(":" +helper.ToString(C.Env("app.port")), middlewares.RemoveTrailingSlash(route))
}