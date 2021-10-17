package info

import (
	"errors"
	"myblog/app/http/requests"
	"myblog/app/models/user"
)

func InfoStore(_user user.User, id int) map[string][]string {
	var errs map[string][]string
	u, err := user.GetById(id)
	if err != nil{
		return map[string][]string{"messages":{errors.New("内部错误，请稍后再试。").Error()}, "Status":{"1"}}
	}
	if errs = requests.InfoVaildate(_user);len(errs) > 0{
		return errs
	}
	u.Name = _user.Name

	if err = u.Edit();err != nil {
		return map[string][]string{"messages":{errors.New("内部错误，请稍后再试。").Error()}, "Status":{"1"}}
	}
	return nil
}