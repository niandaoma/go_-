package test

import (
	"demo/bookStore/dao"
	"demo/bookStore/model"
	"fmt"
	"testing"
)

//测试添加session
func TestAddSession(t *testing.T) {
	sess:=&model.Session{
		SessionId: "123",
		UserName:  "nihao",
		UserId:    1,
	}
	err := dao.AddSession(sess)
	if err != nil {
		fmt.Println(err)
	}
}

//测试删除session
func TestDelSession(t *testing.T) {
	err := dao.DelSession("123")
	if err != nil {
		fmt.Println(err)
	}
}
