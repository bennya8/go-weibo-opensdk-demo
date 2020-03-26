package main

import (
	"github.com/astaxie/beego"
	_ "go-weibo-opensdk-demo/routers"
	"log"
	"path/filepath"
)

func main() {

	path, err := filepath.Abs(filepath.Dir("."))
	if err != nil {
		log.Fatal(err)
	}

	beego.BConfig.WebConfig.ViewsPath = path + "/views"
	beego.Run()
}
