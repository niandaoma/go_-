package model

type Book struct {
	Id      int     //图书Id
	Title   string  //图书标题
	Author  string  //图书作者
	Price   float64 //图书价格
	Sales   int     //图书销量
	Stock   int     //图书库存
	ImgPath string  //图书图片地址
}
