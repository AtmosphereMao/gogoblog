package auth

import (
	"encoding/base64"
	"errors"
	scrypt "github.com/elithrar/simple-scrypt"
	"gorm.io/gorm"
	"myblog/app/http/requests"
	"myblog/app/models/user"
	"myblog/lib/helper"
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

func GetId() int{
	return helper.ToInt(session.Get("uid"))
}

func Logout(){
	session.Forget("uid")
}

func LoginStore(_user user.User)(map[string][]string) {
	errs := requests.LoginVaildate(_user)

	if len(errs) >0 {
		return errs
	}
	var err error
	var u user.User
	if u, err = _user.GetByEmail(); err != nil{
		if err == gorm.ErrRecordNotFound{
			return map[string][]string{"password":{errors.New("账号不存在或者密码不正确。").Error()}}
		}else{
			err = errors.New("内部错误，请稍后再试。")
		}
		return map[string][]string{"messages":{err.Error()}, "Status":{"1"}}
	}
	hash,_ := base64.StdEncoding.DecodeString(u.Password)
	if err = scrypt.CompareHashAndPassword(hash, []byte(_user.Password));err != nil{
		return map[string][]string{"password":{errors.New("账号不存在或者密码不正确。").Error()}}
	}

	Login(u)
	return nil
}