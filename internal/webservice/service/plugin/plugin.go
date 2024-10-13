package plugin

import "C"
import (
	"github.com/ricky97gr/homeOnline/internal/webservice/database/mysql"
	"github.com/ricky97gr/homeOnline/internal/webservice/model"
	"github.com/ricky97gr/homeOnline/pkg/paginate"
)

func ListAllPlugin(q paginate.PageQuery) ([]model.Plugin, error) {
	c := mysql.GetClient()
	var plugins []model.Plugin
	result := c.C.Model(&model.Plugin{}).Scopes(paginate.ParseQuery(q)).Find(&plugins)
	return plugins, result.Error
}

func ListAllPlugins() []model.Plugin {
	var plugins []model.Plugin
	c := mysql.GetClient()
	c.C.Model(&model.Plugin{}).Find(&plugins)
	return plugins
}

func UpdatePluginStatusByName(status int, name string) error {
	c := mysql.GetClient()
	return c.C.Model(&model.Plugin{}).Where("name = ?", name).Update("status", status).Error
}

func DeletePluginByName(name string) error {
	c := mysql.GetClient()
	return c.C.Where("name = ?", name).Delete(&model.Plugin{}).Error
}

func isPluginExist(name string) bool {
	c := mysql.GetClient()

	result := c.C.Model(&model.Plugin{}).Where("name = ?", name).Find(&model.Plugin{})
	if result.Error != nil {
		return true
	}
	return result.RowsAffected != 0
}

func CreatePlugin(p model.Plugin) error {
	c := mysql.GetClient()
	if isPluginExist(p.Name) {
		return nil
	}
	return c.C.Model(&model.Plugin{}).Create(&p).Error
}
