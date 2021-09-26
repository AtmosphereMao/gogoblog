package router

import(
	"github.com/gorilla/mux"
	"myblog/core"
)

func InitRoute() *mux.Router {
	router := mux.NewRouter()
	RegisterRouter(router)
	core.SetRoute(router)
	return router
}