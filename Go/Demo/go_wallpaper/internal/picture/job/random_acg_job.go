package job

import (
	"fmt"
	"github.com/disintegration/imageorient"
	"image"
	"net/http"
)

type RandomAcgJob struct {
}

var url_1 string = "http://api.btstu.cn/sjbz/?lx=suiji"
var url_2 string = "http://api.btstu.cn/sjbz/?lx=dongman"

func NewRandomAcgJob() *RandomAcgJob {
	return &RandomAcgJob{}
}

func (j *RandomAcgJob) FetchLink_1() {

	img, err := j.GetImage(url_1)
	if err != nil {
		fmt.Println("解码图片错误 ", err)
	}
	fmt.Println(img)
}

func (j *RandomAcgJob) GetImage(url string) (image.Image, error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	img, _, err := imageorient.Decode(resp.Body)
	if err != nil {
		return nil, err
	}
	return img, nil
}
