package articles

import (
	"fmt"
	"myblog/app/http/service/articles"
	"myblog/lib/view"
	"net/http"
)

type ArticlesController struct{
}

func (*ArticlesController)Index(w http.ResponseWriter,r *http.Request){
	data := articles.GetAll(r)
	fmt.Println(data[0].Title)
	view.Render(w, view.D{
		"Articles": data,
	}, "articles.articles")
}