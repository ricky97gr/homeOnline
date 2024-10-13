package system

import (
	"github.com/gin-gonic/gin"

	"github.com/ricky97gr/homeOnline/internal/pkg/newlog"
	"github.com/ricky97gr/homeOnline/internal/pkg/response"
	"github.com/ricky97gr/homeOnline/internal/webservice/model"
	"github.com/ricky97gr/homeOnline/pkg/validate"
)

type RegisterInfo struct {
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
	Passwd string `json:"passwd"`
}

func (r *RegisterInfo) Covert() *model.User {
	return &model.User{
		UserName: r.Name,
		Password: r.Passwd,
		Phone:    r.Phone,
		Email:    r.Email,
	}
}

func Register(ctx *gin.Context) {

	var info RegisterInfo
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		response.Failed(ctx, response.ErrStruct, "struct error")
		return
	}
	if !validate.IsEmailFormat(info.Email) {
		return
	}
	if !validate.IsMobileFormat(info.Phone) {
		return
	}

	user := info.Covert()
	newlog.Logger.Infof("%+v\n", user)
	//err = service.Register(user)
	if err != nil {
		return
	}
	response.Success(ctx, "", 1)

}
