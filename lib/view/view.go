package view

import (
	"html/template"
	"io"
	"myblog/app/http/service/auth"
	"myblog/core"
	"myblog/lib/flash"
	"myblog/lib/helper"
	"path/filepath"
	"strings"
)

type D map[string]interface{}

func Render(w io.Writer, data D, templateFiles ...string){
	RenderTemplate(w, "app", data, templateFiles...) // 链接app模板
}

func RenderSimple(w io.Writer, data D, templateFiles ...string){
	RenderTemplate(w, "simple", data, templateFiles...) // 链接simple模板
}

func RenderTemplate(w io.Writer, name string, data D, templateFiles ...string){
	var err error

	data["isLogin"] = auth.Check()
	data["flash"] = flash.All();

	allFiles := getTemplateFiles(templateFiles...)
	tmpl, err := template.New("").Funcs(template.FuncMap{
		"RouteName2URL": core.Name2URL,
	}).ParseFiles(allFiles...)
	helper.LogError(err)
	err = tmpl.ExecuteTemplate(w, name, data)
	helper.LogError(err)
}

func getTemplateFiles(tplFiles ...string) []string{
	viewDir := "resources/views/"
	for i,f := range tplFiles{
		tplFiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".gohtml"
	}
	layoutFiles, err := filepath.Glob((viewDir + "layouts/*.gohtml"))
	helper.LogError(err)
	return append(layoutFiles, tplFiles...)
}