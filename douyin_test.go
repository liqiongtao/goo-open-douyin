package goo_open_douyin

var (
	dy = New(client_key, client_secret)

	client_key    = ""
	client_secret = ""

	oauth_connect_redirect_url   = ""
	oauth_authorize_redirect_url = ""

	code        = ""
	openId      = ""
	accessToken = ""
	itemId      = ""
	videoFile   = "1.mp4"
	imageFile   = "logo.png"
	keyword     = "中国"
	secItemId   = ""
	content     = "中国"
	commentId = ""
)

var (
	scopes = []string{
		// -------------------------------
		// 用户权限
		// -------------------------------

		"user_info", // 授权登录与用户基础信息
		// "following.list", // 关注列表
		// "fans.list",      // 粉丝列表
		"renew_refresh_token", // 授权动态续期，refresh_token支持动态续期，每次续期+30天，续期次数上限为5次
		// "login_id",            // 静默授权，直接获取该用户的open id
		// "mobile_alert",        // 获取用户手机号，用抖音帐号登录第三方平台，获得用户在抖音上的手机号码

		// -------------------------------
		// 视频权限
		// -------------------------------

		"video.create", // 视频发布及管理，服务器端直接发布图片或视频，同时支持获取内容数据以及管理内容状态
		"video.delete", // 同上
		"video.data",   // 同上
		"video.list",   // 同上
		// "toutiao.video.create", // 同上，应用于toutiao
		// "toutiao.video.data",   // 同上，应用于toutiao
		// "xigua.video.data",     // 同上，应用于xigua
		// "xigua.video.create",   // 同上，应用于xigua
		// "im.share",             // 分享给抖音好友/群，从第三方APP分享视频给抖音好友/群
		"video.search",         // 关键词视频管理，包含通过关键词获取抖音视频及该视频下评论，并进行回复的能力
		"video.search.comment", // 关键词视频管理，包含通过关键词获取抖音视频及该视频下评论，并进行回复的能力
		"aweme.share",          // 分享视频至抖音，外部内容分享到抖音，可携带指定话题，并获取分享内容消费数据（浏览量以及点赞量等数据）

		// -------------------------------
		// 互动权限
		// -------------------------------

		// "video.comment", // 互动管理（企业号），获取并管理评论和私信（需要授权用户是企业号）
		// "im",            // 互动管理（企业号），获取并管理评论和私信（需要授权用户是企业号）
		"item.comment", // 评论管理（普通用户），获取并管理评论和私信（需要授权用户是企业号）

		// -------------------------------
		// 数据权限
		// -------------------------------

		"data.external.user", // 用户数据，用户授权后，该接口可用于查询用户的获赞、评论、分享，主页访问等相关数据
		"data.external.item", // 视频管理，用户授权后，该接口可用于查询作品的获赞，评论，分享等相关数据
		"fans.data",          // 粉丝画像数据，用户授权后，该接口可用于获取用户粉丝画像数据
		// "star_top_score_display",    // 星图数据，包含星图达人与达人对应各指数评估分，以及星图6大热门维度下的达人榜单
		// "star_tops",                 // 星图数据，包含星图达人与达人对应各指数评估分，以及星图6大热门维度下的达人榜单
		// "star_author_score_display", // 星图数据，包含星图达人与达人对应各指数评估分，以及星图6大热门维度下的达人榜单
		// "discovery.ent",             // 抖音影视综榜单数据，获取抖音电影榜，电视剧榜以及综艺榜
		// "data.external.sdk_share",   // SDK分享视频数据，获取通过分享SDK分享视频数据
		"hotsearch", // 抖音热点（限时免费），获取抖音热门内容

		// -------------------------------
		// 特殊权限
		// -------------------------------

		// "share_with_source", // 分享携带来源标签，分享携带来源标签，用户可点击标签进入转化页(暂不对外开放)
		// "poi.search",        // 查询POI地点信息，用于查询POI信息
		"micapp.is_legal", // 查询小程序挂载权限，提供一个接口给开发者校验小程序appid是否可挂载到短视频

		// -------------------------------
		// JSBridge
		// -------------------------------

		"js.ticket",         // 签名ticket，获取签名ticket
		"jsb.open.auth",     // 跳转用户授权页面，跳转到原生授权页面
		"jsb.open.showAuth", // 唤起用户授权界面，唤起原生半屏授权界面
	}
)
