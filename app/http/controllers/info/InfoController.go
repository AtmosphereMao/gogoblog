package info

import (
	"fmt"
	"myblog/app/http/service/auth"
	"myblog/app/http/service/info"
	"myblog/app/models/user"
	"myblog/core"
	"myblog/lib/flash"
	"myblog/lib/view"
	"net/http"
)

type InfoController struct{
}

func (*InfoController)Index(w http.ResponseWriter, r *http.Request){
	view.Render(w, view.D{}, "info.info")
}

func (*InfoController)Edit(w http.ResponseWriter, r *http.Request) {
	_user := user.User{
		Name: 				r.PostFormValue("name"),
	}

	if errs := info.InfoStore(_user, auth.GetId()); len(errs) > 0{
		if _,flag := errs["Status"];flag {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "修改用户信息失败，请联系管理员")
		}else{
			view.Render(w, view.D{
				"Errors": errs,
				"User": _user,
			}, "info.info")
		}
	}else{
		flash.Success("保存成功!")
		http.Redirect(w, r, core.Name2URL("info","id",auth.GetId()), http.StatusFound)
	}
}
