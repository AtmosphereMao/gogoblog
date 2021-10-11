package auth

import (
	"myblog/app/models/user"
	"myblog/lib/session"
)

func _getUID() string{
	_uid := session.Get("uid")
	uid, ok := _uid.(string)
	if ok && len(uid) > 0{
		return uid
	}
	return ""
}

func Login(_user user.User){
	session.Put("uid", _user.GetStringID())
}

func Check() bool{
	return len(_getUID()) > 0
}

func Logout(){
	session.Forget("uid")
}