package goo_open_douyin

import (
	"fmt"
	"github.com/liqiongtao/goo"
	"os"
	"testing"
)

func TestTouTiaoVideoUpload(t *testing.T) {
	f, err := os.Open(videoFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()

	vu, err := dy.VideoTouTiao().Upload(openId, accessToken, f.Name(), f)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	goo.Log.Debug(vu.Data)

	vc, err := dy.VideoTouTiao().Create(openId, accessToken, vu.Data.Video.VideoId, "", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	goo.Log.Debug(vc.Data)
}

func TestTouTiaoVedioPartUpload(t *testing.T) {
	pi, err := dy.VideoTouTiao().PartInit(openId, accessToken)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	goo.Log.Debug(pi.Data)

	f, err := os.Open(videoFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()

	pu, err := dy.VideoTouTiao().PartUpload(openId, accessToken, f.Name(), f, pi.Data.UploadId, 1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	goo.Log.Debug(pu.Data)

	pc, err := dy.VideoTouTiao().PartComplete(openId, accessToken, pi.Data.UploadId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	goo.Log.Debug(pc.Data)

	vc, err := dy.VideoTouTiao().Create(openId, accessToken, pc.Data.Video.VideoId, "", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	goo.Log.Debug(vc.Data)
}

func TestTouTiaoVideoList(t *testing.T) {
	vl, err := dy.VideoTouTiao().List(openId, accessToken, 0, 20)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	goo.Log.Debug(vl.Data)
}

func TestTouTiaoVideoData(t *testing.T) {
	vd, err := dy.VideoTouTiao().Data(openId, accessToken, []string{itemId})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	goo.Log.Debug(vd.Data)
}
