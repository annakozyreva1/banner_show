package main

import (
	"flag"
	"github.com/annakozyreva1/banner_show/log"
	"github.com/annakozyreva1/banner_show/selector"
	"github.com/annakozyreva1/banner_show/web"
)

var logger = log.Logger

var (
	webAddress = flag.String("addr", ":7878", "web api address")
	config     = flag.String("conf", "./config.csv", "banner config")
)

func main() {
	flag.Parse()
	sel := selector.InitSelector(*config)
	web.Run(*webAddress, sel)
}
