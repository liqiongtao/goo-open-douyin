package goo_open_douyin

import (
	"fmt"
	"testing"
)

func TestDataExternal_UserItem(t *testing.T) {
	rst, err := dy.DataExternal().UserItem(openId, accessToken, 30)
	fmt.Println(rst.Data.ResultList, err)
}

func TestDataExternal_ItemLike(t *testing.T) {
	rst, err := dy.DataExternal().ItemLike(openId, accessToken, itemId, 15)
	fmt.Println(rst.Data.ResultList, err)
}
