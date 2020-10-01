package main

import (
	"encoding/json"
	"fmt"
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
)

type User struct {
	Name string
}

func main() {
	u := User{
		Name: "jack",
	}
	data, err := json.Marshal(&u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	file := kvs.GetCurrentFilePath("config.ini", 1)

	conf := ini.NewIniFileCompositeConfigSource(file)
	port := conf.GetIntDefault("app.server.port", 18080)
	fmt.Println(port)
	fmt.Println(conf.GetDefault("app.name", "unknow"))

}
