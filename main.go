package main

import (
	"flag"
	"fmt"
	"os"
	"xyz/test/helloworld/config"
	"xyz/test/helloworld/routers"

	"github.com/gin-gonic/gin"

	"go.etcd.io/etcd/api/v3/version"
)

const (
	PROGRAM_NAME    = "helloworld"
	PROGRAM_VERSION = "0.1.0"
)

var configFilepath string

func init() {
	// flag.StringVar(&configFilepath, "conf", "conf/config.default.ini", "config file path")
	rev := flag.Bool("rev", false, "print rev")
	flag.Parse()

	if *rev {
		fmt.Printf("[%s v%s]\n[etcd %s]\n",
			PROGRAM_NAME, PROGRAM_VERSION,
			version.Version,
		)
		os.Exit(0)
	}
}

func main() {
	config, err := config.Init(configFilepath)
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.UseRawPath = true
	routers.InitRouters(r, config)
	r.Run(":" + config.Port)
}
