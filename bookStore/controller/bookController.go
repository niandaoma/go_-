package controller

import (
	"demo/bookStore/dao"
	"demo/bookStore/model"
	"demo/bookStore/utils"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

//首页图片
func IndexHandle(w http.ResponseWriter, r *http.Request){
	pageNoStr:=r.FormValue("pageNo")
	if pageNoStr==""{
		pageNoStr="1"
	}
	pageNo,_:=strconv.Atoi(pageNoStr)
	if pageNo<1 {
		pageNo=1
	}
	page, err := dao.GetPageBooks(pageNo)
	if err != nil {
		fmt.Println(err)
	}
	//判断是否已经登陆
	login, session := utils.IsLogin(r)
	if login {
		//已经登陆
		page.IsLogin=true
		page.Username=session.UserName
	}
	//解析模板
	t:=template.Must(template.ParseFiles("./views/index.html"))
	//执行
	t.Execute(w,page)
}

//获取所有图书
func GetBooks(w http.ResponseWriter, r *http.Request){
	//调用dao层GetBooks函数
	books, err := dao.GetBooks()
	if err != nil {
		fmt.Println(err)
	}
	t:=template.Must(template.ParseFiles("./views/pages/manager/book_manager.html"))
	t.Execute(w,books)
}

//分页获取所有的图书
func GetPageBooks(w http.ResponseWriter, r *http.Request){
	pageNoStr:=r.FormValue("pageNo")
	if pageNoStr==""{
		pageNoStr="1"
	}
	pageNo,_:=strconv.Atoi(pageNoStr)
	if pageNo<1 {
		pageNo=1
	}
	page, err := dao.GetPageBooks(pageNo)
	if err != nil {
		fmt.Println(err)
	}

	t:=template.Must(template.ParseFiles("./views/pages/manager/book_manager.html"))
	t.Execute(w,page)
}


//根据价格查询所有图书
func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request){
	pageNoStr:=r.FormValue("pageNo")
	minPrice,_:=strconv.ParseFloat(r.FormValue("min"),64)
	maxPrice,_:=strconv.ParseFloat(r.FormValue("max"),64)
	if pageNoStr==""{
		pageNoStr="1"
	}
	pageNo,_:=strconv.Atoi(pageNoStr)
	if pageNo<1 {
		pageNo=1
	}
	page:=&model.Page{}
	if r.FormValue("min") == "" {
		//如果没有价格传入,获取所有图片信息
		page, _ = dao.GetPageBooks(pageNo)
	}else{
		//如果有价格传入，查询价格区间期间图片信息
		page, _ = dao.GetPageBooksByPrice(minPrice,maxPrice,pageNo)
	}
	//判断是否已经登陆
	login, session := utils.IsLogin(r)
	if login {
			//已经登陆
			page.IsLogin=true
			page.Username=session.UserName
		}

	t:=template.Must(template.ParseFiles("./views/index.html"))
	t.Execute(w,page)

}

//添加图书
func AddBook(w http.ResponseWriter, r *http.Request){
	book:=&model.Book{}
	book.Title=r.PostFormValue("title")
	book.Author=r.PostFormValue("author")
	book.Price, _ = strconv.ParseFloat(r.FormValue("price"),64)
	book.Sales, _ = strconv.Atoi(r.PostFormValue("sales"))
	book.Stock, _ = strconv.Atoi(r.PostFormValue("stock"))

	err := dao.AddBook(book)
	if err != nil {
		fmt.Println(err)
	}
	GetPageBooks(w,r)

}

//删除图书
func DeleteBook(w http.ResponseWriter, r *http.Request){
	//从url中获取id值
	id,_:=strconv.Atoi(r.FormValue("bookId"))
	err := dao.DeleteBook(id)
	if err != nil {
		fmt.Println(err)
	}
	GetPageBooks(w,r)
}

//根据id获取图书
func GetBookById(w http.ResponseWriter, r *http.Request){
	id,_:=strconv.Atoi(r.FormValue("bookId"))
	book:=dao.GetBookById(id)
	t:=template.Must(template.ParseFiles("./views/pages/manager/book_edit.html"))
	t.Execute(w,book)
}

//更新图书信息
func UpdateBook(w http.ResponseWriter, r *http.Request){
	book:=&model.Book{}
	book.Id,_=strconv.Atoi(r.FormValue("bookId"))
	book.Title=r.FormValue("title")
	book.Author=r.FormValue("author")
	book.Price, _ = strconv.ParseFloat(r.FormValue("price"),64)
	fmt.Println(book.Price)
	book.Sales, _ = strconv.Atoi(r.FormValue("sales"))
	book.Stock, _ = strconv.Atoi(r.FormValue("stock"))
	err := dao.UpdateBook(book)
	fmt.Println(book)
	if err != nil {
		fmt.Println(err)
	}
	GetPageBooks(w,r)
}

//去更新或者添加图书的页面
func ToUpdateBookPage(w http.ResponseWriter, r *http.Request){
	bookId,_:=strconv.Atoi(r.FormValue("bookId"))
	book:=dao.GetBookById(bookId)
	if book.Id>0{
		//在更新图书
		t:=template.Must(template.ParseFiles("./views/pages/manager/book_edit.html"))
		t.Execute(w,book)
	}else {
		//在添加图书
		t:=template.Must(template.ParseFiles("./views/pages/manager/book_edit.html"))
		t.Execute(w,"")
	}
}
