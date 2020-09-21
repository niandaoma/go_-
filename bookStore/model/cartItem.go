package model

//购物项结构体
type CartItem struct {
	CartItemId string  //购物项的id
	Book       *Book   //图书信息
	Count      int     //图书数量
	Amount     float64 //图书总价格
	CartId     string  //购物车的id
}

//获取Amount
func (cartItem *CartItem) GetAmount() float64 {
	price := cartItem.Book.Price
	return float64(cartItem.Count) * price
}
