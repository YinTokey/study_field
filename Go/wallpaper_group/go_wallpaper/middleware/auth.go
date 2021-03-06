package middleware

import (
	"go_wallpaper/internal/account/dao"
	"go_wallpaper/internal/account/serializer"
	//"go_wallpaper/internal/picture/model"
	"go_wallpaper/internal/account/model"
	"go_wallpaper/pkg"
	//"go_wallpaper/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("user_id")
		d := dao.NewUserDao(pkg.InstanceDB())
		if uid != nil {
			user, err := d.GetUser(uid)
			if err == nil {
				c.Set("user", &user)
			}
		}
		c.Next()
	}
}

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*model.User); ok {
				c.Next()
				return
			}
		}

		c.JSON(200, serializer.CheckLogin())
		c.Abort()
	}
}
