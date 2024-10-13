package socre

import (
	"time"

	"github.com/ricky97gr/homeOnline/internal/webservice/database/mysql"
	"github.com/ricky97gr/homeOnline/internal/webservice/model"
	systemService "github.com/ricky97gr/homeOnline/internal/webservice/service/system"
	"github.com/ricky97gr/homeOnline/pkg/uuid"
)

func AddScore(userID string, reason model.ScoreType) error {
	s := &model.Score{
		UUID:       uuid.GetUUID(),
		UserID:     userID,
		Score:      model.GetScoreByReason(reason),
		CreateTime: time.Now().UnixMilli(),
		Type:       reason,
		Reason:     "",
	}
	err := createScore(s)
	if err != nil {
		return err
	}
	return addScoreToUser(userID, s.Score)
}

func createScore(s *model.Score) error {
	c := mysql.GetClient()
	return c.C.Create(s).Error
}

func addScoreToUser(userID string, target int16) error {
	score, err := systemService.GetUserScore(userID)
	if err != nil {
		return err
	}
	return systemService.UpdateUserScore(userID, score+int64(target))
}
