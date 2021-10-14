package auth

import (
	"crypto/md5"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"io"
	"myblog/app/models/password_resets"
	"myblog/app/models/user"
	"myblog/lib/helper"
	"myblog/lib/sendmail"
	"time"
)

func RequestFindPassword(_user user.User)(map[string][]string){
	if _, err := _user.GetByEmail(); err != nil{
		if err == gorm.ErrRecordNotFound{
			return map[string][]string{"email":{errors.New("邮箱不存在。").Error()}}
		}else{
			return map[string][]string{"messages":{errors.New("内部错误，请稍后再试。").Error()}, "Status":{"1"}}
		}
	}
	var _pr password_resets.PasswordResets
	_pr.Email = _user.Email
	if _, err := _pr.GetByEmail(); err != nil{
		if err == gorm.ErrRecordNotFound{
			p := md5.New()
			io.WriteString(p, helper.ToString(time.Now().Unix()))
			_pr.Token = fmt.Sprintf("%x", p.Sum(nil))
			_pr.Create()
		}else{
			return map[string][]string{"messages":{errors.New("内部错误，请稍后再试。").Error()}, "Status":{"1"}}
		}
	}else{
		_pr, _ = _pr.GetByEmail()
	}
	if err := sendmail.Send(_pr); err != nil{
		return map[string][]string{"messages":{errors.New("内部错误，请稍后再试。").Error()}, "Status":{"1"}}
	}
	return nil
}