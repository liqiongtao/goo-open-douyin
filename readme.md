# 授权过程

1. 通过 “抖音获取授权码” 或 “头条获取授权码”，展示授权页面

```
https://open.douyin.com/platform/oauth/connect/?client_key=&optionalScope=&redirect_uri=&response_type=code&scope=user_info%2Crenew_refresh_token%2Cvideo.create%2Cvideo.delete%2Cvideo.data%2Cvideo.list%2Cvideo.search%2Cvideo.search.comment%2Caweme.share%2Citem.comment%2Cdata.external.user%2Cdata.external.item%2Cfans.data%2Chotsearch%2Cmicapp.is_legal%2Cjs.ticket%2Cjsb.open.auth%2Cjsb.open.showAuth&state=
```

注意：“视频发布及管理”的 `scope` 要区分 `toutiao`、`xigua`

2. 用户使用抖音 app 扫描授权后，跳转链接，获取`code`

```
http://douyin.yz210.com/?code=&state=#/login
```

3. 通过 ”获取 access_token“，获取到用户相关信息

-   请求

```
https://open.douyin.com/oauth/access_token/?client_key=&client_secret=&code=&grant_type=authorization_code
```

-   返回

```
{
    "data": {
        "access_token": "",
        "captcha": "",
        "desc_url": "",
        "description": "",
        "error_code": 0,
        "expires_in": 1296000,
        "open_id": "",
        "refresh_expires_in": 2592000,
        "refresh_token": "",
        "scope": "data.external.item,data.external.user,fans.data,item.comment,user_info,video.create,video.data,video.delete,video.list,video.search"
    },
    "message": "success"
}
```

# 获取用户信息

-   请求

```
https://open.douyin.com/oauth/userinfo/?access_token=&open_id=
```

-   返回

```
{
    "data": {
        "avatar": "https://p6-dy-ipv6.byteimg.com/aweme/100x100/26ee40000df8420eb3799.jpeg?from=4010531038",
        "avatar_larger": "https://p3-dy-ipv6.byteimg.com/aweme/1080x1080/26ee40000df8420eb3799.jpeg?from=4010531038",
        "captcha": "",
        "city": "北京",
        "client_key": "",
        "country": "中国",
        "desc_url": "",
        "description": "",
        "district": "",
        "e_account_role": "",
        "error_code": 0,
        "gender": 1,
        "nickname": "",
        "open_id": "",
        "province": "北京",
        "union_id": ""
    },
    "message": "success"
}
```
