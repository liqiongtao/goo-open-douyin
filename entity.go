package goo_open_douyin

import "encoding/json"

type Extra struct {
	SubDescription string `json:"sub_description"`
	SubErrorCode   int    `json:"sub_error_code"`
	Description    string `json:"description"`
	ErrorCode      int    `json:"error_code"`
	Now            int64  `json:"now"`
	Logid          string `json:"logid"`
}

type Result struct {
	Data struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
	} `json:"data"`
	Extra `json:"extra"`
}

type Qrcode struct {
	Data struct {
		Qrcode      string `json:"qrcode"`
		Token       string `json:"token"`
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
	} `json:"data"`
	Message string `json:"message"`
}

type CheckQrcode struct {
	Data struct {
		RedirectUrl string `json:"redirect_url"`
		Status      string `json:"status"`
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
	} `json:"data"`
	Message string `json:"message"`
}

type AccessToken struct {
	Data struct {
		AccessToken      string `json:"access_token"`
		Captcha          string `json:"captcha"`
		DescUrl          string `json:"desc_url"`
		Description      string `json:"description"`
		ErrorCode        int    `json:"error_code"`
		ExpiresIn        int64  `json:"expires_in"`
		OpenId           string `json:"open_id"`
		RefreshExpiresIn int64  `json:"refresh_expires_in"`
		RefreshToken     string `json:"refresh_token"`
		Scope            string `json:"scope"`
	} `json:"data"`
	Message string `json:"message"`
}

