package controller

import (
	"demo/bookStore/dao"
	"demo/bookStore/model"
	"demo/bookStore/utils"
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

//创建订单
func CheckOut(w http.ResponseWriter, r *http.Request) {
	//通过session得到userId,userName
	_, session := utils.IsLogin(r)
	userId := session.UserId
	userName := session.UserName
	//通过uuid创建orderId
	u, _ := uuid.NewV4()
	orderId := u.String()
	timeStr:=time.Now().Format("2006-01-02 15:04:05")
	order := &model.Order{
		OrderId:     orderId,
		CreateTime:  timeStr,
		TotalCount:  0,
		TotalAmount: 0,
		State:       0,
		UserId:      userId,
		UserName:    userName,
	}

	//通过用户id获取购物车id
	cart, err := dao.GetCartByUserId(userId)
	if err != nil {
		fmt.Println(err)
	}
	order.TotalCount = cart.TotalCount
	order.TotalAmount = cart.TotalAmount
	//添加订单
	err = dao.AddOrder(order)
	if err != nil {
		fmt.Println(err)
	}
	//通过购物车id获取购物项的信息
	cartItems, err := dao.GetCartItemByCartId(cart.CartId)
	if err != nil {
		fmt.Println(err)
	}

	//添加订单项
	for _, cartItem := range cartItems {
		//创建订单项
		orderItem := &model.OrderItem{
			Count:   cartItem.Count,
			Amount: cartItem.Amount,
			Title:   cartItem.Book.Title,
			Author: cartItem.Book.Author,
			Price:   cartItem.Book.Price,
			ImgPath: cartItem.Book.ImgPath,
			OrderId: orderId,
		}
		err := dao.AddOrderItem(orderItem)
		if err != nil {
			fmt.Println(err)
		}
		//订单项添加之后，更新图书库存和销量信息
		cartItem.Book.Sales=cartItem.Book.Sales+cartItem.Count
		cartItem.Book.Stock=cartItem.Book.Stock-cartItem.Count
		//更新图书
		err = dao.UpdateBook(cartItem.Book)
		if err != nil {
			fmt.Println(err)
		}
	}

	//添加完订单之后清空购物车
	err = dao.DeleteCartByUserId(userId)
	if err != nil {
		fmt.Println(err)
	}

	t := template.Must(template.ParseFiles("./views/pages/cart/checkout.html"))
	t.Execute(w, order)

}


//获取某一用户的订单
func GetMyOrders(w http.ResponseWriter, r *http.Request){
	//通过session获取userId
	_, session := utils.IsLogin(r)
	userId:=session.UserId
	orders, err := dao.GetOrdersByUserId(userId)
	if err != nil {
		fmt.Println(err)
	}
	session.Orders=orders
	t := template.Must(template.ParseFiles("./views/pages/order/order.html"))
	t.Execute(w, session)

}

//通过订单id获取订单详情
func GetOrderInfo(w http.ResponseWriter, r *http.Request)  {
	//通过url获得orderId
	orderId := r.FormValue("orderId")
	orderItems, err := dao.GetOrderItemsByOrderId(orderId)
	if err != nil {
		fmt.Println(err)
	}
	t := template.Must(template.ParseFiles("./views/pages/order/order_info.html"))
	t.Execute(w, orderItems)
}

//获取所有订单
func GetOrders(w http.ResponseWriter, r *http.Request){
	//userId:=session.UserId
	orders, err := dao.GetOrders()
	if err != nil {
		fmt.Println(err)
	}
	t := template.Must(template.ParseFiles("./views/pages/order/order_manager.html"))
	t.Execute(w, orders)

}

//改变订单状态
func ChangeState(w http.ResponseWriter, r *http.Request){
	//通过url获取orderId
	orderId:=r.FormValue("orderId")
	state,_:=strconv.Atoi(r.FormValue("state"))
	err := dao.UpdateOrderState(orderId, state)
	if err != nil {
		fmt.Println(err)
	}
	//判断是否登陆
	isLogin, _ := utils.IsLogin(r)
	if isLogin {
		GetMyOrders(w,r)
	}else {
		GetOrders(w,r)
	}

}