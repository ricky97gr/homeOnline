package category

import (
	"github.com/gin-gonic/gin"

	"github.com/ricky97gr/homeOnline/internal/pkg/newlog"
	"github.com/ricky97gr/homeOnline/internal/pkg/response"
	"github.com/ricky97gr/homeOnline/internal/pkg/sendlog"
	"github.com/ricky97gr/homeOnline/internal/webservice/service/category"
	"github.com/ricky97gr/homeOnline/pkg/paginate"
)

func NormalGetAllCategory(ctx *gin.Context) {
	cates, err := category.NormalGetAllCategory()
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, cates, int64(len(cates)))

}

func AdminGetAllCategory(ctx *gin.Context) {
	q, err := paginate.GetPageQuery(ctx)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	cates, count, err := category.AdminGetAllCategory(q)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, cates, count)

}

func AdminUpdateCategory(ctx *gin.Context) {
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
	err = category.AdminUpdateCategory(info.Uuid, info.IsShow)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, "update successfully", 1)

}

func AdminDeleteCategory(ctx *gin.Context) {
	type cate struct {
		Uuid string `json:"uuid"`
		Name string `json:"name"`
	}
	info := &cate{}
	err := ctx.ShouldBindJSON(info)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	err = category.AdminDeleteCategory(info.Uuid)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	err = sendlog.SendOperationLog(ctx.Request.Header.Get("userName"), "cn", sendlog.DeleteCategory, info.Name)
	if err != nil {
		newlog.Logger.Errorf("failed to send operation log: %+v, err: %+v\n", info, err)
	}
	response.Success(ctx, "delete successfully", 1)

}

func AdminCreateCategory(ctx *gin.Context) {
	cate := &category.UICategory{}
	err := ctx.ShouldBindJSON(cate)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	cate.Creator = ctx.Request.Header.Get("userName")
	err = category.AdminCreateCategory(cate)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	err = sendlog.SendOperationLog(ctx.Request.Header.Get("userName"), "cn", sendlog.NewCategory, cate.Name)
	if err != nil {
		newlog.Logger.Errorf("failed to send operation log: %+v, err: %+v\n", cate, err)
	}
	response.Success(ctx, "", 1)
}