type RenewRefreshToken struct {
	Data struct {
		Description  string `json:"description"`
		ErrorCode    int    `json:"error_code"`
		ExpiresIn    int64  `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
	} `json:"data"`
	Message string `json:"message"`
}

type ClientToken struct {
	Data struct {
		ExpiresIn   int64  `json:"expires_in"`
		AccessToken string `json:"access_token"`
		Description string `json:"description"`
		ErrorCode   int    `json:"error_code"`
	} `json:"data"`
	Message string `json:"message"`
}

type RefreshToken struct {
	Data struct {
		RefreshToken string `json:"refresh_token"`
		Scope        string `json:"scope"`
		AccessToken  string `json:"access_token"`
		Description  string `json:"description"`
		ErrorCode    int    `json:"error_code"`
		ExpiresIn    int64  `json:"expires_in"`
		OpenId       string `json:"open_id"`
	} `json:"data"`
	Message string `json:"message"`
}

type UserInfo struct {
	Data struct {
		Avatar       string `json:"avatar"`
		AvatarLarger string `json:"avatar_larger"`
		Captcha      string `json:"captcha"`
		City         string `json:"city"`
		ClientKey    string `json:"client_key"`
		Country      string `json:"country"`
		DescUrl      string `json:"desc_url"`
		Description  string `json:"description"`
		District     string `json:"district"`
		EAccountRole string `json:"e_account_role"`
		ErrorCode    int    `json:"error_code"`
		Gender       int    `json:"gender"`
		Nickname     string `json:"nickname"`
		OpenId       string `json:"open_id"`
		Province     string `json:"province"`
		UnionId      string `json:"union_id"`
	} `json:"data"`
	Message string `json:"message"`
}

type FansList struct {
	Data struct {
		ErrorCode   int        `json:"error_code"`
		HasMore     bool       `json:"has_more"`
		List        []FansUser `json:"list"`
		Total       int64      `json:"total"`
		Cursor      int64      `json:"cursor"`
		Description string     `json:"description"`
	} `json:"data"`
	Extra struct {
		Now   int64  `json:"now"`
		Logid string `json:"logid"`
	} `json:"extra"`
}

type FansUser struct {
	OpenId   string `json:"open_id"`
	Province string `json:"province"`
	UnionId  string `json:"union_id"`
	Avatar   string `json:"avatar"`
	City     string `json:"city"`
	Country  string `json:"country"`
	Gender   int    `json:"gender"`
	Nickname string `json:"nickname"`
}

type FollowingList struct {
	Data struct {
		ErrorCode   int             `json:"error_code"`
		HasMore     bool            `json:"has_more"`
		List        []FollowingUser `json:"list"`
		Cursor      int64           `json:"cursor"`
		Description string          `json:"description"`
	} `json:"data"`
	Extra `json:"extra"`
}

type FollowingUser struct {
	Gender   int    `json:"gender"`
	Nickname string `json:"nickname"`
	OpenId   string `json:"open_id"`
	Province string `json:"province"`
	UnionId  string `json:"union_id"`
	Avatar   string `json:"avatar"`
	City     string `json:"city"`
	Country  string `json:"country"`
}

type VideoUpload struct {
	Data struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
		Video       struct {
			VideoId string `json:"video_id"`
			Width   int    `json:"width"`
			Height  int    `json:"height"`
		} `json:"video"`
	} `json:"data"`
	Extra `json:"extra"`
}

type VideoPartInit struct {
	Data struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
		UploadId    string `json:"upload_id"`
	} `json:"data"`
	Extra `json:"extra"`
}

type VideoCreate struct {
	Data struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
		ItemId      string `json:"item_id"`
	} `json:"data"`
	Extra `json:"extra"`
}

type VideoList struct {
	Data struct {
		Cursor      int64           `json:"cursor"`
		ErrorCode   int             `json:"error_code"`
		Description string          `json:"description"`
		HasMore     bool            `json:"has_more"`
		List        []VideoListItem `json:"list"`
	} `json:"data"`
	Extra `json:"extra"`
}

func (vl *VideoList) String() string {
	bf, _ := json.Marshal(vl)
	return string(bf)
}

type VideoListItem struct {
	VideoStatus int    `json:"video_status"` // 表示视频状态。1:已发布;2:不适宜公开;4:审核中
	Cover       string `json:"cover"`
	CreateTime  int64  `json:"create_time"`
	IsTop       bool   `json:"is_top"`
	ItemId      string `json:"item_id"`
	IsReviewed  bool   `json:"is_reviewed"` // 表示是否审核结束。审核通过或者失败都会返回true，审核中返回false。
	ShareUrl    string `json:"share_url"`
	Statistics  struct {
		DiggCount     int64 `json:"digg_count"`     // 点赞数
		DownloadCount int64 `json:"download_count"` // 下载数
		ForwardCount  int64 `json:"forward_count"`  // 转发数
		PlayCount     int64 `json:"play_count"`     // 播放数，只有作者本人可见。公开视频设为私密后，播放数也会返回0。
		ShareCount    int64 `json:"share_count"`    // 分享数
		CommentCount  int64 `json:"comment_count"`  // 评论数
	} `json:"statistics"`
	Title string `json:"title"`
}

type VideoData struct {
	Data struct {
		ErrorCode   int             `json:"error_code"`
		Description string          `json:"description"`
		List        []VideoDataItem `json:"list"`
	} `json:"data"`
	Extra `json:"extra"`
}

func (vd *VideoData) String() string {
	bf, _ := json.Marshal(vd)
	return string(bf)
}

type VideoDataItem struct {
	ShareUrl    string `json:"share_url"`
	VideoStatus int    `json:"video_status"` // 表示视频状态。1:已发布;2:不适宜公开;4:审核中
	Cover       string `json:"cover"`
	CreateTime  int64  `json:"create_time"`
	ItemId      string `json:"item_id"`
	Statistics  struct {
		DiggCount     int64 `json:"digg_count"`     // 点赞数
		DownloadCount int64 `json:"download_count"` // 下载数
		ForwardCount  int64 `json:"forward_count"`  // 转发数
		PlayCount     int64 `json:"play_count"`     // 播放数，只有作者本人可见。公开视频设为私密后，播放数也会返回0。
		ShareCount    int64 `json:"share_count"`    // 分享数
		CommentCount  int64 `json:"comment_count"`  // 评论数
	} `json:"statistics"`
	Title      string `json:"title"`
	IsReviewed bool   `json:"is_reviewed"` // 表示是否审核结束。审核通过或者失败都会返回true，审核中返回false。
	IsTop      bool   `json:"is_top"`
}

type ImageUpload struct {
	Data struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
		Image       struct {
			ImageId string `json:"image_id"`
			Width   int    `json:"width"`
			Height  int    `json:"height"`
		} `json:"image"`
	} `json:"data"`
	Extra `json:"extra"`
}

type ImageCreate struct {
	Data struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
		ItemId      string `json:"item_id"`
	} `json:"data"`
	Extra `json:"extra"`
}

type CommentList struct {
	Data struct {
		Cursor      int64             `json:"cursor"`
		ErrorCode   int               `json:"error_code"`
		Description string            `json:"description"`
		HasMore     bool              `json:"has_more"`
		List        []CommentListItem `json:"list"`
	} `json:"data"`
	Extra `json:"extra"`
}

type CommentListItem struct {
	ReplyCommentTotal int64  `json:"reply_comment_total"`
	Top               bool   `json:"top"`
	CommentId         string `json:"comment_id"`
	CommentUserId     string `json:"comment_user_id"`
	Content           string `json:"content"`
	CreateTime        int64  `json:"create_time"`
	DiggCount         int64  `json:"digg_count"`
}

type CommentReplyList struct {
	Data struct {
		Cursor      int64                  `json:"cursor"`
		ErrorCode   int                    `json:"error_code"`
		Description string                 `json:"description"`
		HasMore     bool                   `json:"has_more"`
		List        []CommentReplyListItem `json:"list"`
	} `json:"data"`
	Extra `json:"extra"`
}

type CommentReplyListItem struct {
	ReplyCommentTotal int64  `json:"reply_comment_total"`
	Top               bool   `json:"top"`
	CommentId         string `json:"comment_id"`
	CommentUserId     string `json:"comment_user_id"`
	Content           string `json:"content"`
	CreateTime        int64  `json:"create_time"`
	DiggCount         int    `json:"digg_count"`
}

type CommentReply struct {
	Data struct {
		CommentId   string `json:"comment_id"`
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
	} `json:"data"`
	Extra `json:"extra"`
}

type DataExternalUserItem struct {
	Data struct {
		ResultList  []DataExternalUserItemResultItem `json:"result_list"`
		ErrorCode   int                              `json:"error_code"`
		Description string                           `json:"description"`
	} `json:"data"`
	Extra `json:"extra"`
}

type DataExternalUserItemResultItem struct {
	Date       string `json:"date"`        // 	日期	yyyy-MM-dd
	NewIssue   int64  `json:"new_issue"`   // 每日发布内容数	200
	NewPlay    int64  `json:"new_play"`    // 每天新增视频播放	200
	TotalIssue int64  `json:"total_issue"` // 	每日内容总数	200
}

type DataExternalUserFans struct {
	Data struct {
		ResultList  []DataExternalUserFansResultItem `json:"result_list"`
		ErrorCode   int                              `json:"error_code"`
		Description string                           `json:"description"`
	} `json:"data"`
	Extra `json:"extra"`
}

type DataExternalUserFansResultItem struct {
	TotalFans int64  `json:"total_fans"` // 	每日总粉丝数	200
	Date      string `json:"date"`       // 	日期	yyyy-MM-dd
	NewFans   int64  `json:"new_fans"`   // 	每天新粉丝数	200
}

type DataExternalUserLike struct {
	Data struct {
		ResultList  []DataExternalUserLikeResultItem `json:"result_list"`
		ErrorCode   int                              `json:"error_code"`
		Description string                           `json:"description"`
	} `json:"data"`
	Extra `json:"extra"`
}

type DataExternalUserLikeResultItem struct {
	Date    string `json:"date"`     // 	日期	yyyy-MM-dd
	NewLike int64  `json:"new_like"` // 	新增点赞	200
}

type DataExternalUserComment struct {
	Data struct {
		ResultList  []DataExternalUserCommentResultItem `json:"result_list"`
		ErrorCode   int                                 `json:"error_code"`
		Description string                              `json:"description"`
	} `json:"data"`
	Extra `json:"extra"`
}

type DataExternalUserCommentResultItem struct {
	Date       string `json:"date"`        // 	日期	yyyy-MM-dd
	NewComment int64  `json:"new_comment"` // 	新增评论	200
}

type DataExternalUserShare struct {
	Data struct {
		ResultList  []DataExternalUserShareResultItem `json:"result_list"`
		ErrorCode   int                               `json:"error_code"`
		Description string                            `json:"description"`
	} `json:"data"`
	Extra `json:"extra"`
}

type DataExternalUserShareResultItem struct {
	Date       string `json:"date"`      // 	日期	yyyy-MM-dd
	NewComment int64  `json:"new_share"` // 	新增分享	200
}

type DataExternalUserProfile struct {
	Data struct {
		ResultList  []DataExternalUserProfileResultItem `json:"result_list"`
		ErrorCode   int                                 `json:"error_code"`
		Description string                              `json:"description"`
	} `json:"data"`
	Extra `json:"extra"`
}

type DataExternalUserProfileResultItem struct {
	Date      string `json:"date"`       // 	日期	yyyy-MM-dd
	ProfileUv int64  `json:"profile_uv"` // 	当日个人主页访问人数	200
}

type DataExternalItemBase struct {
	Data struct {
		Result struct {
			AvgPlayDuration float64 `json:"avg_play_duration"` // 	最近30天平均播放时长	200
			TotalComment    int64   `json:"total_comment"`     // 	最近30天评论数	200
			TotalLike       int64   `json:"total_like"`        // 	最近30天点赞数	200
			TotalPlay       int64   `json:"total_play"`        // 	最近30天播放次数	200
			TotalShare      int64   `json:"total_share"`       // 	最近30天分享数	200
		} `json:"result"`
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
	} `json:"data"`
	Extra `json:"extra"`
}

type DataExternalItemLike struct {
	Data struct {
		ResultList  []DataExternalItemLikeResultItem `json:"result_list"`
		ErrorCode   int                              `json:"error_code"`
		Description string                           `json:"description"`
	} `json:"data"`
	Extra `json:"extra"`
}

type DataExternalItemLikeResultItem struct {
	Date string `json:"date"` // 	日期	yyyy-MM-dd
	Like int64  `json:"like"` // 	每日点赞数	200
}

type DataExternalItemComment struct {
	Data struct {
		ResultList  []DataExternalItemCommentResultItem `json:"result_list"`
		ErrorCode   int                                 `json:"error_code"`
		Description string                              `json:"description"`
	} `json:"data"`
	Extra `json:"extra"`
}

type DataExternalItemCommentResultItem struct {
	Date    string `json:"date"`    // 	日期	yyyy-MM-dd
	Comment int64  `json:"comment"` // 	每日评论数	200
}

type DataExternalItemPlay struct {
	Data struct {
		ResultList  []DataExternalItemPlayResultItem `json:"result_list"`
		ErrorCode   int                              `json:"error_code"`
		Description string                           `json:"description"`
	} `json:"data"`
	Extra `json:"extra"`
}

type DataExternalItemPlayResultItem struct {
	Date string `json:"date"` // 	日期	yyyy-MM-dd
	Play int64  `json:"play"` // 	每日播放数	200
}

type DataExternalItemShare struct {
	Data struct {
		ResultList  []DataExternalItemShareResultItem `json:"result_list"`
		ErrorCode   int                               `json:"error_code"`
		Description string                            `json:"description"`
	} `json:"data"`
	Extra `json:"extra"`
}

type DataExternalItemShareResultItem struct {
	Date  string `json:"date"`  // 	日期	yyyy-MM-dd
	Share int64  `json:"share"` // 	每日分享数	200
}

type FansData struct {
	Data struct {
		ActiveDaysDistributions   []FansProfileDistribution `json:"active_days_distributions"`
		AgeDistributions          []FansProfileDistribution `json:"age_distributions"`
		DeviceDistributions       []FansProfileDistribution `json:"device_distributions"`
		FlowDistributions         []FansProfileDistribution `json:"flow_contributions"`
		GenderDistributions       []FansProfileDistribution `json:"gender_distributions"`
		GeographicalDistributions []FansProfileDistribution `json:"geographical_distributions"`
		InterestDistributions     []FansProfileDistribution `json:"interest_distributions"`
		AllFansNum                int64                     `json:"all_fans_num"`
		ErrorCode                 int                       `json:"error_code"`
		Description               string                    `json:"description"`
	} `json:"data"`
	Extra `json:"extra"`
}

type FansProfileDistribution struct {
	Item  string `json:"item"`  // 	分布的种类
	Value int64  `json:"value"` // 	分布的数值
}

type HotSearchSentences struct {
	Data struct {
		ActiveTime  string                   `json:"active_time"` // 刷新时间
		List        []HotSearchSentencesItem `json:"List"`
		ErrorCode   int                      `json:"error_code"`
		Description string                   `json:"description"`
	} `json:"data"`
	Extra `json:"extra"`
}

type HotSearchSentencesItem struct {
	Item     string `json:"hot_level"` // 	热度 综合点赞、评论、转发等计算得出
	Sentence int64  `json:"sentence"`  // 	热点词
}

type HotSearchTrendingSentences struct {
	Data struct {
		Cursor      int64                            `json:"cursor"` // 	用于下一页请求的cursor
		HasMore     bool                             `json:"has_more"`
		List        []HotSearchTrendingSentencesItem `json:"List"`
		Total       int                              `json:"total"`
		ErrorCode   int                              `json:"error_code"`
		Description string                           `json:"description"`
	} `json:"data"`
	Extra `json:"extra"`
}

type HotSearchTrendingSentencesItem struct {
	HotLevel int64  `json:"hot_level"` // 	热度 综合点赞、评论、转发等计算得出	2.998e+06
	Label    int64  `json:"label"`     // 	标签: * `0` - 无 * `1` - 新 * `2` - 推荐 * `3` - 热 * `4` - 爆 * `5` - 首发
	Sentence string `json:"sentence"`  // 	热点词	苹果发布AirPods Pro
}

type HotSearchVideos struct {
	Data struct {
		List        []HotSearchVideosItem `json:"List"`
		ErrorCode   int                   `json:"error_code"`
		Description string                `json:"description"`
	} `json:"data"`
	Extra `json:"extra"`
}

type HotSearchVideosItem struct {
	IsTop      bool   `json:"is_top"`    // 是否置顶	false
	ShareUrl   string `json:"share_url"` // 视频播放页面。视频播放页可能会失效，请在观看视频前调用/video/data/获取最新的播放页。	https://www.iesdouyin.com/share/video/QDlWd0EzdWVMU2Q0aU5tKzVaOElvVU03ODBtRHFQUCtLUHBSMHFRT21MVkFYYi9UMDYwemRSbVlxaWczNTd6RUJRc3MrM2hvRGlqK2EwNnhBc1lGUkpRPT0=/?region=CN&mid=6753173704399670023&u_code=12h9je425&titleType=title
	Statistics struct {
		PlayCount     int32 `json:"play_count"`     // 播放数，只有作者本人可见。公开视频设为私密后，播放数也会返回0。	300
		ShareCount    int32 `json:"share_count"`    // 分享数	10
		CommentCount  int32 `json:"comment_count"`  // 评论数	100
		DiggCount     int32 `json:"digg_count"`     // 点赞数	200
		DownloadCount int32 `json:"download_count"` // 下载数	10
		ForwardCount  int32 `json:"forward_count"`  // 转发数	10
	} `json:"statistics"` // 统计数据
	CreateTime  int64  `json:"create_time"`  // 视频创建时间戳	1.571075129e+09
	IsReviewed  bool   `json:"is_reviewed"`  // 表示是否审核结束。审核通过或者失败都会返回true，审核中返回false。	true
	Title       string `json:"title"`        // 视频标题	测试视频 #测试话题 @抖音小助手
	VideoStatus int32  `json:"video_status"` // 表示视频状态。1:已发布;2:不适宜公开;4:审核中	1
	Cover       string `json:"cover"`        // 视频封面
	ItemId      string `json:"item_id"`      // 视频id
}

type StarHotList struct {
	Data struct {
		HotListType            int               `json:"hot_list_type"`
		HotListUpdateTimestamp int               `json:"hot_list_update_timestamp"`
		HotListDescription     string            `json:"hot_list_description"`
		List                   []StarHotListItem `json:"List"`
		ErrorCode              int               `json:"error_code"`
		Description            string            `json:"description"`
	} `json:"data"`
	Extra `json:"extra"`
}

type StarHotListItem struct {
	Follower int64    `json:"follower"`  // 粉丝数
	NickName string   `json:"nick_name"` // 达人昵称
	Rank     int64    `json:"rank"`      // 热榜排名
	Score    float64  `json:"score"`     // 热榜类型对应的热榜指数	80.6
	Tags     []string `json:"tags"`
	UniqueId string   `json:"unique_id"` // 抖音号
}

type StarAuthorScore struct {
	Data struct {
		CpScore          float64 `json:"cp_score"`          // 	性价比指数	80.6
		ShopScore        float64 `json:"shop_score"`        // 	种草指数	80.6
		SpreadScore      float64 `json:"spread_score"`      // 	传播指数	80.6
		StarScore        float64 `json:"star_score"`        // 	星图指数	80.6
		CooperationScore float64 `json:"cooperation_score"` // 	合作指数	80.6
		Description      string  `json:"description"`       // 	错误码描述
		ErrorCode        int64   `json:"error_code"`        // 	错误码	0
		Follower         int64   `json:"follower"`          // 	粉丝数
		GrowthScore      float64 `json:"growth_score"`      // 涨粉指数	80.6
		NickName         string  `json:"nick_name"`         // 	达人昵称
		UpdateTimestamp  int64   `json:"update_timestamp"`  // 	达人指数更新时间戳	1.584418477e+09
		UniqueId         string  `json:"unique_id"`         // 	达人抖音号
	} `json:"data"`
	Extra `json:"extra"`
}

type DiscoveryEntRankItem struct {
	Data struct {
		Description string                        `json:"description"` // 	错误码描述
		ErrorCode   int64                         `json:"error_code"`  // 	错误码	0
		ActiveTime  int                           `json:"active_time"`
		HasMore     bool                          `json:"has_more"`
		List        []DiscoveryEntRankVersionItem `json:"list"`
	} `json:"data"`
	Extra `json:"extra"`
}

type DiscoveryEntRankVersion struct {
	Data struct {
		Description string                        `json:"description"` // 	错误码描述
		ErrorCode   int64                         `json:"error_code"`  // 	错误码	0
		Cursor      int64                         `json:"cursor"`
		HasMore     bool                          `json:"has_more"`
		List        []DiscoveryEntRankVersionItem `json:"list"`
	} `json:"data"`
	Extra `json:"extra"`
}

type DiscoveryEntRankVersionItem struct {
	ActiveTime string `json:"active_time"` // 	榜单生成时间	2020-03-30 12:00:00
	EndTime    string `json:"end_time"`    // 	榜单结束时间	2020-03-30 00:00:00
	StartTime  string `json:"start_time"`  // 	榜单起始时间	2020-03-23 00:00:00
	Type       int32  `json:"type"`        // 	类型：1=电影 2=电视剧 3=综艺	1
	Version    int32  `json:"version"`     // 	榜单版本	18
}
