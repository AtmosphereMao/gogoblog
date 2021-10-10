package auth

import (
	"myblog/app/models/user"
	"myblog/lib/session"
)

func Login(_user user.User){
	session.Put("uid", _user.GetStringID())
}