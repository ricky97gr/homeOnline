package user

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/ricky97gr/homeOnline/internal/pkg/newlog"
	"github.com/ricky97gr/homeOnline/internal/pkg/response"
	"github.com/ricky97gr/homeOnline/internal/pkg/sendlog"
	"github.com/ricky97gr/homeOnline/internal/webservice/model"
	comemntService "github.com/ricky97gr/homeOnline/internal/webservice/service/comemnt"
	relationService "github.com/ricky97gr/homeOnline/internal/webservice/service/relation"
	systemService "github.com/ricky97gr/homeOnline/internal/webservice/service/system"
	"github.com/ricky97gr/homeOnline/pkg/groupid"
	"github.com/ricky97gr/homeOnline/pkg/paginate"
	"github.com/ricky97gr/homeOnline/pkg/userid"
)

func AdminGetUserInfo(ctx *gin.Context) {
	type info struct {
		UserID string `json:"userID" form:"userID"`
	}

	id := &info{}
	fmt.Println(id)
	err := ctx.ShouldBindQuery(id)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}

	type result struct {
		ArticleCount int64  `json:"articleCount"`
		NickName     string `json:"nickName"`
		Score        int64  `json:"score"`
		ShortComment int64  `json:"shortComment"`
		FollowCount  int64  `json:"followCount"`
		Description  string `json:"description"`
	}

	user, err := systemService.GetUserByUserID(id.UserID)
	if err != nil {
		return
	}
	shortCount, err := comemntService.GetCommentCountByUserID(id.UserID)
	if err != nil {
		return
	}
	r := &result{
		ArticleCount: 1,
		NickName:     user.NickName,
		Score:        user.Score,
		ShortComment: shortCount,
		FollowCount:  0,
		Description:  "蒸饭机器人",
	}
	response.Success(ctx, r, 1)

}

func AdminUpdateUser(ctx *gin.Context) {
	type tmpT struct {
		UserID string `json:"userID"`
		IsShow bool   `json:"isShow"`
	}
	info := &tmpT{}
	err := ctx.ShouldBind(info)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	err = systemService.UpdateUser(info.UserID, info.IsShow)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, "", 1)
}

func AdminDeleteUser(ctx *gin.Context) {
	type tmpT struct {
		UserID string `json:"userID"`
	}
	info := &tmpT{}
	err := ctx.ShouldBind(info)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	err = systemService.DeleteUser(info.UserID)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, "", 1)
}

func AdminCreateUser(ctx *gin.Context) {
	user := &systemService.UIUser{}
	err := ctx.ShouldBindJSON(user)
	if err != nil {
		newlog.Logger.Errorf("failed bind json, err: %+v\n", err)
		response.Failed(ctx, response.ErrStruct)
		return
	}
	user.UID = userid.GetUserID()
	err = systemService.AdminCreateUser(user)
	if err != nil {
		newlog.Logger.Errorf("failed to create user: %+v, err: %+v\n", user, err)
		response.Failed(ctx, response.ErrDB)
		return
	}
	err = relationService.AddUserGroupRelation(user.UID, strconv.Itoa(groupid.SuperAdministratorGroup), model.GroupNormalMember)
	if err != nil {
		newlog.Logger.Errorf("failed to add user %+v to group: %d, err: %+v\n", user, groupid.SuperAdministratorGroup, err)
		response.Failed(ctx, response.ErrDB)
		return
	}
	err = sendlog.SendOperationLog(ctx.Request.Header.Get("userName"), "cn", sendlog.AddUser, user.NickName)
	if err != nil {
		newlog.Logger.Errorf("failed to send operation log: %+v, err: %+v\n", user, err)
	}
	response.Success(ctx, "", 1)
}

func NormalGetAllUser(ctx *gin.Context) {

	q, err := paginate.GetPageQuery(ctx)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	users, count, err := systemService.GetAllUser(q)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, users, count)
}
