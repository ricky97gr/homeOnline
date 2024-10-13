package model

import (
	"database/sql/driver"
	"encoding/json"
)

func (t *Tags) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

func (t Tags) Value() (driver.Value, error) {
	return json.Marshal(t)
}

type Tags []string

type Article struct {
	AuthorID     string `json:"authorID" bson:"authorID" gorm:"column:authorID"`
	UserName     string `json:"userName" bson:"userName" gorm:"column:userName"`
	IsOriginal   bool   `json:"isOriginal" bson:"isOriginal" gorm:"column:isOriginal"`
	OriginalUrl  string `json:"originalUrl" bson:"originalUrl" gorm:"originalUrl"`
	OriginalUser string `json:"originalUser" bson:"originalUser" gorm:"originalUser"`
	ArticleID    string `json:"articleID" bson:"articleID" gorm:"column:articleID"`
	// 文章内容
	Context    string `json:"context" bson:"context" gorm:"column:context"`
	CreateTime int64  `json:"createTime" bson:"createTime" gorm:"column:createTime"`
	// 点赞数
	LikeCount int32 `json:"likeCount" bson:"likeCount" gorm:"column:likeCount"`
	//1: 审核通过 2: 审核中 3： 封禁 4： 草稿状态, 5 退回
	IsShow         int    `json:"isShow" bson:"isShow" gorm:"column:isShow"`
	IsShortArticle bool   `json:"isShortArticle" bson:"isShortArticle" gorm:"column:isShortArticle"`
	ViewCount      int32  `json:"viewCount" bson:"viewCount" gorm:"column:viewCount"`
	Introduction   string `json:"introduction" bson:"introduction" gorm:"column:introduction"`
	Category       string `json:"category" bson:"category" gorm:"column:category"`
	Tags           Tags   `json:"tags" bson:"tags" gorm:"column:tags"`
	Title          string `json:"title" bson:"title" gorm:"column:title"`
}

const (
	ArticleShow = iota + 1
	ArticleReviewing
	ArticleBanned
	ArticleDraft
	ArticleSendBack
)

func (Article) TableName() string {
	return "article"
}
