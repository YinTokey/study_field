package job

import (
	"fmt"
	"image"
	"io/ioutil"
	"log"
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

	img, err := j.GetImage(url_2)
	if err != nil {
		fmt.Println("解码图片错误 ", err)
	}
	fmt.Println(img)
}

// RedirectFunc 重定向禁止
func RedirectFunc(req *http.Request, via []*http.Request) error {
	fmt.Println(req.RequestURI)
	// 如果返回 非nil 则禁止向下重定向 返回nil 则 一直向下请求 10 次 重定向
	return http.ErrUseLastResponse
}

func (j *RandomAcgJob) GetImage(url string) (image.Image, error) {
	// 禁止重定向
	client := &http.Client{CheckRedirect: RedirectFunc}
	rep, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	//读取响应的结果
	data, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(data[:]))
	//输出响应的头
	for k, v := range rep.Header {
		fmt.Println(k, v)
	}
	defer rep.Body.Close()
	url1, err := rep.Location()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(url1.Host)

	//var r interface{}
	//
	//body, err := ioutil.ReadAll(resp.Body)
	//if err == nil {
	//	err = json.Unmarshal(body, &r)
	//}
	//
	//fmt.Println(resp.Header)

	return nil, nil

	//img, _, err := imageorient.Decode(resp.Body)
	//if err != nil {
	//	return nil, err
	//}
	//return img, nil
}
