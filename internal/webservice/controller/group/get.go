package group

import (
	"github.com/ricky97gr/homeOnline/internal/pkg/newlog"
	"github.com/ricky97gr/homeOnline/internal/pkg/response"
	client_manager "github.com/ricky97gr/homeOnline/internal/webservice/im_server/client"
	group_service "github.com/ricky97gr/homeOnline/internal/webservice/service/group"
	"github.com/ricky97gr/homeOnline/pkg/paginate"
	"github.com/gin-gonic/gin"
)

func GetAllGroupByUserUID(ctx *gin.Context) {
	uid := ctx.Request.Header.Get("userID")
	result, err := group_service.GetAllGroupByUID(uid)
	if err != nil {
		newlog.Logger.Errorf("failed to get group by uid: %+s, err: %+v\n", uid, err)
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, result, int64(len(result)))
}

func GetMemberByGroupUID(ctx *gin.Context) {
	groupID := ctx.Param("id")
	page := paginate.PageQuery{
		Page:     0,
		PageSize: 20,
	}
	result, err := group_service.GetAllMemeberByGroupUID(groupID, page)
	if err != nil {
		newlog.Logger.Errorf("failed to get group member by : %+v, err: %+v\n", groupID, err)
		response.Failed(ctx, response.ErrDB)
		return
	}
	var users []string
	for _, r := range result {
		users = append(users, r.UserID)
	}
	client_manager.SetGroupMember(groupID, users)
	response.Success(ctx, result, int64(len(result)))
}
