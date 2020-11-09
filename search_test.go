package goo_open_douyin

import (
	"github.com/liqiongtao/goo"
	"testing"
)

func TestSearch_List(t *testing.T) {
	si, err := dy.Search().List(openId, accessToken, keyword, 0, 10)
	goo.Log.Debug(si, err)
}

func TestSearch_CommentList(t *testing.T) {
	cl, err := dy.Search().CommentList(accessToken, secItemId, 0, 10)
	goo.Log.Debug(cl, err)
}

func TestSearch_CommentReply(t *testing.T) {
	cr, err := dy.Search().CommentReply(openId, accessToken, secItemId, commentId, content)
	goo.Log.Debug(cr, err)
}

func TestSearch_CommentReplyList(t *testing.T) {
	crl, err := dy.Search().CommentReplyList(accessToken, secItemId, commentId, 0, 10)
	goo.Log.Debug(crl, err)
}
