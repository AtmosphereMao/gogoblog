package session

import (
	"github.com/gorilla/sessions"
	"myblog/lib/config"
	"myblog/lib/helper"
	"net/http"
)

var Store = sessions.NewCookieStore([]byte(helper.ToString(config.Env("app.key"))))

var Session *sessions.Session

var Request *http.Request

var Response http.ResponseWriter

func StartSession(w http.ResponseWriter, r *http.Request){
	var err error
	Session, err = Store.Get(r, helper.ToString(config.Env("session_name")))
	helper.LogError(err)

	Response = w
	Request = r
}

func Put(key string, value interface{}){
	Session.Values[key] = value
	Save()
}

func Get(key string) interface{}{
	return Session.Values[key]
}

func Forget(key string){
	delete(Session.Values, key)
	Save()
}

func Flush(){
	Session.Options.MaxAge = -1
	Save()
}

func Save(){
	err := Session.Save(Request, Response)
	helper.LogError(err)
}