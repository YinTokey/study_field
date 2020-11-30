package api

import (
	"Week02/service"
	"github.com/gin-gonic/gin"
)

func QueryPicture(c *gin.Context) {

	// 表中存有一条记录
///	pic := &models.Picture{URL: "https://wwww.baidu.com", UserID: "123", Author: "JACK"}
	id := 123

	sv := service.NewPictureService()
	err := c.ShouldBind(&sv)
	if err != nil {
		c.JSON(500, "server error")
		return;
	}
	res, err := sv.Query(id)
	if err != nil {
		// 查询错误
		c.JSON(404, "can not find")
		return;
	}
	c.JSON(200, res)

}


