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
func (service *PxCollectService) Papular() Page {
	resp, err := http.Get(pupular_url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	s , err := ioutil.ReadAll(resp.Body)
	jsonStr := string(s)

	var info Page
	json.Unmarshal([]byte(jsonStr), &info)

	//fmt.Println(info.Photos[1].User.Username)

	service.updateToDatabase(info.Photos)

	return info
}

func (service *PxCollectService) updateToDatabase(photos []Photo) {

	fmt.Println(len(photos))

	for _, photo := range photos {
		photo.SavePhoto()


	}
}
