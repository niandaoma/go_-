package controller

import (
	"demo/bookStore/dao"
	"demo/bookStore/model"
	"demo/bookStore/utils"
	"encoding/json"
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	"html/template"
	"net/http"
	"strconv"
)

//将图书添加到购物车
func AddBook2Cart(w http.ResponseWriter, r *http.Request) {
	bookId, _ := strconv.Atoi(r.FormValue("bookId"))
	//根据bookId获得图书信息
	book := dao.GetBookById(bookId)
	//首先判断数据库中是否有了该购物车
	isLogin, session := utils.IsLogin(r)
	if isLogin {
		//如果已经登陆
		userId := session.UserId

		getCart, _ := dao.GetCartByUserId(userId)

		if getCart != nil {
			//说明有该购物车
			//判断是否有该图书的购物项
			cartItem, _ := dao.GetCartItemByBookIdAndCartId(bookId, getCart.CartId)

			if cartItem != nil {
				//说明有该图书的购物项
				//图书总数+1
				cartItem.Count++
				err := dao.UpdateCartItem(cartItem)
				if err != nil {
					panic(err)
				}
			} else {
				//没有该图书的购物项
				//创建一个购物项
				cartItem = &model.CartItem{
					Count:  1,
					Book:   book,
					CartId: getCart.CartId,
				}
				//加入新添加的购物项
				err := dao.AddCartItem(cartItem)
				if err != nil {
					panic(err)
				}
			}
			//存入购物项之后，再获取所有购物项
			cartItems, err := dao.GetCartItemByCartId(getCart.CartId)
			if err != nil {
				panic(err)
			}
			getCart.CartItems = cartItems
			//更新数据库中的购物车信息
			err = dao.UpdateCart(getCart)
			if err != nil {
				panic(err)
			}
		} else {
			//说明还没有该购物车
			//创建一个购物车并添加到数据库中
			//1.创建购物车
			v4, _ := uuid.NewV4()
			uuid := v4.String()
			cart := &model.Cart{
				CartId: uuid,
				UserId: userId,
			}
			//添加购物项
			err := dao.AddCart(cart)
			if err != nil {
				panic(err)
			}
			//2.创建购物项
			//根据图书id获取图书
			book := dao.GetBookById(bookId)
			cartItem := &model.CartItem{
				Count:  1,
				Book:   book,
				CartId: uuid,
			}
			//3.添加购物项
			err = dao.AddCartItem(cartItem)
			if err != nil {
				panic(err)
			}
			//存入购物项之后，再获取所有购物项
			cartItems, err := dao.GetCartItemByCartId(uuid)
			if err != nil {
				panic(err)
			}
			cart.CartItems = cartItems
			//更新数据库中的购物车信息
			err = dao.UpdateCart(cart)
			if err != nil {
				panic(err)
			}
		}
		w.Write([]byte("您刚刚将" + book.Title + "添加到了购物车"))
	} else {
		//如果没有登陆，请先登陆
		w.Write([]byte("请先登陆!"))
	}

}

//从数据库中获取购物车
func GerCartInfo(w http.ResponseWriter, r *http.Request) {
	//获取用户id
	_, session := utils.IsLogin(r)
	userId := session.UserId
	//根据用户id获取购物车
	cart, err := dao.GetCartByUserId(userId)
	if err != nil {
		fmt.Println(err)
	}
	if cart != nil {
		cart.UserName = session.UserName
		//有该购物车
		//for _, cartItem := range cartItems {
		//	cartItem.Book = dao.GetBookById(cartItem.Book.Id)
		//	cartItem.Amount=cartItem.GetAmount()
		//}
		//把购物项返回页面
		t := template.Must(template.ParseFiles("./views/pages/cart/cart.html"))
		t.Execute(w, cart)
	} else {
		//没有该购物车
		cart := &model.Cart{
			UserName: session.UserName,
			IsNull:   true,
		}
		t := template.Must(template.ParseFiles("./views/pages/cart/cart.html"))
		t.Execute(w, cart)
	}

}

//清空购物车
func DeleteCart(w http.ResponseWriter, r *http.Request) {
	_, session := utils.IsLogin(r)
	userId := session.UserId
	err := dao.DeleteCartByUserId(userId)
	if err != nil {
		fmt.Println(err)
	}
	GerCartInfo(w, r)
}

//删除购物项
func DeleteCartItemByCartItemId(w http.ResponseWriter, r *http.Request) {
	cartItemId := r.FormValue("cartItemId")
	err := dao.DeleteCartItemByCartItemId(cartItemId)
	if err != nil {
		fmt.Println(err)
	}
	//从session中获取UserId
	_, session := utils.IsLogin(r)
	userId := session.UserId

	//获取cart
	cart, err := dao.GetCartByUserId(userId)
	if err != nil {
		fmt.Println(err)
	}

	//存入购物项之后，再获取所有购物项
	cartItems, err := dao.GetCartItemByCartId(cart.CartId)
	if err != nil {
		fmt.Println(err)
	}

	if len(cartItems) != 0 {
		//如果该购物车中还有购物项
		cart.CartItems = cartItems
		//更新数据库中的购物车信息
		err = dao.UpdateCart(cart)
		if err != nil {
			panic(err)
		}
		GerCartInfo(w, r)
	} else  {
		//如果该购物车中没有购物项
		DeleteCart(w, r)
	}
}

//更新购物项
func UpdateCartItem(w http.ResponseWriter, r *http.Request) {

	//获取购物项id 和 count
	cartItemId := r.FormValue("cartItemId")
	bookCount, _ := strconv.Atoi(r.FormValue("bookCount"))

	//通过cartItemId来获取cartItem
	cartItem, err := dao.GetCartItemByCartItemId(cartItemId)
	if err != nil {
		fmt.Println(err)
	}
	cartItem.Count = bookCount
	cartItem.Amount = cartItem.GetAmount()
	//判断库存是否足够
	if cartItem.Book.Stock>=bookCount {
		//如果库存足够
		//通过cartItem更新cartItem
		err = dao.UpdateCartItem(cartItem)
		if err != nil {
			fmt.Println(err)
		}
		//从session中获取UserId
		_, session := utils.IsLogin(r)
		userId := session.UserId
		//获取cart
		cart, err := dao.GetCartByUserId(userId)
		if err != nil {
			fmt.Println(err)
		}

		//存入购物项之后，再获取所有购物项
		cartItems, err := dao.GetCartItemByCartId(cart.CartId)
		if err != nil {
			panic(err)
		}
		cart.CartItems = cartItems
		//更新数据库中的购物车信息
		err = dao.UpdateCart(cart)
		if err != nil {
			panic(err)
		}
		cart, _ = dao.GetCartByUserId(userId)
		data := model.Data{}
		for _, cartItem := range cart.CartItems {
			if cartItem.CartItemId==cartItemId {
				data = model.Data{
					Amount:      cartItem.Amount,
					TotalAmount: cart.TotalAmount,
					TotalCount:  cart.TotalCount,
				}
			}
		}
		json, err := json.Marshal(data)
		w.Write(json)
	}else{
		//如果库存不够
		GerCartInfo(w, r)
	}

}
