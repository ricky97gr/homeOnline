package category

import (
	"time"

	"github.com/ricky97gr/homeOnline/internal/webservice/database/mysql"
	"github.com/ricky97gr/homeOnline/internal/webservice/model"
	"github.com/ricky97gr/homeOnline/pkg/paginate"
	"github.com/ricky97gr/homeOnline/pkg/uuid"
)

type UICategory struct {
	Creator string `json:"creator"`
	Name    string `json:"name"`
	IsShow  bool   `json:"isShow"`
}

func (t *UICategory) Convert() *model.Category {
	return &model.Category{
		CreateTime: time.Now().UnixMilli(),
		Creator:    t.Creator,
		Name:       t.Name,
		IsShow:     t.IsShow,
		UUID:       uuid.GetUUID(),
	}

}

func NormalGetAllCategory() ([]model.Category, error) {
	return normalGetAllCategory()
	//c := mysql.GetClient()
	//c.GetAllCategory()
}

func AdminGetAllCategory(q *paginate.PageQuery) ([]model.Category, int64, error) {
	cates, err := getAllCategory(q)
	if err != nil {
		return nil, 0, err
	}
	count, err := getCategoryCount()
	if err != nil {
		return nil, 0, err
	}
	return cates, count, nil
}

func AdminCreateCategory(Category *UICategory) error {
	return createCategory(Category.Convert())
}

func AdminGetCategoryTOP(limit int) ([]model.Category, error) {
	c := mysql.GetClient()
	var cates []model.Category
	result := c.C.Model(&model.Category{}).Select("name", "usedCount").Order("usedCount desc").Limit(limit).Find(&cates)
	return cates, result.Error
}

func AddUsedCountByCategoryName(name string) error {
	c := mysql.GetClient()
	var category model.Category
	result := c.C.Model(&model.Category{}).Where("name = ?", name).Select("usedCount").Find(&category)
	if result.Error != nil {
		return result.Error
	}
	result = c.C.Model(&model.Category{}).Where("name = ?", name).Update("usedCount", category.UsedCount+1)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func AdminGetCategoryCount() (int64, error) {
	return getCategoryCount()
}

func AdminDeleteCategory(uuid string) error {

	return deleteCategory(uuid)
}

func AdminUpdateCategory(uuid string, isShow bool) error {

	if isShow {
		return updateCategoryShow(uuid)
	}
	return updateCategoryNotShow(uuid)
}

func createCategory(cate *model.Category) error {
	c := mysql.GetClient()
	return c.C.Create(cate).Error
}

func updateCategoryShow(uuid string) error {
	c := mysql.GetClient()
	return c.C.Model(&model.Category{}).Where("uuid = ?", uuid).Update("isShow", true).Error
}

func updateCategoryNotShow(uuid string) error {
	c := mysql.GetClient()
	return c.C.Model(&model.Category{}).Where("uuid = ?", uuid).Update("isShow", false).Error
}

func getAllCategory(q *paginate.PageQuery) ([]model.Category, error) {
	c := mysql.GetClient()
	var cates []model.Category
	result := c.C.Model(&model.Category{}).Order("createTime desc").Offset((q.Page - 1) * q.PageSize).Limit(q.PageSize).Find(&cates)
	return cates, result.Error

}
func getCategoryCount() (int64, error) {
	c := mysql.GetClient()
	var count int64
	result := c.C.Model(&model.Category{}).Count(&count)
	return count, result.Error
}

func deleteCategory(uuid string) error {
	c := mysql.GetClient()
	return c.C.Where("uuid = ?", uuid).Delete(&model.Category{}).Error
}

func normalGetAllCategory() ([]model.Category, error) {
	c := mysql.GetClient()
	var cates []model.Category
	result := c.C.Model(&model.Category{}).Where("isShow = true").Order("createTime desc").Find(&cates)
	return cates, result.Error
}
