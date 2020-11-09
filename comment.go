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

type comment struct {
	douyin
}

// 评论列表
// 需要申请权限，需要用户授权，该接口用于获取评论列表。
func (cm *comment) List(openId, accessToken, itemId string, cursor, count int64) (*CommentList, error) {
	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("item_id", itemId)
	params.Add("cursor", fmt.Sprintf("%d", cursor))
	params.Add("count", fmt.Sprintf("%d", count))

	urlStr := fmt.Sprintf("%s/item/comment/list/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "item-comment-list:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "item-comment-list:http-err", err.Error())
		return nil, err
	}

	cl := &CommentList{}
	if err := json.Unmarshal(buf, cl); err != nil {
		goo.Log.Error(tag, "item-comment-list:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "item-comment-list:result", cl)

	if cl.Data.ErrorCode != 0 {
		return nil, errors.New(cl.Data.Description)
	}

	return cl, nil
}

// 评论回复列表
// 需要申请权限，需要用户授权，该接口用于获取评论回复列表。
// 注意参数中item_id作为url参数时，必须encode，只对item_id单独进行encode
// 注意参数中comment_id作为url参数时，必须encode，只对comment_id单独进行encode
func (cm *comment) ReplyList(openId, accessToken, itemId, commentId string, cursor, count int64) (*CommentList, error) {
	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("item_id", itemId)
	params.Add("comment_id", commentId)
	params.Add("cursor", fmt.Sprintf("%d", cursor))
	params.Add("count", fmt.Sprintf("%d", count))

	urlStr := fmt.Sprintf("%s/item/comment/reply/list/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "item-comment-reply-list:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "item-comment-reply-list:http-err", err.Error())
		return nil, err
	}

	cl := &CommentList{}
	if err := json.Unmarshal(buf, cl); err != nil {
		goo.Log.Error(tag, "item-comment-reply-list:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "item-comment-reply-list:result", cl)

	if cl.Data.ErrorCode != 0 {
		return nil, errors.New(cl.Data.Description)
	}

	return cl, nil
}

// 回复视频评论
// 需要申请权限，需要用户授权，该接口用于回复视频评论。
func (cm *comment) Reply(openId, accessToken, itemId, commentId, content string) (*CommentReply, error) {
	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权

	urlStr := fmt.Sprintf("%s/item/comment/reply/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "item-comment-reply:url", urlStr)

	data := map[string]interface{}{
		"item_id":    itemId,
		"comment_id": commentId,
		"content":    content,
	}
	body, _ := json.Marshal(data)

	buf, err := goo.PostJson(urlStr, body)
	if err != nil {
		goo.Log.Error(tag, "item-comment-reply:http-err", err.Error())
		return nil, err
	}

	cr := &CommentReply{}
	if err := json.Unmarshal(buf, cr); err != nil {
		goo.Log.Error(tag, "item-comment-reply:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "item-comment-reply:result", cr)

	if cr.Data.ErrorCode != 0 {
		return nil, errors.New(cr.Data.Description)
	}

	return cr, nil
}

// 评论列表 - 企业号
// 需要申请权限，需要用户授权，该接口用于查看指定视频的实时评论列表。
// 调用本接口，需要授权的抖音用户是企业号企业号 。
func (cm *comment) VideoList(openId, accessToken, itemId string, cursor, count int64) (*CommentList, error) {
	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("item_id", itemId)
	params.Add("cursor", fmt.Sprintf("%d", cursor))
	params.Add("count", fmt.Sprintf("%d", count))

	urlStr := fmt.Sprintf("%s/video/comment/list/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "video-comment-list:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "video-comment-list:http-err", err.Error())
		return nil, err
	}

	cl := &CommentList{}
	if err := json.Unmarshal(buf, cl); err != nil {
		goo.Log.Error(tag, "video-comment-list:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "video-comment-list:result", cl)

	if cl.Data.ErrorCode != 0 {
		return nil, errors.New(cl.Data.Description)
	}

	return cl, nil
}

// 评论回复列表 - 企业号
// 需要申请权限，需要用户授权，该接口用于查看指定视频的评论回复列表。
// 调用本接口，需要授权的抖音用户是企业号企业号 。
// 注意参数中item_id作为url参数时，必须encode，只对item_id单独进行encode
// 注意参数中comment_id作为url参数时，必须encode，只对comment_id单独进行encode
func (cm *comment) VideoReplyList(openId, accessToken, itemId, commentId string, cursor, count int64) (*CommentList, error) {
	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("item_id", itemId)
	params.Add("comment_id", commentId)
	params.Add("cursor", fmt.Sprintf("%d", cursor))
	params.Add("count", fmt.Sprintf("%d", count))

	urlStr := fmt.Sprintf("%s/video/comment/reply/list/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "video-comment-reply-list:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "video-comment-reply-list:http-err", err.Error())
		return nil, err
	}

	cl := &CommentList{}
	if err := json.Unmarshal(buf, cl); err != nil {
		goo.Log.Error(tag, "video-comment-reply-list:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "video-comment-reply-list:result", cl)

	if cl.Data.ErrorCode != 0 {
		return nil, errors.New(cl.Data.Description)
	}

	return cl, nil
}

// 回复视频评论 - 企业号
// 需要申请权限，需要用户授权，该接口用于回复视频评论。
// 调用本接口，需要授权的抖音用户是企业号企业号 。
func (cm *comment) VideoReply(openId, accessToken, itemId, commentId, content string) (*CommentReply, error) {
	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权

	urlStr := fmt.Sprintf("%s/video/comment/reply/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "video-comment-reply:url", urlStr)

	data := map[string]interface{}{
		"item_id":    itemId,
		"comment_id": commentId,
		"content":    content,
	}
	body, _ := json.Marshal(data)

	buf, err := goo.PostJson(urlStr, body)
	if err != nil {
		goo.Log.Error(tag, "video-comment-reply:http-err", err.Error())
		return nil, err
	}

	cr := &CommentReply{}
	if err := json.Unmarshal(buf, cr); err != nil {
		goo.Log.Error(tag, "video-comment-reply:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "video-comment-reply:result", cr)

	if cr.Data.ErrorCode != 0 {
		return nil, errors.New(cr.Data.Description)
	}

	return cr, nil
}

// 置顶视频评论 - 企业号
// 需要申请权限，需要用户授权，该接口用于置顶视频评论。
// 调用本接口，需要授权的抖音用户是企业号企业号 。
func (cm *comment) VideoTop(openId, accessToken, itemId, commentId string, top bool) (*Result, error) {
	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权

	urlStr := fmt.Sprintf("%s/video/comment/top/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "video-comment-top:url", urlStr)

	data := map[string]interface{}{
		"item_id":    itemId,
		"comment_id": commentId,
		"top":        top,
	}
	body, _ := json.Marshal(data)

	buf, err := goo.PostJson(urlStr, body)
	if err != nil {
		goo.Log.Error(tag, "video-comment-top:http-err", err.Error())
		return nil, err
	}

	rst := &Result{}
	if err := json.Unmarshal(buf, rst); err != nil {
		goo.Log.Error(tag, "video-comment-top:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "video-comment-top:result", rst)

	if rst.Data.ErrorCode != 0 {
		return nil, errors.New(rst.Data.Description)
	}

	return rst, nil
}
