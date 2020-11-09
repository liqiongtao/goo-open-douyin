package goo_open_douyin

var (
	dyc = &douyin{}

	Oauth        *oauth
	User         *user
	Search       *search
	Comment      *comment
	DataExternal *dataExternal
	Video        *videoDouYin
	VideoXiGua   *videoXiGua
	VideoTouTiao *videoTouTiao
)

func Init(clientKey, clientSecret string) {
	dyc = New(clientKey, clientSecret)

	Oauth = dyc.OAuth()
	User = dyc.User()
	Search = dyc.Search()
	Comment = dyc.Comment()
	DataExternal = dyc.DataExternal()
	Video = dyc.VideoDouYin()
	VideoXiGua = dyc.VideoXiGua()
	VideoTouTiao = dyc.VideoTouTiao()
}
