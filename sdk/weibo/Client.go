package weibo

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ddliu/go-httpclient"
	"go-weibo-opensdk-demo/sdk/weibo/resp"
	"net/url"
	"strconv"
	"strings"
	"sync"
)

var instance *Client
var once sync.Once

type Client struct {
	host      string
	appKey    string
	appSecret string
	http      *httpclient.HttpClient
}

func NewClient(appKey string, appSecret string) *Client {
	once.Do(func() {
		instance = &Client{}
		instance.host = "https://api.weibo.com/"
		instance.appKey = appKey
		instance.appSecret = appSecret
		instance.http = httpclient.NewHttpClient()
	})
	return instance
}

func (c *Client) OauthAuthorize(redirectUrl string) string {
	host := c.host + "oauth2/authorize?client_id=%s&redirect_uri=%s&scope=all&"
	return fmt.Sprintf(host, c.appKey, url.QueryEscape(redirectUrl))
}

func (c *Client) OauthAccessToken(code string, redirectUrl string) (*resp.OauthAccessTokenRsp, error) {
	body, err := c.http.Post(c.host+"oauth2/access_token", map[string]string{
		"client_id":     c.appKey,
		"client_secret": c.appSecret,
		"grant_type":    "authorization_code",
		"code":          code,
		"redirect_uri":  redirectUrl,
	})
	if err != nil {
		return nil, err
	}
	var rsp *resp.OauthAccessTokenRsp

	bodyStr, _ := body.ToString()
	err = json.Unmarshal([]byte(bodyStr), &rsp)
	if err != nil || strings.Index(bodyStr, `"error":`) != -1 {
		return nil, errors.New(bodyStr)
	}
	return rsp, nil
}

func (c *Client) FriendsFriends(accessToken string) (*resp.FriendsFriendsRsp, error) {
	body, err := c.http.Get(c.host+"2/friendships/friends.json", map[string]string{
		"source":       c.appKey,
		"access_token": accessToken,
	})
	if err != nil {
		return nil, err
	}
	var rsp *resp.FriendsFriendsRsp

	bodyStr, _ := body.ToString()
	err = json.Unmarshal([]byte(bodyStr), &rsp)
	if err != nil || strings.Index(bodyStr, `"error":`) != -1 {
		return nil, errors.New(bodyStr)
	}
	return rsp, nil
}

func (c *Client) StatusesShow(accessToken string, id string) (*resp.StatusesShowRsp, error) {
	body, err := c.http.Get(c.host+"statuses/show", map[string]string{
		"source":       c.appKey,
		"access_token": accessToken,
		"id":           id,
	})
	if err != nil {
		return nil, err
	}
	var rsp *resp.StatusesShowRsp

	bodyStr, _ := body.ToString()
	err = json.Unmarshal([]byte(bodyStr), &rsp)
	if err != nil || strings.Index(bodyStr, `"error":`) != -1 {
		return nil, errors.New(bodyStr)
	}
	return rsp, nil
}

func (c *Client) StatusesPublicLine(accessToken string, count int, page int) (*resp.StatusesPublicLineRsp, error) {
	body, err := c.http.Get(c.host+"2/statuses/public_timeline.json", map[string]string{
		"source":       c.appKey,
		"access_token": accessToken,
		"count":        strconv.Itoa(count),
		"page":         strconv.Itoa(page),
	})
	if err != nil {
		return nil, err
	}
	var rsp *resp.StatusesPublicLineRsp

	bodyStr, _ := body.ToString()
	err = json.Unmarshal([]byte(bodyStr), &rsp)
	if err != nil || strings.Index(bodyStr, `"error":`) != -1 {
		return nil, errors.New(bodyStr)
	}
	return rsp, nil
}

func (c *Client) CommentsShow(accessToken string, id string) (*resp.CommentsShowRsp, error) {
	body, err := c.http.Get(c.host+"2/comments/show.json", map[string]string{
		"source":       c.appKey,
		"access_token": accessToken,
		"id":           id,
	})
	if err != nil {
		return nil, err
	}
	var rsp *resp.CommentsShowRsp

	bodyStr, _ := body.ToString()
	err = json.Unmarshal([]byte(bodyStr), &rsp)
	if err != nil || strings.Index(bodyStr, `"error":`) != -1 {
		return nil, errors.New(bodyStr)
	}
	return rsp, nil
}
