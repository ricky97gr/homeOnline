package friend

import (
	"github.com/ricky97gr/homeOnline/internal/pkg/response"
	friendService "github.com/ricky97gr/homeOnline/internal/webservice/service/friend"
	"github.com/gin-gonic/gin"
)

func UserCreateFriendShip(ctx *gin.Context) {
	type info struct {
		FriendUid string `json:"friendUID"`
	}
	tmp := &info{}
	err := ctx.ShouldBindJSON(tmp)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	uid := ctx.Request.Header.Get("userID")
	err = friendService.CreateFriendShip(uid, tmp.FriendUid)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	response.Success(ctx, "successfully", 1)
}

func UserDeleteFriendShip(ctx *gin.Context) {
	type info struct {
		FriendUid string `json:"friendUID"`
	}
	tmp := &info{}
	err := ctx.ShouldBindJSON(tmp)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	uid := ctx.Request.Header.Get("userID")
	err = friendService.RemoveFriendShip(uid, tmp.FriendUid)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	response.Success(ctx, "successfully", 1)
}

func UserGetFriendFollowList(ctx *gin.Context) {
	uid := ctx.Request.Header.Get("userID")
	result, err := friendService.GetFollowFriendList(uid)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	response.Success(ctx, result, int64(len(result)))
}

func UserGetFriendFollowedList(ctx *gin.Context) {
	uid := ctx.Request.Header.Get("userID")
	result, err := friendService.GetFollowedFriendList(uid)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	response.Success(ctx, result, int64(len(result)))
}
