package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var pupular_url = "https://api.500px.com/v1/photos?feature=popular"

func main () {

	resp, err := http.Get(pupular_url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	s , err := ioutil.ReadAll(resp.Body)
	jsonStr := string(s)
	fmt.Println(jsonStr)

}