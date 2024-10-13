package tag

import (
	"time"

	"github.com/ricky97gr/homeOnline/internal/webservice/database/mysql"
	"github.com/ricky97gr/homeOnline/internal/webservice/model"
	"github.com/ricky97gr/homeOnline/pkg/paginate"
	"github.com/ricky97gr/homeOnline/pkg/uuid"
)

type UITag struct {
	Creator string `json:"creator"`
	Name    string `json:"name"`
	IsShow  bool   `json:"isShow"`
}

func (t *UITag) Convert() *model.Tag {
	return &model.Tag{
		CreateTime: time.Now().UnixMilli(),
		Creator:    t.Creator,
		Name:       t.Name,
		IsShow:     t.IsShow,
		UUID:       uuid.GetUUID(),
	}

}

func NormalGetAllTag() ([]model.Tag, error) {
	return normalGetAllTag()
}

func AdminGetAllTag(q *paginate.PageQuery) ([]model.Tag, int64, error) {

	tags, err := adminGetAllTag(q)
	if err != nil {
		return nil, 0, err
	}
	count, err := getTagCount()
	if err != nil {
		return nil, 0, err
	}
	return tags, count, nil
}

func AddUsedCountByTagName(name string) error {
	c := mysql.GetClient()
	var tag model.Tag
	result := c.C.Model(&model.Tag{}).Where("name = ?", name).Select("usedCount").Find(&tag)
	if result.Error != nil {
		return result.Error
	}
	result = c.C.Model(&model.Tag{}).Where("name = ?", name).Update("usedCount", tag.UsedCount+1)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func AdminGetTagCount() (int64, error) {
	return getTagCount()
}

func AdminGetTagTOP(limit int) ([]model.Tag, error) {
	c := mysql.GetClient()
	var tags []model.Tag
	result := c.C.Model(&model.Tag{}).Select("name", "usedCount").Order("usedCount desc").Limit(limit).Find(&tags)
	return tags, result.Error
}

func adminGetAllTag(q *paginate.PageQuery) ([]model.Tag, error) {
	c := mysql.GetClient()
	var tags []model.Tag
	result := c.C.Model(&model.Tag{}).Order("createTime desc").Offset((q.Page - 1) * q.PageSize).Limit(q.PageSize).Find(&tags)
	return tags, result.Error
}

func getTagCount() (int64, error) {
	c := mysql.GetClient()
	var count int64
	result := c.C.Model(&model.Tag{}).Count(&count)
	return count, result.Error
}

func AdminCreateTag(tag *UITag) error {
	return createTag(tag.Convert())
}

func createTag(tag *model.Tag) error {
	c := mysql.GetClient()
	return c.C.Create(tag).Error
}

func AdminDeleteTag(uuid string) error {
	return deleteTag(uuid)
}

func deleteTag(uuid string) error {
	c := mysql.GetClient()
	return c.C.Where("uuid = ?", uuid).Delete(&model.Tag{}).Error
}

func AdminUpdateTag(uuid string, isShow bool) error {

	if isShow {
		return updateTagShow(uuid)
	}
	return updateTagNotShow(uuid)
}

func updateTagShow(uuid string) error {
	c := mysql.GetClient()
	return c.C.Model(&model.Tag{}).Where("uuid = ?", uuid).Update("isShow", true).Error
}

func updateTagNotShow(uuid string) error {
	c := mysql.GetClient()
	return c.C.Model(&model.Tag{}).Where("uuid = ?", uuid).Update("isShow", false).Error
}

func normalGetAllTag() ([]model.Tag, error) {
	c := mysql.GetClient()
	var tags []model.Tag
	result := c.C.Model(&model.Tag{}).Where("isShow = true").Order("createTime desc").Find(&tags)
	return tags, result.Error
}
