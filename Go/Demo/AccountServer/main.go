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
	infra.Register(&base.DbxDatabaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	//infra.Register(&base.IrisServerStarter{})
}
