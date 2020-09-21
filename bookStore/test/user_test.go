package test

import (
	"demo/bookStore/dao"
	"fmt"
	"log"
	"testing"
)

func TestLogin(t *testing.T) {
	ok, err := dao.Login("ldn2", "123456")
	if err != nil {
		log.Println(err)
	}
	if !ok {
		log.Println("登陆失败")
	}else {
		log.Println("登陆成功")
	}
}

func TestGet(t *testing.T){
	user, err := dao.GetUser("ldn2")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(user)
}

func TestRegist(t *testing.T){
	err := dao.Regist("ldn2", "123456", "1342066923@qq.com")
	if err != nil {
		log.Println(err)
	}
}

func TestUser(t *testing.T){
	log.Println("测试user中的函数")
	t.Run("测试login",TestLogin)
	t.Run("测试get",TestGet)
	//t.Run("测试regist",TestRegist)
}