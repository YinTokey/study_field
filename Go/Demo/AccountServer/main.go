package main

import (
	"AccountServer/infra"
	"AccountServer/infra/base"
)

type User struct {
	Name string
}

func main() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDataBaseStarter{})
}
