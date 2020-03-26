package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"go-weibo-opensdk-demo/sdk/weibo"
)

type AnalyzeController struct {
	beego.Controller
}

func (c *AnalyzeController) Get() {
	token := c.GetString("token")

	appKey := beego.AppConfig.String("openWeiboAppKey")
	appSecret := beego.AppConfig.String("openWeiboAppSecret")

	sdk := weibo.NewClient(appKey, appSecret)
	rsp, err := sdk.StatusesPublicLine(token, 10, 1)
	if err != nil {
		c.Data["Msg"] = fmt.Sprintf("%s", err)
		c.TplName = "msg.html"
		return
	}
	fmt.Println(rsp)

	rsp2, err := sdk.FriendsFriends(token)
	if err != nil {
		c.Data["Msg"] = fmt.Sprintf("%s", err)
		c.TplName = "msg.html"
		return
	}

	c.Data["FriendsFriends"] = rsp2
	c.TplName = "analyze/list.html"

}
