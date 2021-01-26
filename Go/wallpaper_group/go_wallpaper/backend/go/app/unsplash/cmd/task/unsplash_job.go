package task

import (
	"encoding/json"
	"fmt"
	"go_wallpaper/backend/go/internal/picture/dao"
	"go_wallpaper/backend/go/internal/picture/model"
	"go_wallpaper/backend/go/pkg"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

// https://api.unsplash.com/photos/?client_id=YOUR_ACCESS_KEY
var unsplash_base_url string = "https://api.unsplash.com"

type UnsplashJob struct {
}

func NewUnsplashJob() *UnsplashJob {
	return &UnsplashJob{}
}

func (u *UnsplashJob) FetchPics() {
	key := os.Getenv("UNSPLASH_ACCESS_KEY")
	d := dao.NewPictureDao(pkg.InstanceDB())

	for i := 0; i < 200; i++ {
		url := unsplash_base_url + "/photos/" + "?client_id=" + key + "&per_page=3000" + "&page=" + strconv.Itoa(i)

		fmt.Println("开始请求数据, page ", i)

		// 网络请求
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		var mapArr []map[string]interface{}

		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			err = json.Unmarshal(body, &mapArr)
		}

		// 数据筛选

		//result := make([]model.Picture, 0)

		//fmt.Println("开始筛选数据 ")

		for _, picMap := range mapArr {
			urls := picMap["urls"].(map[string]interface{})
			user := picMap["user"].(map[string]interface{})

			var fw float64 = picMap["width"].(float64)
			var fh float64 = picMap["height"].(float64)

			var width int = int(fw)
			if width < 2000 {
				continue
			}

			var height int = int(fh)

			// 过滤掉竖的图
			if width-height < 300 {
				continue
			}

			pic := model.Picture{}
			pic.Height = height
			pic.Width = width

			desc := picMap["description"]
			if desc != nil {
				pic.Description = desc.(string)
			}

			name := picMap["alt_description"]
			if name != nil {
				pic.Name = name.(string)
			}

			flikes := picMap["likes"].(float64)

			pic.Likes = int32(flikes)
			pic.ImageUrl = urls["small"].(string)
			pic.LargeImageUrl = urls["full"].(string)
			pic.Author = user["name"].(string)

			//fmt.Println(picMap["categories"])

			//pic.Categories = picMap["categories"].([]string)

			id, _ := pkg.NewGuid()
			pic.PictureId = id

			d.AddPicture(pic)

			//result = append(result, pic)
		}

	}

	//return nil
}
