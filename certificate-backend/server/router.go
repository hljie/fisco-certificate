package server

import (
	"certificate-backend/api"
	"certificate-backend/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 用户列表

		// 需要登录保护的
		user := v1.Group("")
		user.Use(middleware.AuthRequired())
		{
			// User Routing
			user.GET("user/me", api.UserMe)
			user.GET("user/logout", api.UserLogout)
			user.POST("user/update/me", api.UserUpdateMe)

			user.POST("/certificate/me", api.GetMyCertificateList)
			user.POST("/certificate/apply", api.ApplyCertificate)
			user.POST("/certificate/id", api.GetCertificate)
			user.POST("/certificate/update", api.UpdatePendingCertificate)
		}

		teacher := v1.Group("")
		teacher.Use(middleware.AdminAuthRequired())
		{
			teacher.POST("users", api.UserList)
			teacher.POST("user/update", api.UserUpdate)
		}

		admin := v1.Group("")
		admin.Use(middleware.TeachingAuthRequired())
		{

			admin.POST("certificates", api.GetPendingCertificateList)
			admin.POST("certificate/approve", api.ApproveCertificate)
		}

	}
	return r
}
