package articles

import (
	"myblog/app/http/service/articles"
	"myblog/lib/view"
	"net/http"
)

type ArticlesController struct{
}

func (*ArticlesController)Index(w http.ResponseWriter,r *http.Request){
	data := articles.GetAll(r)
	view.Render(w, view.D{
		"Articles": data,
	}, "articles.articles")
}