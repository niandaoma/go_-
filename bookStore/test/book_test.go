package test

import (
	"demo/bookStore/dao"
	"demo/bookStore/model"
	"fmt"
	"testing"
)

func TestBooks(t *testing.T){
	t.Run("测试获取所有图书",TestGetBooks)
	t.Run("测试添加图书",TestAddBook)
	t.Run("测试删除图书",TestDeleteBook)
	t.Run("测试通过id查询图书",TestGetBookById)
	t.Run("测试更新图书信息",TestUpdateBook)
}

//测试获得所有图书
func TestGetBooks(t *testing.T) {
	books, _ := dao.GetBooks()
	for i, book := range books {
		fmt.Printf("第%v本书：%v\n",i+1,book)
	}
}

//测试添加图书
func TestAddBook(t *testing.T){
	book:=&model.Book{
		Title:   "测试书籍",
		Author:  "李大牛",
		Price:   999999,
		Sales:   999999,
		Stock:   999999,
	}
	dao.AddBook(book)
}

//测试删除图书
func TestDeleteBook(t *testing.T){
	err := dao.DeleteBook(34)
	if err != nil {
		fmt.Println(err)
	}
}

//测试通过id查询图书
func TestGetBookById(t *testing.T){
	fmt.Println(dao.GetBookById(5))
}

//测试更新图书信息
func TestUpdateBook(t *testing.T){
	book:=&model.Book{
		Id:      38,
		Title:   "ccesss",
		Author:  "ldn",
		Price:   56,
		Sales:   78,
		Stock:   66,
		ImgPath: "",
	}
	err := dao.UpdateBook(book)
	if err != nil {
		fmt.Println(err)
	}
}

//测试分页
func TestGetPageBooks(t *testing.T){
	page, err := dao.GetPageBooks(2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("总共有%v条记录\n",page.TotalRecord)
	fmt.Printf("当前页是第%v页\n",page.PageNo)
	fmt.Printf("总页数是%v\n",page.TotalPageNo)

	for i, book := range page.Books {
		fmt.Printf("第%v本书：%v\n",i+1,book)
	}
}