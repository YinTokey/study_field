package api

import (
	"Week02/services"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Response struct {

}


func QueryPicture(c *gin.Context) {

	// 已知表中存有一条记录
///	pic := &models.Picture{URL: "https://wwww.baidu.com", UserID: "123", Author: "JACK"}
	id := 123

	service := services.NewPictureService()
	err := c.ShouldBind(&service)
	if err != nil {
		c.JSON(500, "server error")
		return;
	}
	res, err := service.Query(id)

	if errors.Is(err,sql.ErrNoRows) {
		c.JSON(404, fmt.Sprint(err))
	} else if err != nil {
		c.JSON(404, "query error")
	} else {
		c.JSON(200, res)
	}

}


