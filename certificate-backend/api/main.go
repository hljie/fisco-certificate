package api

import (
	"certificate-backend/model"
	"certificate-backend/serializer"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
)

// Ping 状态检查页面
func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "Pong",
	})
}

// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) *model.User {
	if user, _ := c.Get("user"); user != nil {
		if u, ok := user.(*model.User); ok {
			return u
		}
	}
	return nil
}

// ErrorResponse 返回错误消息
func ErrorResponse(err error) serializer.Response {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, e := range ve {
			return serializer.ParamErr(
				fmt.Sprintf("%s%s", fmt.Sprintf("Field.%s", e.Field()), fmt.Sprintf("Tag.Valid.%s", e.Tag())),
				err,
			)
		}
	}
	var unmarshalTypeError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalTypeError) {
		return serializer.ParamErr("JSON类型不匹配", err)
	}

	return serializer.ParamErr("参数错误", err)
}
