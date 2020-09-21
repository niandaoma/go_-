package dao

import (
	"database/sql"
	"demo/bookStore/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	DB *sql.DB
)

func init() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/bookstore")
	if err != nil {
		panic(err)
	}
	DB = db
}

//登陆
func Login(username string, password string) (bool, error) {
	pwd := ""
	stmt, err := DB.Prepare("select password from users where username=?")
	if err != nil {
		log.Println("数据库操作语句写错了")
		return false, err
	}
	defer stmt.Close()
	//如果用户名不存在，返回错误并打印
	result:= stmt.QueryRow(username)

	err = result.Scan(&pwd)
	if err != nil {
		return false, err
	}
	//检查密码是否正确
	if pwd != password {
		log.Println("密码不正确")
		return false, nil
	}
	return true, nil
}

//当通过登陆后，获取用户信息
func GetUser(username string) (*model.User, error) {
	user := &model.User{}
	stmt, err := DB.Prepare("select * from users where username=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	//如果用户名不存在，返回错误并打印
	result:= stmt.QueryRow(username)
	err = result.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//注册
func Regist(username string, password string, email string)(err error) {
	//先检查改用户是否已经存在
	user,_:=GetUser(username)
	if user!=nil {
		err=fmt.Errorf("该用户已存在")
		return err
	}
	stmt, err := DB.Prepare("insert into users (username,password,email) values (?,?,?)	")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, password, email)
	if err != nil {
		return err
	}
	return nil
}
