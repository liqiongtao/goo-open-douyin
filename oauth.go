package goo_open_douyin

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/liqiongtao/goo"
	"github.com/liqiongtao/goo/utils"
	"net/url"
)

// --------------------------------------------
// -- 抖音开放平台 - 账号授权
// --------------------------------------------

type oauth struct {
	douyin
}

// 抖音获取授权码(code)
// 该接口只适用于抖音获取授权临时票据（code）
func (o *oauth) platformConnect(scope, optionalScope, redirectUri, baseUrl string) string {
	params := url.Values{}
	params.Add("client_key", o.clientKey)
	params.Add("response_type", "code")
	params.Add("scope", scope)                 // 应用授权作用域,多个授权作用域以英文逗号（,）分隔
	params.Add("optionalScope", optionalScope) // 应用授权可选作用域,多个授权作用域以英文逗号（,）分隔，每一个授权作用域后需要加上一个是否默认勾选的参数，1为默认勾选，0为默认不勾选
	params.Add("redirect_uri", redirectUri)
	params.Add("state", utils.NonceStr())

	urlStr := fmt.Sprintf("%s/platform/oauth/connect/?%s", baseUrl, params.Encode())
	goo.Log.Debug(tag, "platform-connect:url", urlStr)

	return urlStr
}

// 抖音获取授权码(code) - 抖音
func (o *oauth) PlatformConnect(scope, optionalScope, redirectUri string) string {
	return o.platformConnect(scope, optionalScope, redirectUri, base_url_douyin)
}

// 抖音获取授权码(code) - 西瓜
func (o *oauth) PlatformConnectXiGua(scope, optionalScope, redirectUri string) string {
	return o.platformConnect(scope, optionalScope, redirectUri, base_url_xigua)
}

func (o *oauth) Qrcode(scope, redirectUrl string) (*Qrcode, string, error) {
	state := utils.NonceStr()

	params := url.Values{}
	params.Add("client_key", o.clientKey)
	params.Add("scope", scope) // 应用授权作用域,多个授权作用域以英文逗号（,）分隔
	params.Add("next", redirectUrl)
	params.Add("state", state)

	urlStr := fmt.Sprintf("%s/oauth/get_qrcode/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "oauth-get-qrcode:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "oauth-get-qrcode:http-err", err.Error())
		return nil, "", err
	}

	qr := &Qrcode{}
	if err := json.Unmarshal(buf, qr); err != nil {
		goo.Log.Error(tag, "oauth-get-qrcode:err", err.Error())
		return nil, "", err
	}

	goo.Log.Debug(tag, "oauth-get-qrcode:result", qr)

	if qr.Message != "success" {
		return nil, "", errors.New(qr.Data.Description)
	}

	qr.Data.Qrcode = fmt.Sprintf("data:image/png;base64,%s", qr.Data.Qrcode)

	return qr, state, nil
}

