package articles

import (
	"gorm.io/gorm/clause"
	"myblog/app/models/article"
	"myblog/core"
	"myblog/lib/helper"
	"net/http"
)

func GetAll(r *http.Request) ([]article.Article){
	db := core.DB.Model(article.Article{}).Order("created_at desc")

	var articles []article.Article
	err := db.Preload(clause.Associations).Find(&articles).Error
	helper.LogError(err)
	return articles
}