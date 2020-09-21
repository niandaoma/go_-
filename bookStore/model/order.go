package model

//订单Order结构
type Order struct {
	OrderId     string  //订单号
	CreateTime  string  //创建时间
	TotalCount  int     //订单中图书的总数量
	TotalAmount float64 //订单中图书的总金额
	State       int     //订单的状态 --> 0未发货1已发货2交易完成
	UserId      int     //订单所属用户id
	UserName    string  //订单所属用户姓名
}

//NoSend 未发货
func (order *Order) NoSend() bool {
	return order.State == 0
}

//SendComplete 已发货
func (order *Order) SendComplete() bool {
	return order.State == 1
}

//Complete 交易完成
func (order *Order) Complete() bool {
	return order.State == 2
}
