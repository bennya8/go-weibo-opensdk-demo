package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"go-weibo-opensdk-demo/sdk/weibo"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	code := c.GetString("code")

	appKey := beego.AppConfig.String("openWeiboAppKey")
	appSecret := beego.AppConfig.String("openWeiboAppSecret")
	redirectUrl := beego.AppConfig.String("openWeiboRedirect")

	sdk := weibo.NewClient(appKey, appSecret)

	if len(code) > 0 {
		rsp, err := sdk.OauthAccessToken(code, redirectUrl)
		if err != nil {
			c.Data["Msg"] = fmt.Sprintf("%s", err)
			c.TplName = "msg.html"
			return
		}

		c.Ctx.Redirect(302, "/analyze?access_token="+rsp.AccessToken)
		return
	}

	url := sdk.OauthAuthorize(redirectUrl)

	c.Data["Url"] = url
	c.TplName = "index.html"
}
