package goo_open_douyin

type douyin struct {
	clientKey    string
	clientSecret string
}

func New(clientKey, clientSecret string) *douyin {
	return &douyin{
		clientKey:    clientKey,
		clientSecret: clientSecret,
	}
}

func (dy *douyin) OAuth() *oauth {
	return &oauth{douyin: *dy}
}

func (dy *douyin) User() *user {
	return &user{douyin: *dy}
}

func (dy *douyin) VideoDouYin() *videoDouYin {
	return &videoDouYin{douyin: *dy}
}

func (dy *douyin) VideoTouTiao() *videoTouTiao {
	return &videoTouTiao{douyin: *dy}
}

func (dy *douyin) VideoXiGua() *videoXiGua {
	return &videoXiGua{douyin: *dy}
}

func (dy *douyin) Comment() *comment {
	return &comment{douyin: *dy}
}

func (dy *douyin) Search() *search {
	return &search{douyin: *dy}
}

func (dy *douyin) DataExternal() *dataExternal {
	return &dataExternal{douyin: *dy}
}
