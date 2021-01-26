package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/goinggo/mapstructure"
	"go_wallpaper/internal/picture/model/px"
	"net/http"
)

const pupular_url = "https://api.500px.com/v1/photos?feature=popular"

var ctx = context.Background()

type PxCollectService struct {
}

func NewPxCollectService() PxCollectService {
	return PxCollectService{}
}

// 拉取 500px 首页
func (service *PxCollectService) Papular() []px.Photo {

	//return px.FindAllPhotos()
	return nil
}

func (service *PxCollectService) updateToDatabase(photos []px.Photo) {

	fmt.Println(len(photos))

	//for _, photo := range photos {
	//	photo.SavePhoto()
	//
	//}
}

func (service *PxCollectService) Request500pxPapuplar() []px.Photo {
	var info []px.Photo

	// 网络请求
	resp, err := http.Get(pupular_url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 本地测试json
	//filePtr, err := os.Open("./px.json")
	//if err != nil {
	//	fmt.Println("文件打开失败 [Err:%s]", err.Error())
	//	return info
	//}
	//defer filePtr.Close()

	// 初始化请求变量结构
	formData := make(map[string]interface{})

	// 创建json解码器
	//json.NewDecoder(filePtr).Decode(&formData)

	// 调用json包的解析，解析请求body
	json.NewDecoder(resp.Body).Decode(&formData)

	// 类型强转
	var photosMapArr []interface{} = formData["photos"].([]interface{})

	//result :=  make([]Photo,0)

	var result []px.Photo

	fmt.Println("准备解析")
	for _, phMap := range photosMapArr {
		var photo px.Photo
		if err := mapstructure.Decode(phMap, &photo); err != nil {
			fmt.Println(err)
		}

		realmap := phMap.(map[string]interface{})

		var urls []interface{} = realmap["image_url"].([]interface{})
		var str string = urls[0].(string)
		photo.ImageURL = str

		//fmt.Println(str)
		result = append(result, photo)

		//photo.SavePhoto()
	}

	info = result

	//json.Unmarshal([]byte(jsonStr), &info)

	return info
}
