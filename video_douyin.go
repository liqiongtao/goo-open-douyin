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
// -- 抖音开放平台 - 视频 - 抖音
// --------------------------------------------

type videoDouYin struct {
	douyin
}

// 上传视频到文件服务器
// 需要申请权限，需要用户授权，该接口用于上传视频文件到文件服务器，获取视频文件video_id。该接口适用于抖音。
// 超过50m的视频建议采用分片上传，可以降低网关超时造成的失败。超过128m的视频必须采用分片上传。视频总大小4GB以内。单个分片建议20MB，最小5MB。
// 支持常用视频格式，推荐使用 mp4 、webm
// 视频文件大小不超过128M，时长在15分钟以内
func (v *videoDouYin) Upload(openId, accessToken, filename string, f io.Reader) (*VideoUpload, error) {
	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权

	urlStr := fmt.Sprintf("%s/video/upload/?%s", base_url_douyin, params.Encode())

	buf, err := goo.Upload(urlStr, "video", filename, f, nil)
	if err != nil {
		goo.Log.Error(tag, "video-upload:http-err", err.Error())
		return nil, err
	}

	vu := &VideoUpload{}
	if err := json.Unmarshal(buf, vu); err != nil {
		goo.Log.Error(tag, "video-upload:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "video-upload:result", vu)

	if vu.Data.ErrorCode != 0 {
		return nil, errors.New(vu.Data.Description + ":" + vu.Extra.SubDescription)
	}

	return vu, nil
}

// 分片初始化上传
// 该接口用于分片上传视频文件到文件服务器，先初始化上传获取upload_id。该接口适用于抖音。
// 超过50m的视频建议采用分片上传，可以降低网关超时造成的失败。超过128m的视频必须采用分片上传。视频总大小4GB以内。单个分片建议20MB，最小5MB。
func (v *videoDouYin) PartInit(openId, accessToken string) (*VideoPartInit, error) {
	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权

	urlStr := fmt.Sprintf("%s/video/part/init/?%s", base_url_douyin, params.Encode())

	buf, err := goo.PostJson(urlStr, []byte{})
	if err != nil {
		goo.Log.Error(tag, "video-part-init:http-err", err.Error())
		return nil, err
	}

	vpi := &VideoPartInit{}
	if err := json.Unmarshal(buf, vpi); err != nil {
		goo.Log.Error(tag, "video-part-init:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "video-part-init:result", vpi)

	if vpi.Data.ErrorCode != 0 {
		return nil, errors.New(vpi.Data.Description + ":" + vpi.Extra.SubDescription)
	}

	return vpi, nil
}

// 上传视频分片到文件服务器
// 该接口用于分片上传视频文件到文件服务器，上传阶段。该接口适用于抖音。
// 超过50m的视频建议采用分片上传，可以降低网关超时造成的失败。超过128m的视频必须采用分片上传。视频总大小4GB以内。单个分片建议20MB，最小5MB。
// 注意参数中upload_id作为url参数时，必须encode，只对upload_id进行encode
func (v *videoDouYin) PartUpload(openId, accessToken, filename string, f io.Reader, uploadId string, partNumber int64) (*Result, error) {
	params := url.Values{}
	params.Add("open_id", openId)                            // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken)                  // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("upload_id", uploadId)                        // 分片上传的标记。有限时间为2小时。
	params.Add("part_number", fmt.Sprintf("%d", partNumber)) // 第几个分片，从1开始

	urlStr := fmt.Sprintf("%s/video/part/upload/?%s", base_url_douyin, params.Encode())

	buf, err := goo.Upload(urlStr, "video", filename, f, nil)
	if err != nil {
		goo.Log.Error(tag, "video-part-upload:http-err", err.Error())
		return nil, err
	}

	rst := &Result{}
	if err := json.Unmarshal(buf, rst); err != nil {
		goo.Log.Error(tag, "video-part-upload:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "video-part-upload:result", rst)

	if rst.Data.ErrorCode != 0 {
		return nil, errors.New(rst.Data.Description + ":" + rst.Extra.SubDescription)
	}

	return rst, nil
}

// 分片完成上传
// 该接口用于分片上传视频文件到文件服务器，完成上传。该接口适用于抖音。
// 超过50m的视频建议采用分片上传，可以降低网关超时造成的失败。超过128m的视频必须采用分片上传。视频总大小4GB以内。单个分片建议20MB，最小5MB。
func (v *videoDouYin) PartComplete(openId, accessToken, uploadId string) (*VideoUpload, error) {
	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("upload_id", uploadId)       // 分片上传的标记。有限时间为2小时。

	urlStr := fmt.Sprintf("%s/video/part/complete/?%s", base_url_douyin, params.Encode())

	buf, err := goo.PostJson(urlStr, []byte{})
	if err != nil {
		goo.Log.Error(tag, "video-part-complete:http-err", err.Error())
		return nil, err
	}

	vu := &VideoUpload{}
	if err := json.Unmarshal(buf, vu); err != nil {
		goo.Log.Error(tag, "video-part-complete:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "video-part-complete:result", vu)

	if vu.Data.ErrorCode != 0 {
		return nil, errors.New(vu.Data.Description + ":" + vu.Extra.SubDescription)
	}

	return vu, nil
}

// 创建视频
// 该接口用于创建抖音视频（支持话题, 小程序等功能）。该接口适用于抖音。
func (v *videoDouYin) Create(openId, accessToken, videoId, text string, atUsers []string) (*VideoCreate, error) {
	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权

	urlStr := fmt.Sprintf("%s/video/create/?%s", base_url_douyin, params.Encode())

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
		goo.Log.Error(tag, "video-create:http-err", err.Error())
		return nil, err
	}

	vc := &VideoCreate{}
	if err := json.Unmarshal(buf, vc); err != nil {
		goo.Log.Error(tag, "video-create:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "video-create:result", vc)

	if vc.Data.ErrorCode != 0 {
		return nil, errors.New(vc.Data.Description + ":" + vc.Extra.SubDescription)
	}

	return vc, nil
}

// 查询授权账号视频数据
// 需要用户授权，需要申请权限，该接口用于分页获取用户所有视频的数据，返回的数据是实时的。该接口适用于抖音。
func (v *videoDouYin) List(openId, accessToken string, cursor, count int64) (*VideoList, error) {
	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("cursor", fmt.Sprintf("%d", cursor))
	params.Add("count", fmt.Sprintf("%d", count))

	urlStr := fmt.Sprintf("%s/video/list/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "video-list:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "video-list:http-err", err.Error())
		return nil, err
	}

	vl := &VideoList{}
	if err := json.Unmarshal(buf, vl); err != nil {
		goo.Log.Error(tag, "video-list:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "video-list:result", vl)

	if vl.Data.ErrorCode != 0 {
		return nil, errors.New(vl.Data.Description + ":" + vl.Extra.SubDescription)
	}

	return vl, nil
}

// 查询指定视频数据
// 需要用户授权，需要申请权限，该接口 用于查询用户特定视频的数据, 如点赞数, 播放数等，返回的数据是实时的。该接口适用于抖音。
func (v *videoDouYin) Data(openId, accessToken string, itemIds []string) (*VideoData, error) {
	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权

	urlStr := fmt.Sprintf("%s/video/data/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "video-data:url", urlStr)

	bd, _ := json.Marshal(map[string]interface{}{"item_ids": itemIds})
	buf, err := goo.PostJson(urlStr, bd)
	if err != nil {
		goo.Log.Error(tag, "video-data:http-err", err.Error())
		return nil, err
	}

	vd := &VideoData{}
	if err := json.Unmarshal(buf, vd); err != nil {
		goo.Log.Error(tag, "video-data:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "video-data:result", vd)

	if vd.Data.ErrorCode != 0 {
		return nil, errors.New(vd.Data.Description + ":" + vd.Extra.SubDescription)
	}

	return vd, nil
}

// 删除授权用户发布的视频
// 需要用户授权，需要申请权限，该接口用于删除授权用户该抖音账号下的视频。该接口适用于抖音。
func (v *videoDouYin) Delete(openId, accessToken string, itemId string) (*Result, error) {
	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权

	urlStr := fmt.Sprintf("%s/video/delete/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "video-delete:url", urlStr)

	bd, _ := json.Marshal(map[string]interface{}{"item_id": itemId})
	buf, err := goo.PostJson(urlStr, bd)
	if err != nil {
		goo.Log.Error(tag, "video-delete:http-err", err.Error())
		return nil, err
	}

	rst := &Result{}
	if err := json.Unmarshal(buf, rst); err != nil {
		goo.Log.Error(tag, "video-delete:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "video-delete:result", rst)

	if rst.Data.ErrorCode != 0 {
		return nil, errors.New(rst.Data.Description + ":" + rst.Extra.SubDescription)
	}

	return rst, nil
}

// 上传图片到文件服务器
// 需要申请权限，需要用户授权，该接口用于上传图片到文件服务器，得到图片的唯一标志image_id。该接口适用于抖音。
// 图片大小不超过100M。
func (v *videoDouYin) ImageUpload(openId, accessToken, filename string, f io.Reader) (*ImageUpload, error) {
	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权

	urlStr := fmt.Sprintf("%s/image/upload/?%s", base_url_douyin, params.Encode())

	buf, err := goo.Upload(urlStr, "image", filename, f, nil)
	if err != nil {
		goo.Log.Error(tag, "image-upload:http-err", err.Error())
		return nil, err
	}

	iu := &ImageUpload{}
	if err := json.Unmarshal(buf, iu); err != nil {
		goo.Log.Error(tag, "image-upload:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "image-upload:result", iu)

	if iu.Data.ErrorCode != 0 {
		return nil, errors.New(iu.Data.Description + ":" + iu.Extra.SubDescription)
	}

	return iu, nil
}

// 发布图片
// 该接口用于发布图片抖音（支持话题，小程序等功能）；该接口适用于抖音。
func (v *videoDouYin) ImageCreate(openId, accessToken, imageId, text string, atUsers []string) (*ImageCreate, error) {
	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权

	urlStr := fmt.Sprintf("%s/image/create/?%s", base_url_douyin, params.Encode())

	data := map[string]interface{}{
		"text":     text,
		"image_id": imageId,
	}
	if atUsers != nil {
		data["at_users"] = atUsers
	}
	body, _ := json.Marshal(data)

	buf, err := goo.PostJson(urlStr, body)
	if err != nil {
		goo.Log.Error(tag, "image-create:http-err", err.Error())
		return nil, err
	}

	ic := &ImageCreate{}
	if err := json.Unmarshal(buf, ic); err != nil {
		goo.Log.Error(tag, "image-create:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "image-create:result", ic)

	if ic.Data.ErrorCode != 0 {
		return nil, errors.New(ic.Data.Description + ":" + ic.Extra.SubDescription)
	}

	return ic, nil
}
