package model

// 记录每一个用户的每一个积分记录
type Score struct {
	UUID       string    `json:"uuid" gorm:"column:uuid"`
	UserID     string    `json:"userID" gorm:"column:userID"`
	Score      int16     `json:"score" gorm:"column:score"`
	CreateTime int64     `json:"createTime" gorm:"column:createTime"`
	Type       ScoreType `json:"type" gorm:"column:type"`
	Reason     string    `json:"reason" gorm:"column:reason"`
}
type ScoreType int8

const (
	Login ScoreType = iota + 1
	PostComment
	PostArticle
	LikeComment
	UnLikeComment
	LikeArticle
	UnLikeArticle
)

var scoreMap = map[ScoreType]int16{
	Login:         10,
	PostComment:   4,
	PostArticle:   15,
	LikeComment:   2,
	UnLikeComment: 2,
	LikeArticle:   2,
	UnLikeArticle: 2,
}

func GetScoreByReason(typ ScoreType) int16 {
	return scoreMap[typ]
}

func (Score) TableName() string {
	return "score"
}
