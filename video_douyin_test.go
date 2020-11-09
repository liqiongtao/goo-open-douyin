package goo_open_douyin

import (
	"fmt"
	"github.com/liqiongtao/goo"
	"os"
	"testing"
)

func TestVideoUpload(t *testing.T) {
	f, err := os.Open(videoFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()

	vu, err := dy.VideoDouYin().Upload(openId, accessToken, f.Name(), f)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	goo.Log.Debug(vu.Data)

	vc, err := dy.VideoDouYin().Create(openId, accessToken, vu.Data.Video.VideoId, "", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	goo.Log.Debug(vc.Data)
}

func TestVedioPartUpload(t *testing.T) {
	pi, err := dy.VideoDouYin().PartInit(openId, accessToken)
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

	pu, err := dy.VideoDouYin().PartUpload(openId, accessToken, f.Name(), f, pi.Data.UploadId, 1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	goo.Log.Debug(pu.Data)

	pc, err := dy.VideoDouYin().PartComplete(openId, accessToken, pi.Data.UploadId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	goo.Log.Debug(pc.Data)

	vc, err := dy.VideoDouYin().Create(openId, accessToken, pc.Data.Video.VideoId, "", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	goo.Log.Debug(vc.Data)
}

func TestVideoList(t *testing.T) {
	vl, err := dy.VideoDouYin().List(openId, accessToken, 0, 20)
	fmt.Println(vl.Data.List, err)
}

func TestVideoData(t *testing.T) {
	vd, err := dy.VideoDouYin().Data(openId, accessToken, []string{itemId})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	goo.Log.Debug(vd.Data)
}

func TestVideoDelete(t *testing.T) {
	vd, err := dy.VideoDouYin().Delete(openId, accessToken, itemId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	goo.Log.Debug(vd.Data)
}

func TestImageUpload(t *testing.T) {
	f, err := os.Open(imageFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()

	iu, err := dy.VideoDouYin().ImageUpload(openId, accessToken, f.Name(), f)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	goo.Log.Debug(iu.Data)

	ic, err := dy.VideoDouYin().ImageCreate(openId, accessToken, iu.Data.Image.ImageId, "", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	goo.Log.Debug(ic.Data)
}
