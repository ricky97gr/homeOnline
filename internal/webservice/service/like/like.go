package like

import "C"
import (
	"errors"

	"github.com/ricky97gr/homeOnline/internal/webservice/database/mysql"
	"github.com/ricky97gr/homeOnline/internal/webservice/model"
	commentService "github.com/ricky97gr/homeOnline/internal/webservice/service/comemnt"
)

type UILike struct {
	ID string `json:"id"`
	// 1 短评 2： 文章
	Type     int8 `json:"type"`
	Like     bool `json:"like"`
	IsCancel bool `json:"isCancel"`
}

func (u *UILike) Convert(uid string) *model.GiveLike {
	return &model.GiveLike{
		ID:       u.ID,
		UserID:   uid,
		UserName: "",
		Like:     u.Like,
		Type:     u.Type,
	}
}

func GiveLike(uid string, uilike *UILike) error {
	like := uilike.Convert(uid)
	if isRecordExist(like) {
		return errors.New("already review this item")
	}
	//获取点赞数
	if like.Type == 1 {
		count, err := commentService.GetCommentLikeCount(uilike.ID, uilike.Like)
		if err != nil {
			return err
		}
		if uilike.IsCancel {
			//更新点赞数 取消赞
			err = deleteGiveLike(like)
			if err != nil {
				return err
			}
			return commentService.UpdateCommentLike(uilike.ID, uilike.Like, count-1)
		}
		err = createGiveLike(like)
		if err != nil {
			return err
		}

		//更新点赞数
		return commentService.UpdateCommentLike(uilike.ID, uilike.Like, count+1)
	}

	//todo: 更新文章点赞数
	count, err := commentService.GetCommentLikeCount(uilike.ID, uilike.Like)
	if err != nil {
		return err
	}
	if uilike.IsCancel {
		//更新点赞数 取消赞
		err = deleteGiveLike(like)
		if err != nil {
			return err
		}
		return commentService.UpdateCommentLike(uilike.ID, uilike.Like, count-1)
	}
	err = createGiveLike(like)
	if err != nil {
		return err
	}

	//更新点赞数
	return commentService.UpdateCommentLike(uilike.ID, uilike.Like, count+1)

}

func createGiveLike(like *model.GiveLike) error {
	c := mysql.GetClient()
	return c.C.Create(like).Error
}

func deleteGiveLike(like *model.GiveLike) error {
	c := mysql.GetClient()
	return c.C.Where(map[string]interface{}{"ID": like.ID, "userID": like.UserID, "type": like.Type}).Delete(&model.GiveLike{}).Error
}

func isRecordExist(like *model.GiveLike) bool {
	c := mysql.GetClient()
	var info model.GiveLike
	result := c.C.Model(&model.GiveLike{}).Where(map[string]interface{}{"ID": like.ID, "userID": like.UserID, "type": like.Type}).First(&info)
	if result.Error != nil {
		return false
	}
	if result.RowsAffected != 0 {
		return true
	}
	return false
}
