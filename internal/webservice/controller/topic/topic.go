package topic

import (
	"github.com/gin-gonic/gin"

	"github.com/ricky97gr/homeOnline/internal/pkg/newlog"
	"github.com/ricky97gr/homeOnline/internal/pkg/response"
	"github.com/ricky97gr/homeOnline/internal/pkg/sendlog"
	"github.com/ricky97gr/homeOnline/internal/webservice/service/topic"
	"github.com/ricky97gr/homeOnline/pkg/paginate"
)

func NormalGetAllTopic(ctx *gin.Context) {
	topics, err := topic.NormalGetAllTopic()
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, topics, int64(len(topics)))

}

func AdminGetAllTopic(ctx *gin.Context) {
	q, err := paginate.GetPageQuery(ctx)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	topics, count, err := topic.AdminGetAllTopic(q)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, topics, count)
}

func AdminUpdateTopic(ctx *gin.Context) {
	type tmpT struct {
		Uuid   string `json:"uuid"`
		IsShow bool   `json:"isShow"`
	}
	info := &tmpT{}
	err := ctx.ShouldBind(&info)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	err = topic.AdminUpdateTopic(info.Uuid, info.IsShow)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, "update successfully", 1)

}

func AdminDeleteTopic(ctx *gin.Context) {
	type tmpT struct {
		Uuid string `json:"uuid"`
		Name string `json:"name"`
	}
	info := &tmpT{}
	err := ctx.ShouldBind(&info)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	err = topic.AdminDeleteTopic(info.Uuid)
	if err != nil {
		newlog.Logger.Errorf("failed to delete topic: %+v, err: %+v\n", info, err)
		response.Failed(ctx, response.ErrStruct)
		return
	}
	//TODO: 有问题
	err = sendlog.SendOperationLog(ctx.Request.Header.Get("userName"), "cn", sendlog.DeleteTopic, info.Name)
	if err != nil {
		newlog.Logger.Errorf("failed to send operation log: %+v, err: %+v\n", info, err)
	}
	response.Success(ctx, "delete successfully", 1)
}

func AdminCreateTopic(ctx *gin.Context) {
	topics := &topic.UITopic{}
	err := ctx.ShouldBindJSON(topics)
	if err != nil {
		newlog.Logger.Errorf("failed bind json, err: %+v\n", err)
		response.Failed(ctx, response.ErrStruct)
		return
	}
	topics.Creator = ctx.Request.Header.Get("userName")
	err = topic.AdminCreateTopic(topics)
	if err != nil {
		newlog.Logger.Errorf("failed to create topic: %+v, err: %+v\n", topics, err)
		response.Failed(ctx, response.ErrDB)
		return
	}
	err = sendlog.SendOperationLog(ctx.Request.Header.Get("userName"), "cn", sendlog.NewTopic, topics.Name)
	if err != nil {
		newlog.Logger.Errorf("failed to send operation log: %+v, err: %+v\n", topics, err)
	}
	response.Success(ctx, "", 1)
}
