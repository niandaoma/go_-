package dao

import (
	"demo/bookStore/model"
	"fmt"
)

//获取所有图书
func GetBooks()(books []*model.Book,err error){
	stmt, err := DB.Prepare("select id,title,author,price,sales,stock,img_path from books ")
	if err != nil {
		fmt.Println("GetBooks err :",err)
	}
	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("GetBooks err :",err)
	}
	//用for 不是if！！！！！！
	for rows.Next() {
		book:=&model.Book{}
		rows.Scan(&book.Id,&book.Title,&book.Author,&book.Price,&book.Sales,&book.Stock,&book.ImgPath)
		books=append(books,book)
	}
	return
}


//添加图书
func AddBook(book *model.Book)error{
	stmt, err := DB.Prepare("insert into books(title,author,price,sales,stock)values (?,?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(book.Title, book.Author, book.Price, book.Sales, book.Stock)
	if err != nil {
		return err
	}
	return nil
}

//删除图书
func DeleteBook(id int)error{
	stmt, err := DB.Prepare("delete from books where id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

//根据id获取图书
func GetBookById(id int)*model.Book{
	book:=&model.Book{}
	stmt, err := DB.Prepare("select * from books where id = ?")
	if err != nil {
		fmt.Println(err)
	}
	row := stmt.QueryRow(id)
	row.Scan(&book.Id,&book.Title,&book.Author,&book.Price,&book.Sales,&book.Stock,&book.ImgPath)
	return book
}

//更新图书信息
func UpdateBook(book *model.Book)error{
	stmt, err := DB.Prepare("update books set title=?,author=?,price=?,sales=?,stock=? where id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(book.Title, book.Author, book.Price, book.Sales, book.Stock, book.Id)
	if err != nil {
		return err
	}
	return nil
}

//获取带分页的图书信息
func GetPageBooks(pageNo int)(page *model.Page,err error){
	page = &model.Page{}
	page.PageNo=pageNo
	//获取数据库中图书的总记录值
	stmt, err := DB.Prepare("select count(*) from books")
	if err != nil {
		return
	}
	row := stmt.QueryRow()
	row.Scan(&page.TotalRecord)
	//设置每页显示的条数为4
	page.PageSize=4
	//接受总页数
	if page.TotalRecord%page.PageSize==0{
		page.TotalPageNo=page.TotalRecord/page.PageSize
	}else {
		page.TotalPageNo=page.TotalRecord/page.PageSize+1
	}

	//获取当前页的图书信息
	stmt, err = DB.Prepare("select id,title,author,price,sales,stock,img_path from books limit ?, ?")
	if err != nil {
		return
	}
	rows, err := stmt.Query((page.PageNo-1)*page.PageSize,page.PageSize)
	if err != nil {
		return
	}
	//用for 不是if！！！！！！
	for rows.Next() {
		book:=&model.Book{}
		rows.Scan(&book.Id,&book.Title,&book.Author,&book.Price,&book.Sales,&book.Stock,&book.ImgPath)
		page.Books= append(page.Books, book)
	}
	return
}

//根据价格区间查询书籍
func GetPageBooksByPrice(minPrice float64,maxPrice float64,pageNo int)(page *model.Page,err error){
	page = &model.Page{}
	page.MinPrice=minPrice
	page.MaxPrice=maxPrice
	page.PageNo=pageNo
	//获取数据库中图书的总记录值
	stmt, err := DB.Prepare("select count(*) from books where price between ? and ?")
	if err != nil {
		return
	}
	row := stmt.QueryRow(minPrice,maxPrice)
	row.Scan(&page.TotalRecord)
	//设置每页显示的条数为4
	page.PageSize=4
	//接受总页数
	if page.TotalRecord%page.PageSize==0{
		page.TotalPageNo=page.TotalRecord/page.PageSize
	}else {
		page.TotalPageNo=page.TotalRecord/page.PageSize+1
	}

	//获取当前页的图书信息
	stmt, err = DB.Prepare("select id,title,author,price,sales,stock,img_path from books where price between ? and ? limit ?, ?")
	if err != nil {
		return
	}
	rows, err := stmt.Query(minPrice,maxPrice,(page.PageNo-1)*page.PageSize,page.PageSize)
	if err != nil {
		return
	}
	//用for 不是if！！！！！！
	for rows.Next() {
		book:=&model.Book{}
		rows.Scan(&book.Id,&book.Title,&book.Author,&book.Price,&book.Sales,&book.Stock,&book.ImgPath)
		page.Books= append(page.Books, book)
	}
	return
}