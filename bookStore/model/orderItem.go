package model

//订单项
type OrderItem struct {
	OrderItemId int     //订单项的id
	Count       int     //订单项中图书的数量
	Amount      float64 //订单项下中图书的总金额
	Title       string  //订单项图书的书名
	Author      string  //订单项图书的作者
	Price       float64 //订单项图书的价格
	ImgPath     string  //订单项图书的图片地址
	OrderId     string  //订单项所属的订单id
	IsEnough    bool    //判断订单是否合法
}