func (o *oauth) CheckQrcode(scope, redirectUrl, state, token string) (*CheckQrcode, error) {
	params := url.Values{}
	params.Add("client_key", o.clientKey)
	params.Add("scope", scope) // 应用授权作用域,多个授权作用域以英文逗号（,）分隔
	params.Add("next", redirectUrl)
	params.Add("state", state)
	params.Add("token", token)

	urlStr := fmt.Sprintf("%s/oauth/check_qrcode/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "oauth-check-qrcode:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "oauth-check-qrcode:http-err", err.Error())
		return nil, err
	}

	cq := &CheckQrcode{}
	if err := json.Unmarshal(buf, cq); err != nil {
		goo.Log.Error(tag, "oauth-check-qrcode:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "oauth-check-qrcode:result", cq)

	if cq.Data.ErrorCode != 0 {
		return nil, errors.New(cq.Data.Description)
	}

	return cq, nil
}

// 头条获取授权码(code)
// 该接口只适用于头条获取授权临时票据（code）
func (o *oauth) AuthorizeSnssdk(scope, redirectUri string) string {
	params := url.Values{}
	params.Add("client_key", o.clientKey)
	params.Add("response_type", "code")
	params.Add("scope", scope) // 应用授权作用域,多个授权作用域以英文逗号（,）分隔
	params.Add("redirect_uri", redirectUri)
	params.Add("state", utils.NonceStr())

	urlStr := fmt.Sprintf("%s/oauth/authorize/?%s", base_url_snssdk, params.Encode())
	goo.Log.Debug(tag, "authorize-snssdk:url", urlStr)

	return urlStr
}

// 获取access_token
// 该接口用于获取用户授权第三方接口调用的凭证access_token；该接口适用于抖音/头条授权。
func (o *oauth) accessToken(code, baseUrl string) (*AccessToken, error) {
	params := url.Values{}
	params.Add("client_key", o.clientKey)
	params.Add("client_secret", o.clientSecret)
	params.Add("code", code)
	params.Add("grant_type", "authorization_code")

	urlStr := fmt.Sprintf("%s/oauth/access_token/?%s", baseUrl, params.Encode())
	goo.Log.Debug(tag, "oauth-access-token:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "oauth-access-token:http-err", err.Error())
		return nil, err
	}

	at := &AccessToken{}
	if err := json.Unmarshal(buf, at); err != nil {
		goo.Log.Error(tag, "oauth-access-token:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "oauth-access-token:result", at)

	if at.Data.ErrorCode != 0 {
		return nil, errors.New(at.Data.Description)
	}

	return at, nil
}

// 获取access_token - 抖音
func (o *oauth) AccessTokenDouYin(code string) (*AccessToken, error) {
	return o.accessToken(code, base_url_douyin)
}

// 获取access_token - 头条
func (o *oauth) AccessTokenSnsSdk(code string) (*AccessToken, error) {
	return o.accessToken(code, base_url_snssdk)
}

// 获取access_token - 西瓜
func (o *oauth) AccessTokenXiGua(code string) (*AccessToken, error) {
	return o.accessToken(code, base_url_xigua)
}

// 刷新refresh_token
// 不需要授权，该接口用于刷新refresh_token的有效期；该接口适用于抖音授权。
// 通过旧的refresh_token获取新的refresh_token，调用后旧refresh_token会失效，新refresh_token有30天有效期。
// 最多只能获取5次新的refresh_token，5次过后需要用户重新授权
func (o *oauth) RenewRefreshToken(refreshToken string) (*RenewRefreshToken, error) {
	params := url.Values{}
	params.Add("client_key", o.clientKey)
	params.Add("refresh_token", refreshToken)

	urlStr := fmt.Sprintf("%s/oauth/renew_refresh_token/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug("[pen-douyin", "oauth-renew-refresh-token:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "oauth-renew-refresh-token:http-err", err.Error())
		return nil, err
	}

	rt := &RenewRefreshToken{}
	if err := json.Unmarshal(buf, rt); err != nil {
		goo.Log.Error(tag, "oauth-renew-refresh-token:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "oauth-renew-refresh-token:result", rt)

	if rt.Data.ErrorCode != 0 {
		return nil, errors.New(rt.Data.Description)
	}

	return rt, nil
}

// 生成client_token
// 该接口用于获取接口调用的凭证client_access_token，主要用于调用不需要用户授权就可以调用的接口；该接口适用于抖音/头条授权
// client_access_token的有效时间为2个小时，重复获取token后会使上次的token失效(但有5分钟的缓冲时间)
func (o *oauth) clientToken(baseUrl string) (*ClientToken, error) {
	params := url.Values{}
	params.Add("client_key", o.clientKey)
	params.Add("client_secret", o.clientSecret)
	params.Add("grant_type", "client_credential")

	urlStr := fmt.Sprintf("%s/oauth/client_token/?%s", baseUrl, params.Encode())
	goo.Log.Debug(tag, "client-token:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "client-token:http-err", err.Error())
		return nil, err
	}

	ct := &ClientToken{}
	if err := json.Unmarshal(buf, ct); err != nil {
		goo.Log.Error(tag, "client-token:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "client-token:result", ct)

	if ct.Data.ErrorCode != 0 {
		return nil, errors.New(ct.Data.Description)
	}

	return ct, nil
}

// 生成client_token - 抖音
func (o *oauth) ClientTokenDouYin(baseUrl string) (*ClientToken, error) {
	return o.clientToken(base_url_douyin)
}

// 生成client_token - 头条
func (o *oauth) ClientTokenSnsSdk(baseUrl string) (*ClientToken, error) {
	return o.clientToken(base_url_snssdk)
}

// 生成client_token - 西瓜
func (o *oauth) ClientTokenXiGua(baseUrl string) (*ClientToken, error) {
	return o.clientToken(base_url_xigua)
}

// 刷新access_token
// 该接口用于刷新access_token的有效期；该接口适用于抖音/头条授权。
// 当access_token过期（过期时间15天）后，可以通过该接口使用refresh_token（过期时间30天）进行刷新
func (o *oauth) refreshAccessToken(refreshToken, baseUrl string) (*RefreshToken, error) {
	params := url.Values{}
	params.Add("client_key", o.clientKey)
	params.Add("grant_type", "refresh_token")
	params.Add("refresh_token", refreshToken)

	urlStr := fmt.Sprintf("%s/oauth/refresh_token/?%s", baseUrl, params.Encode())
	goo.Log.Debug(tag, "refresh-access-token:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "refresh-access-token:http-err", err.Error())
		return nil, err
	}

	rt := &RefreshToken{}
	if err := json.Unmarshal(buf, rt); err != nil {
		goo.Log.Error(tag, "refresh-access-token:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "refresh-access-token:result", rt)

	if rt.Data.ErrorCode != 0 {
		return nil, errors.New(rt.Data.Description)
	}

	return rt, nil
}

// 刷新access_token - 抖音
func (o *oauth) RefreshAccessTokenDouYin(refreshToken string) (*RefreshToken, error) {
	return o.refreshAccessToken(refreshToken, base_url_douyin)
}

// 刷新access_token - 头条
func (o *oauth) RefreshAccessTokenSnsSdk(refreshToken string) (*RefreshToken, error) {
	return o.refreshAccessToken(refreshToken, base_url_snssdk)
}

// 刷新access_token - 西瓜
func (o *oauth) RefreshAccessTokenXiGua(refreshToken string) (*RefreshToken, error) {
	return o.refreshAccessToken(refreshToken, base_url_xigua)
}

// 获取授权码(code)
// 不需要授权，该接口适用于抖音获取静默授权临时票据（code）
func (o *oauth) AuthorizeV2(redirectUri string) string {
	params := url.Values{}
	params.Add("client_key", o.clientKey)
	params.Add("response_type", "code")     // 填写code
	params.Add("scope", "login_id")         // 填login_id
	params.Add("redirect_uri", redirectUri) // 授权成功后的回调地址，必须以http/https开头。域名要跟申请应用时填写的授权回调域一致。用于调用https://open.douyin.com/oauth/access_token/换token。
	params.Add("state", utils.NonceStr())   // 用于保持请求和回调状态，授权请求后会原样返回给接入方,如果是App则不用传该参数

	urlStr := fmt.Sprintf("%s/oauth/authorize/v2/?%s", base_url_aweme, params.Encode())
	goo.Log.Debug(tag, "oauth-authorize-v2:url", urlStr)
	return urlStr
}
