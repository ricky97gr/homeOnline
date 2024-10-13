package model

type CommunityComment struct {
	UserName  string `json:"userName" gorm:"column:userName"`
	AuthorID  string `json:"authorID" bson:"authorID" gorm:"column:authorID"`
	CommentID string `json:"commentID" bson:"commentID" gorm:"column:commentID"`
	// 评论内容
	Context    string `json:"context" bson:"context" gorm:"column:context"`
	CreateTime int64  `json:"createTime" bson:"createTime" gorm:"column:createTime"`
	// 点赞数
	LikeCount    int32  `json:"likeCount" bson:"likeCount" gorm:"column:likeCount"`
	UnLikeCount  int32  `json:"unLikeCount" bson:"unLikeCount" gorm:"column:unLikeCount"`
	ParentID     string `json:"parentID" bson:"parentID" gorm:"column:parentID"`
	IsShow       bool   `json:"isShow" bson:"isShow" gorm:"column:isShow"`
	IsFirst      bool   `json:"isFirst" bson:"isFirst" gorm:"column:isFirst"`
	TopCommentID string `json:"topCommentID" gorm:"column:topCommentID"`
	ReplayTo     string `json:"replayTo" gorm:"column:replayTo"`
	Topic        string `json:"topic" gorm:"column:topic"`
	Address      string `json:"address" gorm:"column:address"`
	ReplayToUser string `json:"replayToUser" gorm:"column:replayToUser"`
}

func (CommunityComment) TableName() string {
	return "community_comment"
}
