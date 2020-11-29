package api

import (
	"Week02/models"
	"Week02/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AddPicture(c *gin.Context) {

	pic := &models.Picture{URL: "https://wwww.baidu.com", UserID: "123", Author: "JACK"}

	sv := service.NewPictureService()
	err := c.ShouldBind(&sv)
	if err != nil {
		//c.JSON(200, ErrorResponse(err))
		return
	}
	err = sv.AddPicture(pic)
	if err != nil {
		// 查询错误
		fmt.Println(err)
		return
	}
	//c.JSON(200, res)

}

func QueryPicture(c *gin.Context) {
	sv := service.NewPictureService()
	err := c.ShouldBind(&sv)
	if err != nil {
		//c.JSON(200, ErrorResponse(err))
		return;
	}
	res, err := sv.GetAll()
	if err != nil {
		// 查询错误
		return;
	}
	c.JSON(200, res)

}


