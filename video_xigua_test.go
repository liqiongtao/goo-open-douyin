package goo_open_douyin

import (
	"fmt"
	"os"
	"testing"
)

func TestXiGuaVideoUpload(t *testing.T) {
	f, err := os.Open(videoFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()

	vu, err := dy.VideoXiGua().Upload(openId, accessToken, f.Name(), f)
	fmt.Println(vu, err)

	vc, err := dy.VideoDouYin().Create(openId, accessToken, vu.Data.Video.VideoId, "", nil)
	fmt.Println(vc, err)
}

func TestXiGuaVedioPartUpload(t *testing.T) {
	pi, err := dy.VideoDouYin().PartInit(openId, accessToken)
	fmt.Println(pi, err)

	f, err := os.Open(videoFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()

	pu, err := dy.VideoDouYin().PartUpload(openId, accessToken, f.Name(), f, pi.Data.UploadId, 1)
	fmt.Println(pu, err)

	pc, err := dy.VideoDouYin().PartComplete(openId, accessToken, pi.Data.UploadId)
	fmt.Println(pc, err)

	vc, err := dy.VideoDouYin().Create(openId, accessToken, pc.Data.Video.VideoId, "", nil)
	fmt.Println(vc, err)
}

func TestXiGuaVideoList(t *testing.T) {
	vl, err := dy.VideoDouYin().List(openId, accessToken, 0, 20)
	fmt.Println(vl, err)
}

func TestXiGuaVideoData(t *testing.T) {
	vd, err := dy.VideoDouYin().Data(openId, accessToken, []string{itemId})
	fmt.Println(vd, err)
}
