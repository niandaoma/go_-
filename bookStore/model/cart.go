package model

//购物车
type Cart struct {
	CartId      string      //购物车Id
	CartItems   []*CartItem //购物项
	TotalCount  int         //购物车中图书总数量
	TotalAmount float64     //购物车中图书总价格
	UserName    string      //购物车所属用户姓名
	UserId      int         //购物车所属用户Id
	IsNull      bool        //判断购物车是否为空
}

//获取TotalCount图书总数量
func (cart *Cart) GetTotalCount() int {
	totalCount := 0
	for _, v := range cart.CartItems {
		totalCount = totalCount + v.Count
	}
	return totalCount
}

//获取TotalAmount图书总价格
func (cart *Cart) GetTotalAmount() float64 {
	totalAmount := 0.0
	for _, v := range cart.CartItems {
		totalAmount = totalAmount + v.GetAmount()
	}
	return totalAmount
}
