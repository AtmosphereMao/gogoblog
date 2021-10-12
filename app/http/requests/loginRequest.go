package requests

import (
	"github.com/thedevsaddam/govalidator"
	"myblog/app/models/user"
)

func login_rules() (govalidator.MapData){
	return govalidator.MapData{
		"email":	[]string{"required","email"},
		"password":	[]string{"required"},
	}
}

func login_messages() (govalidator.MapData){
	return govalidator.MapData{
		"email": []string{
			"required:邮箱不能为空",
			"email:E-mail必须为邮箱格式",
		},
		"password": []string{
			"required:密码为必填项",
		},
	}
}

func LoginVaildate(data user.User) map[string][]string{
	opts:= govalidator.Options{
		Data: &data,
		Rules: login_rules(),
		TagIdentifier: "valid",
		Messages: login_messages(),
	}
	errs := govalidator.New(opts).ValidateStruct()

	return errs
}