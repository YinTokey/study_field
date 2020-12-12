package service

import (
	"encoding/json"
	"fmt"
	"go_wallpaper/model"
	"io/ioutil"
	"net/http"
	"os"
)

var um_all_url string = "https://api.unsplash.com/photos/"
// https://api.unsplash.com/photos/?client_id=YOUR_ACCESS_KEY

type UmCollectService struct {

}

func NewUmCollectService() UmCollectService {
	return UmCollectService{
	}
}

func (service *UmCollectService) Papular() []model.Photo {

	key := os.Getenv("UNSPLASH_ACCESS_KEY")
	fmt.Println("key ", key)

	url := um_all_url + "?client_id=" + key

	// 网络请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	body,err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, &result)
	}

	fmt.Println(result)

	return nil
}