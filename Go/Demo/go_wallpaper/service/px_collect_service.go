package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	. "go_wallpaper/model"
)

const pupular_url = "https://api.500px.com/v1/photos?feature=popular"

type PxCollectService struct {

}

func NewPxCollectService() PxCollectService {
	return PxCollectService{
	}
}

// 拉取 500px 首页
func (service *PxCollectService) Papular() {
	resp, err := http.Get(pupular_url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	s , err := ioutil.ReadAll(resp.Body)
	jsonStr := string(s)

	var info Page
	json.Unmarshal([]byte(jsonStr), &info)

	fmt.Println(info.Photos[1].User.Username)

	//return jsonStr
}
