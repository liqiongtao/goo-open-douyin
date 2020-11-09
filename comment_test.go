package goo_open_douyin

import (
	"fmt"
	"testing"
)

func TestComment_List(t *testing.T) {
	li, err := dy.Comment().List(openId, accessToken, itemId, 0, 20)
	fmt.Println(li, err)
}

func TestComment_ReplyList(t *testing.T) {
	rl, err := dy.Comment().ReplyList(openId, accessToken, itemId, commentId, 0, 20)
	fmt.Println(rl, err)
}

func TestComment_Reply(t *testing.T) {
	re, err := dy.Comment().Reply(openId, accessToken, itemId, commentId, "挺棒的!!!!!!!")
	fmt.Println(re, err)
}
