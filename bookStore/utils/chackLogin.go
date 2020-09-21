package utils

import (
	"demo/bookStore/dao"
	"demo/bookStore/model"
	"net/http"
)

//判断用户是否已经登陆
func IsLogin(r *http.Request)(bool,*model.Session){
	cookie,_:=r.Cookie("user")
	if cookie==nil {
		//没有登陆
		return false,nil
	}
		//已经登陆
		//获取cookie中的值
		cookieValue:=cookie.Value
		//数据库中查询对应session
		session,_:=dao.GetSession(cookieValue)
		if session.UserId>0 {
			//已经登陆
			return true,session
		}
		return false,nil
}
