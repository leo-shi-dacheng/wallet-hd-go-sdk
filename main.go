package main

import (
	"flag"
	"github.com/leo-shi-dacheng/wallet-hd-go-sdk/config"
)

func main() {
	var f = flag.String("c", "config.yml", "config file")
	flag.Parse()
	//println(*f, "f")

	conf, err := config.New(*f)

	if err != nil {
		panic(err)
	}

	println(conf, "conf")
}
