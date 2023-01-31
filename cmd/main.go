package main

import (
	"gin_shop/conf"
	"gin_shop/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)

}
