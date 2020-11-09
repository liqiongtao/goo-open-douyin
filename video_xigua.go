package goo_open_douyin

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/liqiongtao/goo"
	"io"
	"net/url"
)

// --------------------------------------------
// -- 抖音开放平台 - 视频 - 西瓜
// --------------------------------------------

type videoXiGua struct {
	douyin
}

// 上传视频到文件服务器
// 该接口用于上传视频文件到文件服务器，获取视频文件video_id。该接口适用于头条。
// 超过50m的视频建议采用分片上传，可以降低网关超时造成的失败。超过128m的视频必须采用分片上传。视频总大小4GB以内。单个分片建议20MB，最小5MB。
// 支持常用视频格式，推荐使用 mp4 、webm
// 视频文件大小不超过128M，时长在15分钟以内
func (v *videoXiGua) Upload(openId, accessToken, filename string, f io.Reader) (*VideoUpload, error) {
	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权

	urlStr := fmt.Sprintf("%s/xihua/video/upload/?%s", base_url_douyin, params.Encode())

	buf, err := goo.Upload(urlStr, "video", filename, f, nil)
	if err != nil {
		goo.Log.Error(tag, "xihua-video-upload:http-err", err.Error())
		return nil, err
	}

	vu := &VideoUpload{}
	if err := json.Unmarshal(buf, vu); err != nil {
		goo.Log.Error(tag, "xihua-video-upload:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "xihua-video-upload:result", vu)

	if vu.Data.ErrorCode != 0 {
		return nil, errors.New(vu.Data.Description + ":" + vu.Extra.SubDescription)
	}

	return vu, nil
}

// 分片初始化上传
// 该接口用于分片上传视频文件到文件服务器，先初始化上传获取upload_id。该接口适用于头条。
// 超过50m的视频建议采用分片上传，可以降低网关超时造成的失败。超过128m的视频必须采用分片上传。视频总大小4GB以内。单个分片建议20MB，最小5MB。
func (v *videoXiGua) PartInit(openId, accessToken string) (*VideoPartInit, error) {
	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权

	urlStr := fmt.Sprintf("%s/xihua/video/part/init/?%s", base_url_douyin, params.Encode())

	buf, err := goo.PostJson(urlStr, []byte{})
	if err != nil {
		goo.Log.Error(tag, "xihua-video-part-init:http-err", err.Error())
		return nil, err
	}

	vpi := &VideoPartInit{}
	if err := json.Unmarshal(buf, vpi); err != nil {
		goo.Log.Error(tag, "xihua-video-part-init:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "xihua-video-part-init:result", vpi)

	if vpi.Data.ErrorCode != 0 {
		return nil, errors.New(vpi.Data.Description + ":" + vpi.Extra.SubDescription)
	}

	return vpi, nil
}

// 上传视频分片到文件服务器
// 该接口用于分片上传视频文件到文件服务器，先初始化上传获取upload_id。该接口适用于头条。
// 超过50m的视频建议采用分片上传，可以降低网关超时造成的失败。超过128m的视频必须采用分片上传。视频总大小300MB以内。单个分片建议20MB，最小5MB。
// 注意参数中upload_id作为url参数时，必须encode，只对upload_id进行encode
func (v *videoXiGua) PartUpload(openId, accessToken, filename string, f io.Reader, uploadId string, partNumber int64) (*Result, error) {
	params := url.Values{}
	params.Add("open_id", openId)                            // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken)                  // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("upload_id", uploadId)                        // 分片上传的标记。有限时间为2小时。
	params.Add("part_number", fmt.Sprintf("%d", partNumber)) // 第几个分片，从1开始

	urlStr := fmt.Sprintf("%s/xihua/video/part/upload/?%s", base_url_douyin, params.Encode())

	buf, err := goo.Upload(urlStr, "video", filename, f, nil)
	if err != nil {
		goo.Log.Error(tag, "xihua-video-part-upload:http-err", err.Error())
		return nil, err
	}

	rst := &Result{}
	if err := json.Unmarshal(buf, rst); err != nil {
		goo.Log.Error(tag, "xihua-video-part-upload:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "xihua-video-part-upload:result", rst)

	if rst.Data.ErrorCode != 0 {
		return nil, errors.New(rst.Data.Description + ":" + rst.Extra.SubDescription)
	}

	return rst, nil
}

// 分片完成上传
// 该接口用于分片上传视频文件到文件服务器，先初始化上传获取upload_id。该接口适用于头条。
// 超过50m的视频建议采用分片上传，可以降低网关超时造成的失败。超过128m的视频必须采用分片上传。视频总大小300MB以内。单个分片建议20MB，最小5MB。
func (v *videoXiGua) PartComplete(openId, accessToken, uploadId string) (*VideoUpload, error) {
	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("upload_id", uploadId)       // 分片上传的标记。有限时间为2小时。

	urlStr := fmt.Sprintf("%s/xihua/video/part/complete/?%s", base_url_douyin, params.Encode())

	buf, err := goo.PostJson(urlStr, []byte{})
	if err != nil {
		goo.Log.Error(tag, "xihua-video-part-complete:http-err", err.Error())
		return nil, err
	}

	vu := &VideoUpload{}
	if err := json.Unmarshal(buf, vu); err != nil {
		goo.Log.Error(tag, "xihua-video-part-complete:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "xihua-video-part-complete:result", vu)

	if vu.Data.ErrorCode != 0 {
		return nil, errors.New(vu.Data.Description + ":" + vu.Extra.SubDescription)
	}

	return vu, nil
}

// 创建视频
// 该接口用于发布视频到头条。该接口适用于头条。
func (v *videoXiGua) Create(openId, accessToken, videoId, text string, atUsers []string) (*VideoCreate, error) {
	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权

	urlStr := fmt.Sprintf("%s/xihua/video/create/?%s", base_url_douyin, params.Encode())

	data := map[string]interface{}{
		"text":     text,
		"video_id": videoId,
	}
	if atUsers != nil {
		data["at_users"] = atUsers
	}
	body, _ := json.Marshal(data)

	buf, err := goo.PostJson(urlStr, body)
	if err != nil {
		goo.Log.Error(tag, "xihua-video-create:http-err", err.Error())
		return nil, err
	}

	vc := &VideoCreate{}
	if err := json.Unmarshal(buf, vc); err != nil {
		goo.Log.Error(tag, "xihua-video-create:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "xihua-video-create:result", vc)

	if vc.Data.ErrorCode != 0 {
		return nil, errors.New(vc.Data.Description + ":" + vc.Extra.SubDescription)
	}

	return vc, nil
}

// 查询授权帐号的视频列表
// 该接口用于分页获取用户所有视频的数据；适用于头条。
func (v *videoXiGua) List(openId, accessToken string, cursor, count int64) (*VideoList, error) {
	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("cursor", fmt.Sprintf("%d", cursor))
	params.Add("count", fmt.Sprintf("%d", count))

	urlStr := fmt.Sprintf("%s/xihua/video/list/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "xihua-video-list:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "xihua-video-list:http-err", err.Error())
		return nil, err
	}

	vl := &VideoList{}
	if err := json.Unmarshal(buf, vl); err != nil {
		goo.Log.Error(tag, "xihua-video-list:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "xihua-video-list:result", vl)

	if vl.Data.ErrorCode != 0 {
		return nil, errors.New(vl.Data.Description + ":" + vl.Extra.SubDescription)
	}

	return vl, nil
}

// 查询特定视频的视频信息
// 该接口用于查询用户特定视频的数据, 如点赞数, 播放数等；适用于头条。
func (v *videoXiGua) Data(openId, accessToken string, itemIds []string) (*VideoData, error) {
	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权

	urlStr := fmt.Sprintf("%s/xihua/video/data/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "xihua-video-data:url", urlStr)

	bd, _ := json.Marshal(map[string]interface{}{"item_ids": itemIds})
	buf, err := goo.PostJson(urlStr, bd)
	if err != nil {
		goo.Log.Error(tag, "xihua-video-data:http-err", err.Error())
		return nil, err
	}

	vd := &VideoData{}
	if err := json.Unmarshal(buf, vd); err != nil {
		goo.Log.Error(tag, "xihua-video-data:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "xihua-video-data:result", vd)

	if vd.Data.ErrorCode != 0 {
		return nil, errors.New(vd.Data.Description + ":" + vd.Extra.SubDescription)
	}

	return vd, nil
}
