package article

import "C"
import (
	"github.com/ricky97gr/homeOnline/internal/webservice/service/category"
	"github.com/ricky97gr/homeOnline/internal/webservice/service/tag"
	"time"

	"github.com/ricky97gr/homeOnline/internal/webservice/database/mysql"
	"github.com/ricky97gr/homeOnline/internal/webservice/model"
	"github.com/ricky97gr/homeOnline/pkg/paginate"
	"github.com/ricky97gr/homeOnline/pkg/uuid"
)

type UIArticle struct {
	AuthorID string `json:"authorID"`
	UserName string `json:"userName"`

	ArticleID string `json:"articleID"`
	// 文章内容
	Context    string `json:"context" `
	CreateTime int64  `json:"createTime" `
	// 点赞数
	LikeCount    int32    `json:"likeCount"`
	IsOriginal   bool     `json:"isOriginal"`
	OriginalUrl  string   `json:"originalUrl"`
	OriginalUser string   `json:"originalUser"`
	IsShow       int      `json:"isShow"`
	ViewCount    int32    `json:"viewCount"`
	Introduction string   `json:"introduction"`
	Category     string   `json:"category"`
	Tags         []string `json:"tags"`
	Title        string   `json:"title"`
}

func (u *UIArticle) Convert() *model.Article {
	return &model.Article{
		AuthorID:     u.AuthorID,
		ArticleID:    uuid.GetUUID(),
		UserName:     u.UserName,
		Context:      u.Context,
		CreateTime:   time.Now().UnixMilli(),
		IsOriginal:   u.IsOriginal,
		OriginalUrl:  u.OriginalUrl,
		OriginalUser: u.OriginalUser,
		Title:        u.Title,
		Tags:         u.Tags,
		Category:     u.Category,
		Introduction: u.Introduction,
		IsShow:       u.IsShow,
	}
}

func CreateArticle(a *UIArticle) error {
	return createArticle(a.Convert())
}

func PublishArticle(articleID string) error {
	return publishArticle(articleID)
}

func BanArticle(articleID string) error {
	return banArticle(articleID)
}

func SendBackArticle(articleID string) error {
	return sendBackArticle(articleID)
}

func AdminGetArticleList(q *paginate.PageQuery) ([]model.Article, int64, error) {
	articles, err := getAllArticle(q)
	if err != nil {
		return nil, 0, err
	}
	count, err := getArticleCount()
	if err != nil {
		return nil, 0, err
	}
	return articles, count, nil
}

func NormalGetArticleList(q *paginate.PageQuery) ([]model.Article, error) {
	return normalGetAllArticle(q)
}

func GetArticleInfoByArticleID(id string) (model.Article, error) {
	return getArticleInfoByID(id)
}

func publishArticle(articleID string) error {
	c := mysql.GetClient()
	return c.C.Model(&model.Article{}).Where("articleID = ?", articleID).Update("isShow", model.ArticleShow).Error
}

func banArticle(articleID string) error {
	c := mysql.GetClient()
	return c.C.Model(&model.Article{}).Where("articleID = ?", articleID).Update("isShow", model.ArticleBanned).Error
}

func sendBackArticle(articleID string) error {
	c := mysql.GetClient()
	return c.C.Model(&model.Article{}).Where("articleID = ?", articleID).Update("isShow", model.ArticleSendBack).Error
}

func getAllArticle(q *paginate.PageQuery) ([]model.Article, error) {
	c := mysql.GetClient()
	var articles []model.Article
	result := c.C.Model(&model.Article{}).Where("isShow != ?", model.ArticleDraft).Scopes(paginate.ParseQuery(*q)).Find(&articles)
	return articles, result.Error
}

func normalGetAllArticle(q *paginate.PageQuery) ([]model.Article, error) {
	c := mysql.GetClient()
	var articles []model.Article
	result := c.C.Model(&model.Article{}).Where("isShow = ?", model.ArticleShow).Scopes(paginate.ParseQuery(*q)).Find(&articles)
	return articles, result.Error
}

func getArticleInfoByID(id string) (model.Article, error) {
	c := mysql.GetClient()
	var article model.Article
	result := c.C.Model(&model.Article{}).Where("articleID= ?", id).Find(&article)
	return article, result.Error
}

func createArticle(a *model.Article) error {
	c := mysql.GetClient()
	err := category.AddUsedCountByCategoryName(a.Category)
	if err != nil {
		return err
	}
	for _, t := range a.Tags {
		err = tag.AddUsedCountByTagName(t)
		if err != nil {
			return err
		}
	}

	return c.C.Model(&model.Article{}).Create(a).Error
}

func getArticleCount() (int64, error) {
	c := mysql.GetClient()
	var count int64
	result := c.C.Model(&model.Article{}).Count(&count)
	return count, result.Error
}

func GetArticleByAuthorID(authorID string) ([]model.Article, error) {
	c := mysql.GetClient()
	var articles []model.Article
	result := c.C.Model(&model.Article{}).Where("authorID = ?", authorID).Find(&articles)
	return articles, result.Error
}
func getArticleByCategory(category string) ([]model.Article, error) {
	c := mysql.GetClient()
	var articles []model.Article
	result := c.C.Model(&model.Article{}).Where("category= ?", category).Find(&articles)
	return articles, result.Error
}

func getArticleByTags(tag string) ([]model.Article, error) {
	return nil, nil
}
