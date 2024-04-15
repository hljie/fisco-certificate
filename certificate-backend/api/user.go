package api

import (
	"certificate-backend/model"
	"certificate-backend/serializer"
	"certificate-backend/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var registerService service.UserRegisterService
	if err := c.ShouldBind(&registerService); err == nil {
		res := registerService.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserMe 用户详情
func UserMe(c *gin.Context) {
	user := CurrentUser(c)
	res := serializer.BuildUserResponse(*user)
	c.JSON(200, res)
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}

// 用户列表

type UserListAo struct {
	serializer.PageAo
	Username string `json:"username" required:"false"`
	Nickname string `json:"nickname" required:"false"`
	Status   string `json:"status" required:"false"`
	Role     string `json:"role" required:"false"`
}

func UserList(c *gin.Context) {
	ao := new(UserListAo)
	if err := c.ShouldBind(ao); err == nil {
		// 分页查询
		if ao.PageSize == 0 {
			ao.PageSize = 10
		}
		if ao.Page == 0 {
			ao.Page = 1
		} else {
			ao.Page = ao.Page - 1
		}
		var sql = "1 = 1"
		if ao.Username != "" {
			sql += " and username like '%" + ao.Username + "%'"
		}
		if ao.Nickname != "" {
			sql += " and nickname like '%" + ao.Nickname + "%'"
		}
		if ao.Status != "" {
			sql += " and status = '" + ao.Status + "'"
		}
		if ao.Role != "" {
			sql += " and role = '" + ao.Role + "'"
		}
		var users []model.User
		tx := model.DB.Find(&users, sql).Order("id desc").Limit(ao.PageSize).Offset(ao.Page * ao.PageSize)
		if tx.Error != nil {
			c.JSON(200, ErrorResponse(tx.Error))
			return
		}
		var count int64
		tx = model.DB.Model(&model.User{}).Where(sql).Count(&count)
		if tx.Error != nil {
			c.JSON(200, ErrorResponse(tx.Error))
			return
		}
		c.JSON(200, serializer.BuildPage(users, count))
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 更新用户的状态和角色
type UserUpdateAo struct {
	ID     uint   `json:"id" binding:"required"`
	Status string `json:"status" `
	Role   string `json:"role" `
}

func UserUpdate(c *gin.Context) {
	ao := new(UserUpdateAo)
	if err := c.ShouldBind(ao); err == nil {
		var user model.User
		tx := model.DB.First(&user, ao.ID)
		if tx.RowsAffected == 0 {
			c.JSON(200, serializer.Response{
				Code: 404,
				Msg:  "用户不存在",
			})
			return
		}
		if ao.Status != "" {
			user.Status = ao.Status
		}
		if ao.Role != "" {
			user.Role = ao.Role
		}
		tx = model.DB.Save(&user)
		if tx.Error != nil {
			c.JSON(200, ErrorResponse(tx.Error))
			return
		}
		c.JSON(200, serializer.BuildUserResponse(user))
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 自己修改呢称或者密码
type UserUpdateMeAo struct {
	Nickname    string `json:"nickname"`
	OldPassword string `json:"oldPassword"`
	Password    string `json:"password"`
}

func UserUpdateMe(c *gin.Context) {
	ao := new(UserUpdateMeAo)
	if err := c.ShouldBind(ao); err == nil {
		user := CurrentUser(c)
		if ao.Nickname != "" {
			user.Nickname = ao.Nickname
		}
		if ao.OldPassword != "" && ao.Password != "" {
			if user.CheckPassword(ao.OldPassword) {
				err := user.SetPassword(ao.Password)
				if err != nil {
					c.JSON(200, serializer.Response{
						Code: 400,
						Msg:  "密码错误",
					})
					return
				}
			} else {
				c.JSON(200, serializer.Response{
					Code: 400,
					Msg:  "密码错误",
				})
				return
			}
		}
		tx := model.DB.Save(&user)
		if tx.Error != nil {
			c.JSON(200, ErrorResponse(tx.Error))
			return
		}
		c.JSON(200, serializer.BuildUserResponse(*user))
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
