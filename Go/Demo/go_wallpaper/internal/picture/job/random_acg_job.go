package job

import (
	"fmt"
	"github.com/disintegration/imageorient"
	"image"
	_ "image/jpeg"
	_ "image/png"
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

	validUrl := j.GetRealImageUrl(url_2)

	// 过滤掉长度过段的
	if len(validUrl) < 35 {
		return
	}
	fmt.Println(validUrl)

	img, err := GetImage(validUrl)

	if err != nil {
		fmt.Println("get img err ", err)
	}

	fmt.Println(img.Bounds())

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
	fmt.Println("start request ", url)
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
