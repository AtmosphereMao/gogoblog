package auth

import (
	"myblog/app/http/requests"
	"myblog/app/models/user"
)

func RegisterStore(_user user.User)(map[string][]string){
	errs := requests.RegisterVaildate(_user)

	if len(errs) >0 {
		return errs
	}

	_user.Create()
	if _user.ID > 0 {
		return nil
	} else {
		return map[string][]string{"messages":{"创建用户失败"}, "Status":{"1"}}
	}
}