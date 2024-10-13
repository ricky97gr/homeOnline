package middleware

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"

	"github.com/ricky97gr/homeOnline/internal/pkg/newlog"
	"github.com/ricky97gr/homeOnline/internal/pkg/response"
	"github.com/ricky97gr/homeOnline/internal/webservice/database/redis"
	"github.com/ricky97gr/homeOnline/internal/webservice/model"
)

var JwtStr = []byte("这是jwt认证密钥")

const (
	expiration = 60 * time.Minute
)

type Claims struct {
	UserID   string
	UserName string
	Role     int
	jwt.StandardClaims
}

// 普通用户
func AuthNormal() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ""
		if ctx.Request.URL.String() == "/normalUser/ws" {
			token = ctx.Request.Header.Get("Sec-Websocket-Protocol")
		} else {
			token = ctx.Request.Header.Get("token")
		}
		fmt.Println(token)
		if token == "" {
			ctx.Abort()
			response.Failed(ctx, response.ErrAuth)
			return
		}
		if !isTokenExist(token) {
			ctx.Abort()
			response.Failed(ctx, response.ErrAuth)
			return
		}
		if restoreToken(token) != nil {
			ctx.Abort()
			response.Failed(ctx, response.ErrAuth)
			return
		}
		claims, err := parseToken(token)
		if err != nil {
			newlog.Logger.Errorf("failed to parse token, err:%+v\n", err)
		}
		ctx.Request.Header.Set("userName", claims.UserName)
		ctx.Request.Header.Set("role", strconv.Itoa(claims.Role))
		ctx.Request.Header.Set("userID", claims.UserID)
		ctx.Request.Header.Set("clientIP", ctx.ClientIP())
		newlog.Logger.Infof("user:%s, auth successfully\n", claims.UserName)
		ctx.Next()
	}
}

// 普通管理员用户
func AuthAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token == "" {
			ctx.Abort()
			response.Failed(ctx, response.ErrAuth)
			return
		}
		if !isTokenExist(token) {
			ctx.Abort()
			response.Failed(ctx, response.ErrAuth)
			return
		}
		claims, err := parseToken(token)
		if err != nil {
			newlog.Logger.Errorf("failed to parse token, err:%+v\n", err)
		}
		if claims.Role != model.Admin && claims.Role != model.SuperAdmin {
			ctx.Abort()
			return
		}
		if restoreToken(token) != nil {
			ctx.Abort()
			response.Failed(ctx, response.ErrRedis)
			return
		}

		ctx.Request.Header.Set("userName", claims.UserName)
		ctx.Request.Header.Set("role", strconv.Itoa(claims.Role))
		ctx.Request.Header.Set("userID", claims.UserID)
		newlog.Logger.Infof("user:%s, auth successfully\n", claims.UserName)
		ctx.Next()
	}
}

// 超级管理员用户
func AuthSuperAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token == "" {
			ctx.Abort()
			response.Failed(ctx, response.ErrAuth)
			return
		}
		if !isTokenExist(token) {
			ctx.Abort()
			response.Failed(ctx, response.ErrAuth)
			return
		}
		claims, err := parseToken(token)
		if err != nil {
			newlog.Logger.Errorf("failed to parse token, err:%+v\n", err)
		}
		if claims.Role != model.SuperAdmin {
			ctx.Abort()
			return
		}
		if restoreToken(token) != nil {
			ctx.Abort()
			response.Failed(ctx, response.ErrRedis)
			return
		}

		ctx.Request.Header.Set("userName", claims.UserName)
		ctx.Request.Header.Set("role", strconv.Itoa(claims.Role))
		ctx.Request.Header.Set("userID", claims.UserID)
		newlog.Logger.Infof("user:%s, auth successfully", claims.UserName)
		ctx.Next()
	}
}

func GenerateToken(userID, userName string, role int) (string, error) {
	claim := &Claims{
		UserName: userName,
		UserID:   userID,
		Role:     role,
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString(JwtStr)
}

func parseToken(token string) (*Claims, error) {
	tokenClaim, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return JwtStr, nil
	})
	if _, ok := tokenClaim.Claims.(*Claims); ok {
		if tokenClaim.Claims.Valid() == nil {
			return tokenClaim.Claims.(*Claims), nil
		}
	}
	return nil, err
}

func StoreToken(token string) error {
	c, err := redis.GetRedisClient()
	if err != nil {
		return err
	}
	return c.Set(token, nil, expiration).Err()
}

func restoreToken(token string) error {
	return StoreToken(token)
}

func isTokenExist(token string) bool {
	c, err := redis.GetRedisClient()
	if err != nil {
		return false
	}
	return c.Get(token).Err() == nil
}
