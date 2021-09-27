package articles

import (
	"gorm.io/gorm/clause"
	"myblog/app/models/article"
	"myblog/core"
	"net/http"
)

func GetAll(r *http.Request) ([]article.Article){
	db := core.DB.Model(article.Article{}).Order("created_at desc")

	var articles []article.Article
	result := db.Preload(clause.Associations).Find(&articles).Error
	println(result)
	return articles
}