package goo_open_douyin

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/liqiongtao/goo"
	"net/url"
)

// --------------------------------------------
// -- 抖音开放平台 - 用户
// --------------------------------------------

type user struct {
	douyin
}

// 获取用户信息
// 需要用户授权
func (u *user) userInfo(openId, accessToken, baseUrl string) (*UserInfo, error) {
	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权

	urlStr := fmt.Sprintf("%s/oauth/userinfo/?%s", baseUrl, params.Encode())
	goo.Log.Debug(tag, "oauth-userinfo:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "oauth-userinfo:http-err", err.Error())
		return nil, err
	}

	ui := &UserInfo{}
	if err := json.Unmarshal(buf, ui); err != nil {
		goo.Log.Error(tag, "oauth-userinfo:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "oauth-userinfo:result", ui)

	if ui.Data.ErrorCode != 0 {
		return nil, errors.New(ui.Data.Description)
	}

	return ui, nil
}

// 获取用户信息 - 抖音
func (u *user) UserInfoDouYin(openId, accessToken string) (*UserInfo, error) {
	return u.userInfo(openId, accessToken, base_url_douyin)
}

// 获取用户信息 - 头条
func (u *user) UserInfoSnsSdk(openId, accessToken string) (*UserInfo, error) {
	return u.userInfo(openId, accessToken, base_url_snssdk)
}

// 获取用户信息 - 西瓜
func (u *user) UserInfoXiGua(openId, accessToken string) (*UserInfo, error) {
	return u.userInfo(openId, accessToken, base_url_xigua)
}

// 获取粉丝列表
// 需要用户授权，获取用户最近的粉丝列表，不保证顺序，用于发布视频时@用户。目前可查询的粉丝数上限5千。该接口适用于抖音。
func (u *user) FansList(openId, accessToken string, cursor, count int64) (*FansList, error) {
	params := url.Values{}
	params.Add("open_id", openId)                   // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken)         // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("cursor", fmt.Sprintf("%d", cursor)) // 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据
	params.Add("count", fmt.Sprintf("%d", count))   // 每页数量

	urlStr := fmt.Sprintf("%s/fans/list/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "fans-list:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "fans-list:http-err", err.Error())
		return nil, err
	}

	fl := &FansList{}
	if err := json.Unmarshal(buf, fl); err != nil {
		goo.Log.Error(tag, "fans-list:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "fans-list:result", fl)

	if fl.Data.ErrorCode != 0 {
		return nil, errors.New(fl.Data.Description)
	}

	return fl, nil
}

// 获取关注列表
// 需要用户授权，获取用户的关注列表；该接口适用于抖音。
func (u *user) FollowingList(openId, accessToken string, cursor, count int64) (*FollowingList, error) {
	params := url.Values{}
	params.Add("open_id", openId)                   // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken)         // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("cursor", fmt.Sprintf("%d", cursor)) // 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据
	params.Add("count", fmt.Sprintf("%d", count))   // 每页数量

	urlStr := fmt.Sprintf("%s/following/list/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "following-list:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "following-list:http-err", err.Error())
		return nil, err
	}

	fl := &FollowingList{}
	if err := json.Unmarshal(buf, fl); err != nil {
		goo.Log.Error(tag, "following-list:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "following-list:result", fl)

	if fl.Data.ErrorCode != 0 {
		return nil, errors.New(fl.Data.Description)
	}

	return fl, nil
}
