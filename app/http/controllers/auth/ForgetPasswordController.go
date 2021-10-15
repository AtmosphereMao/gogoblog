package auth

import (
	"fmt"
	"myblog/app/http/service/auth"
	"myblog/app/models/password_resets"
	"myblog/app/models/user"
	"myblog/core"
	"myblog/lib/flash"
	"myblog/lib/view"
	"net/http"
)

type ForgetPasswordController struct {
}

func (*ForgetPasswordController)Index(w http.ResponseWriter, r *http.Request){
	view.Render(w, view.D{},"auth.forget")
}

func (*ForgetPasswordController)Find(w http.ResponseWriter, r *http.Request){
	var _user user.User
	_user.Email = r.PostFormValue("email")
	errs := auth.RequestFindPassword(_user)
	if len(errs) > 0 {
		if _, flag := errs["Status"]; flag {
			flash.Danger(errs["messages"][0])
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			fmt.Println(errs)
			view.Render(w, view.D{
				"Errors": errs,
				"User":   _user,
			}, "auth.forget")
		}
	}
	flash.Success("已发送修改密码邮件至："+_user.Email)
	http.Redirect(w, r, core.Name2URL("auth.forget"), http.StatusFound)
}

func (*ForgetPasswordController)Reset(w http.ResponseWriter, r *http.Request){
	var _pr password_resets.PasswordResets
	_pr.Token = r.URL.Query()["token"][0]
	if err, _ := auth.ResetTokenIsExist(_pr);err != nil{
		flash.Danger(err.Error())
		http.Redirect(w, r, "/", http.StatusFound)
	}
	view.Render(w, view.D{
		"Token": _pr.Token,
	}, "auth.reset")
}

func (*ForgetPasswordController)ResetStore(w http.ResponseWriter, r *http.Request) {
	var _pr password_resets.PasswordResets
	var prEmail string
	_pr.Token = r.URL.Query()["token"][0]
	if err, email := auth.ResetTokenIsExist(_pr);err != nil{
		flash.Danger(err.Error())
		http.Redirect(w, r, "/", http.StatusFound)
	}else{
		prEmail = email
	}

	data := map[string]string{
		"Password" : r.PostFormValue("password"),
		"PasswordConfirm" : r.PostFormValue("password_confirm"),
	}

	var _user user.User
	var errs map[string][]string

	if data["Password"] != data["PasswordConfirm"] {
		errs["password_confirm"] = append(errs["password_confirm"], "两次输入的密码不同")
	}
	if data["Password"] == ""{
		errs["password"] = append(errs["password"], "密码为必填项")
	}
	if len([]rune(data["Password"])) < 6{
		errs["password"] = append(errs["password"], "密码长度需大于 6")
	}
	if len(errs) > 0{
		view.Render(w, view.D{
			"Errors": errs,
			"Token": _pr.Token,
		}, "auth.reset")
	}

	_user.Email = prEmail
	_user,_ = _user.GetByEmail()
	_user.Password = data["Password"]

	if err := _user.EditPassword();err != nil{
		flash.Danger("发生内部错误，请联系管理员")
		http.Redirect(w, r, core.Name2URL("auth.forget"), http.StatusFound)
	}
	_pr,_ = _pr.GetByToken()
	if err := _pr.DeleteByEmail();err != nil{
		flash.Danger("发生内部错误，请联系管理员")
		http.Redirect(w, r, core.Name2URL("auth.forget"), http.StatusFound)
	}
	flash.Success("密码保存成功")
	http.Redirect(w, r, core.Name2URL("auth.login"), http.StatusFound)
}
