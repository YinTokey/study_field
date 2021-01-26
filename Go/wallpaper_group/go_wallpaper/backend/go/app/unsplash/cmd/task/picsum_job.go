package task

import (
	"encoding/json"
	"fmt"
	"go_wallpaper/backend/go/internal/picture/dao"
	"go_wallpaper/backend/go/internal/picture/model"
	"go_wallpaper/backend/go/pkg"
	"io/ioutil"
	"net/http"
)

//https://picsum.photos/v2/list?page=2&limit=100
var url_1 string = "https://picsum.photos/v2/list?"

type PicsumJob struct {
}

func NewPicsumJob() *PicsumJob {
	return &PicsumJob{}
}

func (j *PicsumJob) StartWork() {
	for i := 11; i < 500; i++ {
		j.FetchList(i)
	}
}

func (j *PicsumJob) FetchList(index int) {
	limit := 1000
	url := fmt.Sprintf("%spage=%v&limit=%v", url_1, index, limit)

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
	d := dao.NewPicsumDao(pkg.InstanceDB())

	for _, picMap := range mapArr {

		var fw float64 = picMap["width"].(float64)
		var fh float64 = picMap["height"].(float64)

		var width int = int(fw)
		if width < 2000 {
			continue
		}

		var height int = int(fh)

		model := &model.Picsum{}
		model.Width = width
		model.Height = height

		author := picMap["author"]
		if author != nil {
			model.Author = author.(string)
		}

		downloadUrl := picMap["download_url"]
		if author != nil {
			model.ImageUrl = downloadUrl.(string)
			model.LargeImageUrl = downloadUrl.(string)
		}

		id, _ := pkg.NewGuid()
		model.PictureId = id

		d.AddPicsum(model)
	}
}
