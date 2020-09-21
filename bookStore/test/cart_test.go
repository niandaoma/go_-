package test

import (
	"demo/bookStore/dao"
	"demo/bookStore/model"
	"fmt"
	"testing"
)


//测试增加购物车
func TestAddCart(t *testing.T) {
	book1:=&model.Book{
		Id: 1,
		Price: 24.9,
	}
	book2:=&model.Book{
		Id: 2,
		Price: 23.0,
	}
	cartItem1:=&model.CartItem{
		Book:       book1,
		Count:      2,
		CartId:     "123456789",
	}
	cartItem2:=&model.CartItem{
		Book:       book2,
		Count:      3,
		CartId:     "123456789",
	}
	cartItems:=[]*model.CartItem{}
	cartItems=append(cartItems,cartItem1)
	cartItems=append(cartItems,cartItem2)
	cart:=&model.Cart{
		CartId:      "123456789",
		CartItems:   cartItems,
		UserId:      9,
	}
	err := dao.AddCart(cart)
	if err != nil {
		fmt.Println(err)
	}
}

//测试获取购物项
func TestGetCartItemByCartId(t *testing.T) {
	cartItems, err := dao.GetCartItemByCartId("123456789")
	if err != nil {
		panic(err)
	}
	for i, item := range cartItems {
		fmt.Printf("第%v个购物项:%v",i,item)
	}
}

func TestGetCartItemByBookIdAndCartId(t *testing.T) {
	fmt.Println("*************")
	cartItem, err := dao.GetCartItemByBookIdAndCartId(1, "123456789")
	if err != nil {
		panic(err)
	}
	fmt.Println(cartItem)
}

