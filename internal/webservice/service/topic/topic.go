package topic

import (
	"time"

	"github.com/ricky97gr/homeOnline/internal/webservice/database/mysql"
	"github.com/ricky97gr/homeOnline/internal/webservice/model"
	"github.com/ricky97gr/homeOnline/pkg/paginate"
	"github.com/ricky97gr/homeOnline/pkg/uuid"
)

type UITopic struct {
	Creator string `json:"creator"`
	Name    string `json:"name"`
	IsShow  bool   `json:"isShow"`
}

func (t *UITopic) Convert() *model.Topic {
	return &model.Topic{
		CreateTime: time.Now().UnixMilli(),
		Creator:    t.Creator,
		Name:       t.Name,
		IsShow:     t.IsShow,
		UUID:       uuid.GetUUID(),
	}

}

func NormalGetAllTopic() ([]model.Topic, error) {
	return normalGetAllTopic()
}

func AdminGetAllTopic(q *paginate.PageQuery) ([]model.Topic, int64, error) {

	topics, err := adminGetAllTopic(q)
	if err != nil {
		return nil, 0, err
	}
	count, err := getTopicCount()
	if err != nil {
		return nil, 0, err
	}
	return topics, count, nil
}
func AdminGetTopicCount() (int64, error) {
	return getTopicCount()
}

func AdminGetTopicTOP(limit int) ([]model.Topic, error) {
	c := mysql.GetClient()
	var topics []model.Topic
	result := c.C.Model(&model.Topic{}).Select("name", "usedCount").Order("usedCount desc").Limit(limit).Find(&topics)
	return topics, result.Error
}

func adminGetAllTopic(q *paginate.PageQuery) ([]model.Topic, error) {
	c := mysql.GetClient()
	var topics []model.Topic
	result := c.C.Model(&model.Topic{}).Order("createTime desc").Offset((q.Page - 1) * q.PageSize).Limit(q.PageSize).Find(&topics)
	return topics, result.Error
}

func getTopicCount() (int64, error) {
	c := mysql.GetClient()
	var count int64
	result := c.C.Model(&model.Topic{}).Count(&count)
	return count, result.Error
}

func AddUsedCountByTopicName(name string) error {
	c := mysql.GetClient()
	var topic model.Topic
	result := c.C.Model(&model.Topic{}).Where("name = ?", name).Select("usedCount").Find(&topic)
	if result.Error != nil {
		return result.Error
	}
	result = c.C.Model(&model.Topic{}).Where("name = ?", name).Update("usedCount", topic.UsedCount+1)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func AdminCreateTopic(topic *UITopic) error {
	return createTopic(topic.Convert())
}

func createTopic(topic *model.Topic) error {
	c := mysql.GetClient()
	return c.C.Create(topic).Error
}

func AdminDeleteTopic(uuid string) error {
	return deleteTopic(uuid)
}

func deleteTopic(uuid string) error {
	c := mysql.GetClient()
	return c.C.Where("uuid = ?", uuid).Delete(&model.Topic{}).Error
}

func AdminUpdateTopic(uuid string, isShow bool) error {

	if isShow {
		return updateTopicShow(uuid)
	}
	return updateTopicNotShow(uuid)
}

func updateTopicShow(uuid string) error {
	c := mysql.GetClient()
	return c.C.Model(&model.Topic{}).Where("uuid = ?", uuid).Update("isShow", true).Error
}

func updateTopicNotShow(uuid string) error {
	c := mysql.GetClient()
	return c.C.Model(&model.Topic{}).Where("uuid = ?", uuid).Update("isShow", false).Error
}

func normalGetAllTopic() ([]model.Topic, error) {
	c := mysql.GetClient()
	var topics []model.Topic
	result := c.C.Model(&model.Topic{}).Where("isShow = true").Order("createTime desc").Find(&topics)
	return topics, result.Error
}
