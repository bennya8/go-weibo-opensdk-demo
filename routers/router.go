package routers

import (
	"github.com/astaxie/beego"
	"go-weibo-opensdk-demo/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/analyze", &controllers.AnalyzeController{})
}
