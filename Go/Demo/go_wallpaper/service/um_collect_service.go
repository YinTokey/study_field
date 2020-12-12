package service

import (
	"encoding/json"
	"fmt"
	"go_wallpaper/model/px"
	"io/ioutil"
	"net/http"
	"os"
)

// https://api.unsplash.com/photos/?client_id=YOUR_ACCESS_KEY
var unsplash_base_url string = "https://api.unsplash.com"

type UmCollectService struct {
}

func NewUmCollectService() UmCollectService {
	return UmCollectService{}
}

func (service *UmCollectService) Papular() []px.Photo {

	key := os.Getenv("UNSPLASH_ACCESS_KEY")

	url := unsplash_base_url + "/photos/" + "?client_id=" + key

	// 网络请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var result []map[string]interface{}

	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, &result)
	}

	fmt.Println(result)

	return nil
}
