package unsplash

import (
	"encoding/json"
	"fmt"
	"go_wallpaper/model"
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

	for i := 0; i < 3; i++ {
		url := unsplash_base_url + "/photos/" + "?client_id=" + key + "&per_page=100" + "&page=" + strconv.Itoa(i)

		fmt.Println("开始请求数据, url = ", url)

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

		result := make([]model.Picture, 0)

		for _, picMap := range mapArr {
			urls := picMap["urls"].(map[string]interface{})
			user := picMap["user"].(map[string]interface{})

			pic := model.Picture{}
			pic.PictureId = picMap["id"].(string)
			pic.Height = picMap["height"].(float64)
			pic.Width = picMap["width"].(float64)
			pic.Name = picMap["alt_description"].(string)

			desc := picMap["description"]
			if desc != nil {
				pic.Description = desc.(string)
			}

			pic.Likes = picMap["likes"].(float64)
			pic.ImageUrl = urls["small"].(string)
			pic.LargeImageUrl = urls["full"].(string)
			pic.Author = user["name"].(string)
			//pic.UpdatedAt = picMap["updated_at"]
			//pic.Categories = picMap["categories"].([]string)

			result = append(result, pic)
		}

	}

	//return nil
}
