package requests

import (
	"github.com/thedevsaddam/govalidator"
	"myblog/app/models/user"
)

func info_rules() (govalidator.MapData){
	return govalidator.MapData{
		"name": 	[]string{"required","alpha_num","between:3,10", "not_exists:users,name"},
	}
}

func info_messages() (govalidator.MapData){
	return govalidator.MapData{
		"name": []string{
			"required:用户名不能为空",
			"alpha_num:用户名格式只允许数字和英文",
			"between:用户名长度为3~10之间",
		},
	}
}

func InfoVaildate(data user.User) map[string][]string{
	opts:= govalidator.Options{
		Data: &data,
		Rules: info_rules(),
		TagIdentifier: "valid",
		Messages: info_messages(),
	}
	errs := govalidator.New(opts).ValidateStruct()

	return errs
}
