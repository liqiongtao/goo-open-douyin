package goo_open_douyin

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	ui, err := dy.User().UserInfoDouYin(openId, accessToken)
	fmt.Println(ui, err)
}

func TestFansList(t *testing.T) {
	fl, err := dy.User().FansList(openId, accessToken, 0, 20)
	fmt.Println(fl, err)
}

func TestFollowingList(t *testing.T) {
	fl, err := dy.User().FollowingList(openId, accessToken, 0, 20)
	fmt.Println(fl, err)
}
