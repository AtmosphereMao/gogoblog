package requests

import (
	"github.com/thedevsaddam/govalidator"
	"myblog/app/models/user"
)

func register_rules() (govalidator.MapData){
	return govalidator.MapData{
		"name": 	[]string{"required","alpha_num","between:3,10", "not_exists:users,name"},
		"email":	[]string{"required","email","between:4,30","not_exists:users,email"},
		"password":	[]string{"required","min:6"},
		"password_confirm": []string{"required"},
	}
}

func register_messages() (govalidator.MapData){
	return govalidator.MapData{
		"name": []string{
			"required:用户名不能为空",
			"alpha_num:用户名格式只允许数字和英文",
			"between:用户名长度为3~10之间",
		},
		"email": []string{
			"required:邮箱不能为空",
			"between:邮箱长度需在 4~30 之间",
			"email:E-mail必须为邮箱格式",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"password_confirm": []string{
			"required:确认密码为必填项",
		},
	}
}

func RegisterVaildate(data user.User) map[string][]string{
	opts:= govalidator.Options{
		Data: &data,
		Rules: register_rules(),
		TagIdentifier: "valid",
		Messages: register_messages(),
	}
	errs := govalidator.New(opts).ValidateStruct()

	if data.Password != data.PasswordConfirm {
		errs["password_confirm"] = append(errs["password_confirm"], "两次输入的密码不同")
	}

	return errs
}