package goo_open_douyin

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/liqiongtao/goo"
)

// --------------------------------------------
// -- 抖音开放平台 - 数据开放服务
// --------------------------------------------

type dataExternal struct {
	douyin
}

// 用户数据 - 获取用户视频情况
// 该接口用于获取用户视频情况。
// 注：用户首次授权应用后，需要第二天才会产生全部的数据；
func (cm *dataExternal) UserItem(openId, accessToken string, dateType int64) (*DataExternalUserItem, error) {
	params := url.Values{}
	params.Add("open_id", openId)                        // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken)              // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("date_type", fmt.Sprintf("%d", dateType)) // 近7/15天；输入7代表7天、15代表15天、30代表30天

	urlStr := fmt.Sprintf("%s/data/external/user/item/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "data-external-user-item:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "data-external-user-item:http-err", err.Error())
		return nil, err
	}

	ui := &DataExternalUserItem{}
	if err := json.Unmarshal(buf, ui); err != nil {
		goo.Log.Error(tag, "data-external-user-item:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "data-external-user-item:result", ui)

	if ui.Data.ErrorCode != 0 {
		return nil, errors.New(ui.Data.Description)
	}

	return ui, nil
}

// 用户数据 - 获取用户粉丝数
// 该接口用于获取用户粉丝数。
// 注：用户首次授权应用后，需要第二天才会产生全部的数据；
func (cm *dataExternal) UserFans(openId, accessToken string, dateType int64) (*DataExternalUserFans, error) {
	params := url.Values{}
	params.Add("open_id", openId)                        // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken)              // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("date_type", fmt.Sprintf("%d", dateType)) // 近7/15天；输入7代表7天、15代表15天、30代表30天

	urlStr := fmt.Sprintf("%s/data/external/user/fans/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "data-external-user-fans:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "data-external-user-fans:http-err", err.Error())
		return nil, err
	}

	uf := &DataExternalUserFans{}
	if err := json.Unmarshal(buf, uf); err != nil {
		goo.Log.Error(tag, "data-external-user-fans:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "data-external-user-fans:result", uf)

	if uf.Data.ErrorCode != 0 {
		return nil, errors.New(uf.Data.Description)
	}

	return uf, nil
}

// 用户数据 - 获取用户点赞数
// 该接口用于获取用户点赞数。
// 注：用户首次授权应用后，需要第二天才会产生全部的数据；
func (cm *dataExternal) UserLike(openId, accessToken string, dateType int64) (*DataExternalUserLike, error) {
	params := url.Values{}
	params.Add("open_id", openId)                        // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken)              // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("date_type", fmt.Sprintf("%d", dateType)) // 近7/15天；输入7代表7天、15代表15天、30代表30天

	urlStr := fmt.Sprintf("%s/data/external/user/like/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "data-external-user-like:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "data-external-user-like:http-err", err.Error())
		return nil, err
	}

	ul := &DataExternalUserLike{}
	if err := json.Unmarshal(buf, ul); err != nil {
		goo.Log.Error(tag, "data-external-user-like:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "data-external-user-like:result", ul)

	if ul.Data.ErrorCode != 0 {
		return nil, errors.New(ul.Data.Description)
	}

	return ul, nil
}

// 用户数据 - 获取用户评论数
// 该接口用于获取用户评论数。
// 注：用户首次授权应用后，需要第二天才会产生全部的数据；
func (cm *dataExternal) UserComment(openId, accessToken string, dateType int64) (*DataExternalUserComment, error) {
	params := url.Values{}
	params.Add("open_id", openId)                        // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken)              // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("date_type", fmt.Sprintf("%d", dateType)) // 近7/15天；输入7代表7天、15代表15天、30代表30天

	urlStr := fmt.Sprintf("%s/data/external/user/comment/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "data-external-user-comment:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "data-external-user-comment:http-err", err.Error())
		return nil, err
	}

	ul := &DataExternalUserComment{}
	if err := json.Unmarshal(buf, ul); err != nil {
		goo.Log.Error(tag, "data-external-user-comment:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "data-external-user-comment:result", ul)

	if ul.Data.ErrorCode != 0 {
		return nil, errors.New(ul.Data.Description)
	}

	return ul, nil
}

// 用户数据 - 获取用户分享数
// 该接口用于获取用户分享数。
// 注：用户首次授权应用后，需要第二天才会产生全部的数据；
func (cm *dataExternal) UserShare(openId, accessToken string, dateType int64) (*DataExternalUserShare, error) {
	params := url.Values{}
	params.Add("open_id", openId)                        // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken)              // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("date_type", fmt.Sprintf("%d", dateType)) // 近7/15天；输入7代表7天、15代表15天、30代表30天

	urlStr := fmt.Sprintf("%s/data/external/user/share/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "data-external-user-share:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "data-external-user-share:http-err", err.Error())
		return nil, err
	}

	ul := &DataExternalUserShare{}
	if err := json.Unmarshal(buf, ul); err != nil {
		goo.Log.Error(tag, "data-external-user-share:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "data-external-user-share:result", ul)

	if ul.Data.ErrorCode != 0 {
		return nil, errors.New(ul.Data.Description)
	}

	return ul, nil
}

// 用户数据 - 获取用户主页访问数
// 该接口用于获取用户主页访问数。
// 注：用户首次授权应用后，需要第二天才会产生全部的数据；
func (cm *dataExternal) UserProfile(openId, accessToken string, dateType int64) (*DataExternalUserProfile, error) {
	params := url.Values{}
	params.Add("open_id", openId)                        // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken)              // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("date_type", fmt.Sprintf("%d", dateType)) // 近7/15天；输入7代表7天、15代表15天、30代表30天

	urlStr := fmt.Sprintf("%s/data/external/user/profile/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "data-external-user-profile:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "data-external-user-profile:http-err", err.Error())
		return nil, err
	}

	ul := &DataExternalUserProfile{}
	if err := json.Unmarshal(buf, ul); err != nil {
		goo.Log.Error(tag, "data-external-user-profile:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "data-external-user-profile:result", ul)

	if ul.Data.ErrorCode != 0 {
		return nil, errors.New(ul.Data.Description)
	}

	return ul, nil
}

// 视频数据 - 获取视频基础数据
// 注：用户首次授权应用后，需要第二天才会产生全部的数据；
func (cm *dataExternal) ItemBase(openId, accessToken, itemId string) (*DataExternalItemBase, error) {
	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("item_id", itemId)           // item_id，仅能查询access_token对应用户上传的视频

	urlStr := fmt.Sprintf("%s/data/external/item/base/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "data-external-item-base:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "data-external-item-base:http-err", err.Error())
		return nil, err
	}

	ul := &DataExternalItemBase{}
	if err := json.Unmarshal(buf, ul); err != nil {
		goo.Log.Error(tag, "data-external-item-base:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "data-external-item-base:result", ul)

	if ul.Data.ErrorCode != 0 {
		return nil, errors.New(ul.Data.Description)
	}

	return ul, nil
}

// 视频数据 - 获取视频点赞数据
// 注：用户首次授权应用后，需要第二天才会产生全部的数据；
func (cm *dataExternal) ItemLike(openId, accessToken, itemId string, dateType int64) (*DataExternalItemLike, error) {
	params := url.Values{}
	params.Add("open_id", openId)                        // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken)              // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("item_id", itemId)                        // item_id，仅能查询access_token对应用户上传的视频
	params.Add("date_type", fmt.Sprintf("%d", dateType)) // 近7/15天；输入7代表7天、15代表15天、30代表30天

	urlStr := fmt.Sprintf("%s/data/external/item/like/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "data-external-item-like:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "data-external-item-like:http-err", err.Error())
		return nil, err
	}

	ul := &DataExternalItemLike{}
	if err := json.Unmarshal(buf, ul); err != nil {
		goo.Log.Error(tag, "data-external-item-like:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "data-external-item-like:result", ul)

	if ul.Data.ErrorCode != 0 {
		return nil, errors.New(ul.Data.Description)
	}

	return ul, nil
}

// 视频数据 - 获取视频评论数据
// 注：用户首次授权应用后，需要第二天才会产生全部的数据；
func (cm *dataExternal) ItemComment(openId, accessToken, itemId string, dateType int64) (*DataExternalItemComment, error) {
	params := url.Values{}
	params.Add("open_id", openId)                        // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken)              // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("item_id", itemId)                        // item_id，仅能查询access_token对应用户上传的视频
	params.Add("date_type", fmt.Sprintf("%d", dateType)) // 近7/15天；输入7代表7天、15代表15天、30代表30天

	urlStr := fmt.Sprintf("%s/data/external/item/comment/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "data-external-item-comment:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "data-external-item-comment:http-err", err.Error())
		return nil, err
	}

	ul := &DataExternalItemComment{}
	if err := json.Unmarshal(buf, ul); err != nil {
		goo.Log.Error(tag, "data-external-item-comment:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "data-external-item-comment:result", ul)

	if ul.Data.ErrorCode != 0 {
		return nil, errors.New(ul.Data.Description)
	}

	return ul, nil
}

// 视频数据 - 获取视频播放数据
// 注：用户首次授权应用后，需要第二天才会产生全部的数据；
func (cm *dataExternal) ItemPlay(openId, accessToken, itemId string, dateType int64) (*DataExternalItemPlay, error) {
	params := url.Values{}
	params.Add("open_id", openId)                        // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken)              // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("item_id", itemId)                        // item_id，仅能查询access_token对应用户上传的视频
	params.Add("date_type", fmt.Sprintf("%d", dateType)) // 近7/15天；输入7代表7天、15代表15天、30代表30天

	urlStr := fmt.Sprintf("%s/data/external/item/play/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "data-external-item-play:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "data-external-item-play:http-err", err.Error())
		return nil, err
	}

	ul := &DataExternalItemPlay{}
	if err := json.Unmarshal(buf, ul); err != nil {
		goo.Log.Error(tag, "data-external-item-play:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "data-external-item-play:result", ul)

	if ul.Data.ErrorCode != 0 {
		return nil, errors.New(ul.Data.Description)
	}

	return ul, nil
}

// 视频数据 - 获取视频分享数据
// 注：用户首次授权应用后，需要第二天才会产生全部的数据；
func (cm *dataExternal) ItemShare(openId, accessToken, itemId string, dateType int64) (*DataExternalItemShare, error) {
	params := url.Values{}
	params.Add("open_id", openId)                        // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken)              // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("item_id", itemId)                        // item_id，仅能查询access_token对应用户上传的视频
	params.Add("date_type", fmt.Sprintf("%d", dateType)) // 近7/15天；输入7代表7天、15代表15天、30代表30天

	urlStr := fmt.Sprintf("%s/data/external/item/share/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "data-external-item-share:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "data-external-item-share:http-err", err.Error())
		return nil, err
	}

	ul := &DataExternalItemShare{}
	if err := json.Unmarshal(buf, ul); err != nil {
		goo.Log.Error(tag, "data-external-item-share:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "data-external-item-share:result", ul)

	if ul.Data.ErrorCode != 0 {
		return nil, errors.New(ul.Data.Description)
	}

	return ul, nil
}

// 粉丝画像数据 - 获取用户粉丝数据
// 该接口用于查询用户的粉丝数据，如性别分布，年龄分布，地域分布等。
// 注：用户首次授权应用后，需要间隔2天才会产生全部的数据；并只提供粉丝大于100的用户数据。
func (cm *dataExternal) FansData(openId, accessToken string) (*FansData, error) {
	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权

	urlStr := fmt.Sprintf("%s/fans/data/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "fans-data:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "fans-data:http-err", err.Error())
		return nil, err
	}

	fd := &FansData{}
	if err := json.Unmarshal(buf, fd); err != nil {
		goo.Log.Error(tag, "fans-data:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "fans-data:result", fd)

	if fd.Data.ErrorCode != 0 {
		return nil, errors.New(fd.Data.Description)
	}

	return fd, nil
}

// 热点视频数据 - 获取实时热点词
// 注意： 热点榜约每两个小时刷新一次。
func (cm *dataExternal) HotSearchSentences(accessToken string) (*HotSearchSentences, error) {
	params := url.Values{}
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权

	urlStr := fmt.Sprintf("%s/hotsearch/sentences/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "hotsearch-sentences:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "hotsearch-sentences:http-err", err.Error())
		return nil, err
	}

	fd := &HotSearchSentences{}
	if err := json.Unmarshal(buf, fd); err != nil {
		goo.Log.Error(tag, "hotsearch-sentences:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "hotsearch-sentences:result", fd)

	if fd.Data.ErrorCode != 0 {
		return nil, errors.New(fd.Data.Description)
	}

	return fd, nil
}

// 热点视频数据 - 获取上升词
// 注意： 热点榜约每两个小时刷新一次。
func (cm *dataExternal) HotSearchTrendingSentences(accessToken string, cursor, count int64) (*HotSearchTrendingSentences, error) {
	params := url.Values{}
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("cursor", fmt.Sprintf("%d", cursor))
	params.Add("count", fmt.Sprintf("%d", count))

	urlStr := fmt.Sprintf("%s/hotsearch/trending/sentences/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "hotsearch-trending-sentences:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "hotsearch-trending-sentences:http-err", err.Error())
		return nil, err
	}

	fd := &HotSearchTrendingSentences{}
	if err := json.Unmarshal(buf, fd); err != nil {
		goo.Log.Error(tag, "hotsearch-trending-sentences:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "hotsearch-trending-sentences:result", fd)

	if fd.Data.ErrorCode != 0 {
		return nil, errors.New(fd.Data.Description)
	}

	return fd, nil
}

// 热点视频数据 - 获取热点词聚合的视频
func (cm *dataExternal) HotSearchVideos(accessToken, hotSentence string) (*HotSearchVideos, error) {
	params := url.Values{}
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("hot_sentence", hotSentence) // 热点词

	urlStr := fmt.Sprintf("%s/hotsearch/videos/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "hotsearch-videos:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "hotsearch-videos:http-err", err.Error())
		return nil, err
	}

	fd := &HotSearchVideos{}
	if err := json.Unmarshal(buf, fd); err != nil {
		goo.Log.Error(tag, "hotsearch-videos:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "hotsearch-videos:result", fd)

	if fd.Data.ErrorCode != 0 {
		return nil, errors.New(fd.Data.Description)
	}

	return fd, nil
}

// 星图数据 - 星图榜单数据 - 获取抖音星图达人热榜
// 该接口用于查询抖音星图榜单相关数据。【调用接口给用户展示时必须带有【星图指数】或【星图达人榜】字样】。
func (cm *dataExternal) StarHotList(accessToken string, hotListType int64) (*StarHotList, error) {
	params := url.Values{}
	params.Add("access_token", accessToken)                     // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("hot_list_type", fmt.Sprintf("%d", hotListType)) // 达人热榜类型 * `1` - 星图指数榜 * `2` - 涨粉指数榜 * `3` - 性价比指数榜 * `4` - 种草指数榜 * `5` - 精选指数榜 * `6` - 传播指数榜

	urlStr := fmt.Sprintf("%s/star/hot_list/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "star-hot-list:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "star-hot-list:http-err", err.Error())
		return nil, err
	}

	shl := &StarHotList{}
	if err := json.Unmarshal(buf, shl); err != nil {
		goo.Log.Error(tag, "star-hot-list:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "star-hot-list:result", shl)

	if shl.Data.ErrorCode != 0 {
		return nil, errors.New(shl.Data.Description)
	}

	return shl, nil
}

// 星图数据 - 星图达人指数数据 - 获取抖音星图达人指数
// 该接口用于查询抖音星图达人相关数据。【调用接口给用户展示时必须带有【星图指数】或【星图达人榜】字样】。
func (cm *dataExternal) StarAuthorScore(openId, accessToken string) (*StarAuthorScore, error) {
	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权

	urlStr := fmt.Sprintf("%s/star/author_score/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "star-author-score:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "star-author-score:http-err", err.Error())
		return nil, err
	}

	sas := &StarAuthorScore{}
	if err := json.Unmarshal(buf, sas); err != nil {
		goo.Log.Error(tag, "star-author-score:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "star-author-score:result", sas)

	if sas.Data.ErrorCode != 0 {
		return nil, errors.New(sas.Data.Description)
	}

	return sas, nil
}

// 星图数据 - 星图达人指数数据 - 获取抖音星图达人指数数据V2
// 该接口用于查询抖音星图达人相关数据。【调用接口给用户展示时必须带有【星图指数】或【星图达人榜】字样】。
func (cm *dataExternal) StarAuthorScoreV2(accessToken, uniqueId string) (*StarAuthorScore, error) {
	params := url.Values{}
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("unique_id", uniqueId)       // 达人抖音号

	urlStr := fmt.Sprintf("%s/star/author_score_v2/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "star-author-score-v2:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "star-author-score-v2:http-err", err.Error())
		return nil, err
	}

	fd := &StarAuthorScore{}
	if err := json.Unmarshal(buf, fd); err != nil {
		goo.Log.Error(tag, "star-author-score-v2:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "star-author-score-v2:result", fd)

	if fd.Data.ErrorCode != 0 {
		return nil, errors.New(fd.Data.Description)
	}

	return fd, nil
}

// 抖音影视综艺榜单数据 - 获取抖音电影榜、抖音电视剧榜、抖音综艺榜
// 该接口用于查询抖音影视综榜单数据，可以根据版本返回往期榜单。返回的内容按热度降序排列，电影榜最多返回top30、电视剧榜和综艺榜最多返回top10。
//
// 上榜规则：
// 电影上映时间在当前时间的前60后45天。
// 电视剧上映时间在当前时间的前60天后30天。
// 综艺上映时间在当前时间的前90天后30天。（除部分全年持续播出综艺）
//
// 榜单规则：
// 电影榜/电视剧榜/综艺榜 分为“本周榜单”和“往期榜单”两部分。
// 本周榜单：展示本周上榜影片的动态变化，热度值在每天的12:00~24:00间，每日更新一次 。
// 往期榜单：每周日最后一张榜单（次日周一12:00发布），即当周最终上榜的影片排名及数据详情，进入往期榜单。
// 榜单统计周期：周一00:00-周日24:00。
func (cm *dataExternal) DiscoveryEntRankItem(accessToken string, ttype, version int) (*DiscoveryEntRankItem, error) {
	params := url.Values{}
	params.Add("access_token", accessToken)           // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("type", fmt.Sprintf("%d", ttype))      // 榜单类型： * 1 - 电影 * 2 - 电视剧 * 3 - 综艺
	params.Add("version", fmt.Sprintf("%d", version)) // 榜单版本：空值默认为本周榜单

	urlStr := fmt.Sprintf("%s/discovery/ent/rank/item/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "discovery-ent-rank-item:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "discovery-ent-rank-item:http-err", err.Error())
		return nil, err
	}

	fd := &DiscoveryEntRankItem{}
	if err := json.Unmarshal(buf, fd); err != nil {
		goo.Log.Error(tag, "discovery-ent-rank-item:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "discovery-ent-rank-item:result", fd)

	if fd.Data.ErrorCode != 0 {
		return nil, errors.New(fd.Data.Description)
	}

	return fd, nil
}

// 抖音影视综艺榜单数据 - 获取抖音影视综榜单版本
// 该接口用于查询抖音影视综榜单数据，可以根据版本返回往期榜单。返回的内容按热度降序排列，电影榜最多返回top30、电视剧榜和综艺榜最多返回top10。
//
// 上榜规则
// 电影上映时间在当前时间的前60后45天。
// 电视剧上映时间在当前时间的前60天后30天。
// 综艺上映时间在当前时间的前90天后30天。（除部分全年持续播出综艺）
// 榜单规则
// 电影榜/电视剧榜/综艺榜 分为“本周榜单”和“往期榜单”两部分。
//
// 本周榜单：展示本周上榜影片的动态变化，热度值在每天的12:00~24:00间，每日更新一次 。
// 往期榜单：每周日最后一张榜单（次日周一12:00发布），即当周最终上榜的影片排名及数据详情，进入往期榜单。
// 榜单统计周期：周一00:00-周日24:00。
func (cm *dataExternal) DiscoveryEntRankVersion(accessToken string, cursor, count int64, typ int) (*DiscoveryEntRankVersion, error) {
	params := url.Values{}
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("cursor", fmt.Sprintf("%d", cursor))
	params.Add("count", fmt.Sprintf("%d", count))
	params.Add("type", fmt.Sprintf("%d", typ)) // 榜单类型： * 1 - 电影 * 2 - 电视剧 * 3 - 综艺

	urlStr := fmt.Sprintf("%s/discovery/ent/rank/version/?%s", base_url_douyin, params.Encode())
	goo.Log.Debug(tag, "discovery-ent-rank-version:url", urlStr)

	buf, err := goo.Get(urlStr)
	if err != nil {
		goo.Log.Error(tag, "discovery-ent-rank-version:http-err", err.Error())
		return nil, err
	}

	fd := &DiscoveryEntRankVersion{}
	if err := json.Unmarshal(buf, fd); err != nil {
		goo.Log.Error(tag, "discovery-ent-rank-version:err", err.Error())
		return nil, err
	}

	goo.Log.Debug(tag, "discovery-ent-rank-version:result", fd)

	if fd.Data.ErrorCode != 0 {
		return nil, errors.New(fd.Data.Description)
	}

	return fd, nil
}
