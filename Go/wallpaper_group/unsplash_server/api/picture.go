package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"unsplash_server/service"
)

func FetchPapular(c *gin.Context) {

	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	fmt.Println("page %s , page sie %s", page, pageSize)

	fetchFromServcie(c, page, pageSize)
}

func fetchFromServcie(c *gin.Context, page int, pageSize int) {
	service := service.NewPictureService()

	if err := c.ShouldBind(&service); err == nil {
		res, err := service.Papular(page, pageSize)
		if err != nil {
			c.JSON(200, "error happend")
		} else {
			c.JSON(200, res)
		}

	} else {
		c.JSON(200, err)
	}

}
