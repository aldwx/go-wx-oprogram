package auth

import (
	httpclient "github.com/aldwx/go-http-client"
	"github.com/aldwx/go-wx-oprogram/common"
)

const (
	apiLogin          = "/sns/component/jscode2session"
	apiGetAccessToken = "/cgi-bin/component/api_component_token"
)

type Auth struct{}

type LoginResponse struct {
	common.CommonError
	OpenId     string `json:"openid"`      // 第三方平台 access_token
	SessionKey string `json:"session_key"` // 有效期，单位：秒
}

// jscode2session 小程序登录
// appid					小程序的 AppID
// js_code					wx.login 获取的 code
// grant_type				填 authorization_code
// component_appid			第三方平台 appid
// component_access_token	令牌
func (a *Auth) Login(appId, jsCode, componentAppId, componentAccessToken string) (*LoginResponse, error) {
	api := common.BaseURL + apiLogin
	return jsCode2session(api, appId, jsCode, "authorization_code", componentAppId, componentAccessToken)
}

func jsCode2session(api, appId, jsCode, grantType, componentAppId, componentAccessToken string) (*LoginResponse, error) {
	url, err := httpclient.EncodeURL(api, httpclient.RequestQueries{
		"appid":                  appId,
		"js_code":                jsCode,
		"grant_type":             grantType,
		"component_appid":        componentAppId,
		"component_access_token": componentAccessToken,
	})
	if err != nil {
		return nil, err
	}

	res := new(LoginResponse)
	if err := httpclient.GetJSON(url, res); err != nil {
		return nil, err
	}
	return res, nil
}

// TokenResponse 获取 access_token 成功返回数据
type TokenResponse struct {
	common.CommonError
	AccessToken string `json:"component_access_token"` // 获取到的凭证
	ExpiresIn   uint   `json:"expires_in"`             // 凭证有效时间，单位：秒。目前是7200秒之内的值。
}

// GetAccessToken 获取小程序全局唯一后台接口调用凭据（access_token）。
// 调调用绝大多数后台接口时都需使用 access_token，开发者需要进行妥善保存，注意缓存。
func (a *Auth) GetAccessToken(componentAppID, componentSecret, ticket string) (*TokenResponse, error) {
	api := common.BaseURL + apiGetAccessToken
	return getAccessToken(componentAppID, componentSecret, ticket, api)
}

func getAccessToken(appID, secret, ticket, api string) (*TokenResponse, error) {
	queries := httpclient.RequestQueries{
		"component_appid":         appID,
		"component_appsecret":     secret,
		"component_verify_ticket": ticket,
	}
	res := new(TokenResponse)
	if err := httpclient.PostJSON(api, queries, res); err != nil {
		return nil, err
	}

	return res, nil
}
