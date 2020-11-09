package goo_open_douyin

import (
	"fmt"
	"strings"
	"testing"
)

func TestOauth_PlatformConnect(t *testing.T) {
	urlstr := dy.OAuth().PlatformConnect(strings.Join(scopes, ","), "", oauth_connect_redirect_url)
	fmt.Println(urlstr)
}

func TestOauth_AccessTokenDouYin(t *testing.T) {
	at, err := dy.OAuth().AccessTokenDouYin(code)
	fmt.Println(at, err)
}
