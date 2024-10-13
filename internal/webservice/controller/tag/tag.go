package tag

import (
	"github.com/gin-gonic/gin"

	"github.com/ricky97gr/homeOnline/internal/pkg/newlog"
	"github.com/ricky97gr/homeOnline/internal/pkg/response"
	"github.com/ricky97gr/homeOnline/internal/pkg/sendlog"
	"github.com/ricky97gr/homeOnline/internal/webservice/service/tag"
	"github.com/ricky97gr/homeOnline/pkg/paginate"
)

func NormalGetAllTag(ctx *gin.Context) {
	tags, err := tag.NormalGetAllTag()
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, tags, int64(len(tags)))

}

func AdminGetAllTag(ctx *gin.Context) {
	q, err := paginate.GetPageQuery(ctx)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	tags, count, err := tag.AdminGetAllTag(q)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, tags, count)
}

func AdminUpdateTag(ctx *gin.Context) {
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
	err = tag.AdminUpdateTag(info.Uuid, info.IsShow)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, "update successfully", 1)

}

func AdminDeleteTag(ctx *gin.Context) {
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
	err = tag.AdminDeleteTag(info.Uuid)
	if err != nil {
		newlog.Logger.Errorf("failed to delete tag: %+v, err: %+v\n", info, err)
		response.Failed(ctx, response.ErrStruct)
		return
	}
	//TODO: 有问题
	err = sendlog.SendOperationLog(ctx.Request.Header.Get("userName"), "cn", sendlog.DeleteTag, info.Name)
	if err != nil {
		newlog.Logger.Errorf("failed to send operation log: %+v, err: %+v\n", info, err)
	}
	response.Success(ctx, "delete successfully", 1)
}

func AdminCreateTag(ctx *gin.Context) {
	tags := &tag.UITag{}
	err := ctx.ShouldBindJSON(tags)
	if err != nil {
		newlog.Logger.Errorf("failed bind json, err: %+v\n", err)
		response.Failed(ctx, response.ErrStruct)
		return
	}
	tags.Creator = ctx.Request.Header.Get("userName")
	err = tag.AdminCreateTag(tags)
	if err != nil {
		newlog.Logger.Errorf("failed to create tag: %+v, err: %+v\n", tags, err)
		response.Failed(ctx, response.ErrDB)
		return
	}
	err = sendlog.SendOperationLog(ctx.Request.Header.Get("userName"), "cn", sendlog.NewTag, tags.Name)
	if err != nil {
		newlog.Logger.Errorf("failed to send operation log: %+v, err: %+v\n", tags, err)
	}
	response.Success(ctx, "", 1)
}
