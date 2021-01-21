package job

import (
	"encoding/json"
	"fmt"
	"github.com/disintegration/imageorient"
	"go_wallpaper/internal/picture/model"
	"go_wallpaper/pkg"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
)

type RandomAcgJob struct {
}

var url_1 string = "http://api.btstu.cn/sjbz/?lx=suiji"
var url_2 string = "http://api.btstu.cn/sjbz/?lx=dongman"

func (j *RandomAcgJob) FetchJob_1() {
	j.FetchLink(url_1)
}

func (j *RandomAcgJob) FetchJob_2() {
	j.FetchLink(url_2)
}

func NewRandomAcgJob() *RandomAcgJob {
	return &RandomAcgJob{}
}

func (j *RandomAcgJob) FetchLink(url string) {

	validUrl := j.GetRealImageUrl(url)

	// 过滤掉长度过段的
	if len(validUrl) < 35 {
		return
	}
	img, err := GetImage(validUrl)

	if err != nil {
		fmt.Println("get img err ", err)
	}

	// 构建model
	var width = img.Bounds().Max.X
	var height = img.Bounds().Max.Y

	model := j.ComposeModel(validUrl, width, height)

	// 写入redis
	j.StoreRedis(model)

	fmt.Println("写入 redis ", model.PictureId)
}

// RedirectFunc 重定向禁止
func RedirectFunc(req *http.Request, via []*http.Request) error {
	//fmt.Println(req.RequestURI)
	// 如果返回 非nil 则禁止向下重定向 返回nil 则 一直向下请求 10 次 重定向
	return http.ErrUseLastResponse
}

func (j *RandomAcgJob) GetRealImageUrl(url string) string {
	// 禁止重定向
	client := &http.Client{CheckRedirect: RedirectFunc}
	rep, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer rep.Body.Close()

	//读取响应的结果
	localtion := rep.Header["Location"]

	return localtion[0]

}

func GetImage(url string) (image.Image, error) {
	//fmt.Println("start request ", url)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("get image request err ", err)
		return nil, err
	}
	img, _, err := imageorient.Decode(resp.Body)
	if err != nil {
		fmt.Println("decode image err ", err)
		return nil, err
	}
	return img, nil
}

func (j *RandomAcgJob) ComposeModel(url string, width int, height int) *model.Acg {

	id, _ := pkg.NewGuid()

	acgModel := &model.Acg{
		PictureId:     id,
		ImageUrl:      url,
		LargeImageUrl: url,
		Name:          "",
		Description:   "",
		Author:        "",
		Width:         width,
		Height:        height,
		Likes:         0,
		Categories:    "acg",
	}

	return acgModel

}

func (j *RandomAcgJob) StoreRedis(model *model.Acg) {

	key := model.ImageUrl

	buf, err := json.Marshal(model)
	if err != nil {
		fmt.Println("json.Marshal(model) 失败", err)
		return
	}
	// 0 表示不会过期
	pkg.RedisClient.Set(key, buf, 0)

}

func (j *RandomAcgJob) StoreMysql(model *model.Acg) {

}
