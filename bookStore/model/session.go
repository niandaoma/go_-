package model

type Session struct {
	SessionId string  //sessionId
	UserName  string  //用户姓名
	UserId    int     //用户Id
	Orders    []*Order //订单
	OrderItems []*OrderItem //订单项
}
