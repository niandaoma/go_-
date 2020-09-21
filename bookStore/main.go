package main

import (
	"demo/bookStore/controller"
	"net/http"
)

func main() {
	//设置处理静态资源
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("./views/pages"))))
	http.HandleFunc("/index", controller.IndexHandle)

	//处理登陆问题
	{
		//处理用户登陆
		http.HandleFunc("/login", controller.Login)
		//处理用户注销
		http.HandleFunc("/logout", controller.Logout)
		//处理用户注册
		http.HandleFunc("/regist", controller.Regist)
		//检测用户名是否可用
		http.HandleFunc("/checkUserName", controller.CheckUserName)
	}

	//处理图书
	{
		//获取所有图书
		http.HandleFunc("/getBooks", controller.GetBooks)
		//获取带分页的图书
		http.HandleFunc("/getPageBooks", controller.GetPageBooks)
		//添加图书
		http.HandleFunc("/addBook", controller.AddBook)
		//删除图书
		http.HandleFunc("/deleteBook", controller.DeleteBook)
		//根据id获得某本书的信息
		http.HandleFunc("/GetBookById", controller.GetBookById)
		//修改图书的信息
		http.HandleFunc("/updateBook", controller.UpdateBook)
		//更新或者添加图书
		http.HandleFunc("/ToUpdateBookPage", controller.ToUpdateBookPage)
		//根据价格区间查询图书
		http.HandleFunc("/getPageBooksByPrice", controller.GetPageBooksByPrice)
	}


	//处理购物车
	{
		//将图书添加到购物车
		http.HandleFunc("/addBook2Cart",controller.AddBook2Cart)
		//从数据库中获取购物车
		http.HandleFunc("/getCart",controller.GerCartInfo)
		//清空购物车
		http.HandleFunc("/deleteCart",controller.DeleteCart)
		//删除购物项
		http.HandleFunc("/deleteCartItem",controller.DeleteCartItemByCartItemId)
		//更新购物项
		http.HandleFunc("/updateCartItem",controller.UpdateCartItem)
	}

	//订单项
	{
		//创建订单
		http.HandleFunc("/checkOut",controller.CheckOut)
		//获取我的订单
		http.HandleFunc("/getMyOrder",controller.GetMyOrders)
		//通过订单Id获取订单详情
		http.HandleFunc("/getOrderInfo",controller.GetOrderInfo)
		//获取所有订单
		http.HandleFunc("/getOrders",controller.GetOrders)
		//改变订单状态
		http.HandleFunc("/changeState",controller.ChangeState)
	}

	http.ListenAndServe(":8080", nil)
}
