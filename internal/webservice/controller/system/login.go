package system

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ricky97gr/homeOnline/internal/pkg/newlog"
	"github.com/ricky97gr/homeOnline/internal/pkg/response"
	"github.com/ricky97gr/homeOnline/internal/pkg/sendlog"
	"github.com/ricky97gr/homeOnline/internal/webservice/middleware"
	"github.com/ricky97gr/homeOnline/internal/webservice/model"
	"github.com/ricky97gr/homeOnline/internal/webservice/service/socre"
	"github.com/ricky97gr/homeOnline/internal/webservice/service/system"
)

type LoginInfo struct {
	Phone  string `json:"phone"`
	Passwd string `json:"password"`
}

func Login(ctx *gin.Context) {
	var info LoginInfo
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		response.Failed(ctx, response.ErrStruct, "struct error")
		return
	}
	if info.Phone == "" {
		response.Failed(ctx, response.ErrStruct, "please input phone")
		return
	}
	user, err := system.GetUserByPhone(info.Phone, info.Passwd)
	if err != nil {
		//登陆失败
		response.Failed(ctx, response.ErrUserNameOrPassword)
		return
	}

	if user.UserID == "" {
		response.Failed(ctx, response.ErrUserNameOrPassword, "user not exist")
		return
	}
	if system.UpdateUserLastLogin(user.UserID) != nil {
		response.Failed(ctx, response.ErrDB, "")
		return
	}

	newlog.Logger.Infof("user <%s> login successfully\n", user.NickName)
	token, err := middleware.GenerateToken(user.UserID, user.NickName, user.Role)
	if err != nil {
		response.Failed(ctx, response.ErrUserNameOrPassword)
		return
	}
	err = middleware.StoreToken(token)
	if err != nil {
		response.Failed(ctx, response.ErrRedis)
		return
	}
	err = sendlog.SendOperationLog(user.UserID, "cn", sendlog.LoginCode)
	if err != nil {
		newlog.Logger.Errorf("failed to send operation log, err: %+v\n", err)
	}
	err = socre.AddScore(user.UserID, model.Login)
	if err != nil {
		newlog.Logger.Errorf("failed to add score, err: %+v\n", err)
		//return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"code":     200,
			"msg":      "handle successfully",
			"token":    token,
			"userID":   user.UserID,
			"role":     user.Role,
			"userName": user.NickName,
		})
}
