package goo_open_douyin

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/liqiongtao/goo"
	"net/url"
)

// --------------------------------------------
// -- 抖音开放平台 - 互动 - 评论
// --------------------------------------------

type search struct {
	douyin
}

// 关键词视频搜索
// 需要申请权限，需要用户授权
// 该接口用于通过关键词搜索全站视频,类似抖音端上搜索。使用前请到 管理中心-应用详情-关键词视频管理-关键词管理 创建关键词。
// 该接口只返回最近1天的视频
func (se *search) List(openId, accessToken, keyword string, cursor, count int64) (*CommentList, error) {
	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("keyword", keyword)
	params.Add("cursor", fmt.Sprintf("%d", cursor))
	params.Add("count", fmt.Sprintf("%d", count))

	urlStr := fmt.Sprintf("%s/video/search/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "video-search:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "video-search:http-err", err.Error())
		return nil, err
	}

	cl := &CommentList{}
	if err := json.Unmarshal(buf, cl); err != nil {
		goo.Log.Error(tag, "video-search:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "video-search:result", cl)

	if cl.Data.ErrorCode != 0 {
		return nil, errors.New(cl.Data.Description)
	}

	return cl, nil
}

// 关键词视频评论管理 - 评论列表
func (se *search) CommentList(accessToken, secItemId string, cursor, count int64) (*CommentList, error) {
	params := url.Values{}
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("cursor", fmt.Sprintf("%d", cursor))
	params.Add("count", fmt.Sprintf("%d", count))
	params.Add("sec_item_id", secItemId) // 视频搜索接口返回的加密的视频id

	urlStr := fmt.Sprintf("%s/video/search/comment/list/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "video-search-comment-list:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "video-search-comment-list:http-err", err.Error())
		return nil, err
	}

	cl := &CommentList{}
	if err := json.Unmarshal(buf, cl); err != nil {
		goo.Log.Error(tag, "video-search-comment-list:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "video-search-comment-list:result", cl)

	if cl.Data.ErrorCode != 0 {
		return nil, errors.New(cl.Data.Description)
	}

	return cl, nil
}

// 关键词视频评论管理 - 回复视频评论
func (se *search) CommentReply(openId, accessToken, secItemId, commentId, content string) (*CommentReply, error) {
	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权

	urlStr := fmt.Sprintf("%s/video/search/comment/reply/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "video-search-comment-list:url", urlStr)

	data := map[string]interface{}{
		"sec_item_id": secItemId, // 视频搜索接口返回的加密的视频id
		"comment_id":  commentId, // 需要回复的评论id（如果需要回复的是视频不传此字段）
		"content":     content,   // 评论内容
	}
	body, _ := json.Marshal(data)

	buf, err := goo.PostJson(urlStr, body)
	if err != nil {
		goo.Log.Error(tag, "video-search-comment-list:http-err", err.Error())
		return nil, err
	}

	cr := &CommentReply{}
	if err := json.Unmarshal(buf, cr); err != nil {
		goo.Log.Error(tag, "video-search-comment-list:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "video-search-comment-list:result", cr)

	if cr.Data.ErrorCode != 0 {
		return nil, errors.New(cr.Data.Description)
	}

	return cr, nil
}

// 关键词视频评论管理 - 评论回复列表
func (se *search) CommentReplyList(accessToken, secItemId, commentId string, cursor, count int64) (*CommentList, error) {
	params := url.Values{}
	params.Add("access_token", accessToken)         // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("cursor", fmt.Sprintf("%d", cursor)) // 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。
	params.Add("count", fmt.Sprintf("%d", count))   // 每页数量
	params.Add("sec_item_id", secItemId)            // 视频搜索接口返回的加密的视频id
	params.Add("comment_id", commentId)             // 评论id

	urlStr := fmt.Sprintf("%s/video/search/comment/reply/list/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "video-search-comment-reply-list:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "video-search-comment-reply-list:http-err", err.Error())
		return nil, err
	}

	cl := &CommentList{}
	if err := json.Unmarshal(buf, cl); err != nil {
		goo.Log.Error(tag, "video-search-comment-reply-list:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "video-search-comment-reply-list:result", cl)

	if cl.Data.ErrorCode != 0 {
		return nil, errors.New(cl.Data.Description)
	}

	return cl, nil
}
