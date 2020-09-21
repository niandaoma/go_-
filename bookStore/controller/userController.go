package controller

import (
	"demo/bookStore/dao"
	"demo/bookStore/model"
	"demo/bookStore/utils"
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	"html/template"
	"log"
	"net/http"
)

//处理登陆
func Login(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	login, err := dao.Login(username, password)

	if err != nil {
		log.Println(err)
	}
	if !login {
		//如果登陆失败
		te:=template.Must(template.ParseFiles("./views/pages/user/login.html"))
		te.Execute(w,"用户名或密码错误")
		fmt.Println("登陆失败")
	} else {
		//如果登陆成功，转到登陆成功页面
		isLogin, _ := utils.IsLogin(r)
		user, _ := dao.GetUser(username)
		if isLogin {
			//已经登陆过了
		}else{

			//创建一个session
			v4, _ := uuid.NewV4()
			sessionId := v4.String()
			sess:=&model.Session{
				SessionId: sessionId,
				UserName:  user.Username,
				UserId:    user.Id,
			}
			//把session存入到数据库中
			dao.AddSession(sess)
			cookie:=http.Cookie{
				Name: "user",
				Value: sess.SessionId,
				HttpOnly: true,
			}
			http.SetCookie(w,&cookie)
		}


		ts := template.Must(template.ParseFiles("./views/pages/user/login_success.html"))
		ts.Execute(w, user)
	}

}

//处理注销
func Logout(w http.ResponseWriter, r *http.Request){
	cookie,_:=r.Cookie("user")
	if cookie!=nil {
	//删除数据库中session
	err := dao.DelSession(cookie.Value)
	if err != nil {
		fmt.Println(err)
	}
	//设置cookie立即失效
	cookie.MaxAge=-1
	//把修改之后的cookie发给浏览器
	http.SetCookie(w,cookie)
	}
	//删除session和cookie之后，去往首页
	IndexHandle(w,r)


}

//处理注册
func Regist(w http.ResponseWriter, r *http.Request) {
	//获取表单中的值
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")

	//编译模板引擎
	te := template.Must(template.ParseFiles("./views/pages/user/regist.html"))

	//调用dao层注册函数
	err := dao.Regist(username, password, email)
	if err != nil {
		log.Println(err)
		te.Execute(w, "用户已存在，注册失败")
	} else {
		user, _ := dao.GetUser(username)

			//创建一个session
			v4, _ := uuid.NewV4()
			sessionId := v4.String()
			sess:=&model.Session{
				SessionId: sessionId,
				UserName:  user.Username,
				UserId:    user.Id,
			}
			//把session存入到数据库中
			dao.AddSession(sess)
			cookie:=http.Cookie{
				Name: "user",
				Value: sess.SessionId,
				HttpOnly: true,
			}
			http.SetCookie(w,&cookie)
		ts := template.Must(template.ParseFiles("./views/pages/user/regist_success.html"))
		ts.Execute(w, user)
	}
}

//检测用户名是否可用
func CheckUserName(w http.ResponseWriter, r *http.Request){
	username := r.PostFormValue("username")
	user, _ := dao.GetUser(username)
	if user!=nil{
		//说明用户存在
		w.Write([]byte("用户名已存在"))
	}else {
		//用户不存在
		w.Write([]byte("用户名可用"))
	}
}