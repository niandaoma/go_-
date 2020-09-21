package dao

import (
	"demo/bookStore/model"
)

//像数据库中插入购物项
func AddCartItem(item *model.CartItem) error {
	stmt, err := DB.Prepare("insert into cart_items(count,amount,book_id,cart_id) values(?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(item.Count, item.GetAmount(), item.Book.Id, item.CartId)
	if err != nil {
		return err
	}
	return nil
}

//更新购物项
func UpdateCartItem(item *model.CartItem) error {
	stmt, err := DB.Prepare("update cart_items set  count =?,amount=? where id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(item.Count, item.GetAmount(), item.CartItemId)
	if err != nil {
		return err
	}
	return nil
}

//根据图书id和购物车id来获取一辆购物车中某一本图书对应的购物项
func GetCartItemByBookIdAndCartId(bookId int, CartId string) (*model.CartItem, error) {
	book := GetBookById(bookId)
	cartItem := &model.CartItem{
		CartId: CartId,
		Book:   book,
	}
	stmt, err := DB.Prepare("select id,count,amount from cart_items where book_id =? and cart_id =?")
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(bookId, CartId)
	err = row.Scan(&cartItem.CartItemId, &cartItem.Count, &cartItem.Amount)
	if err != nil {
		return nil, err
	}

	return cartItem, nil
}

//根据购物车id来获取购物车中所有购物项
func GetCartItemByCartId(CartId string) ([]*model.CartItem, error) {
	cartItems := []*model.CartItem{}
	stmt, err := DB.Prepare("select id,count,amount,book_id from cart_items where cart_id =?")
	if err != nil {
		return nil, err
	}

	row, _ := stmt.Query(CartId)
	for row.Next() {
		cartItem := &model.CartItem{
			CartId: CartId,
		}
		var bookId int
		err = row.Scan(&cartItem.CartItemId, &cartItem.Count, &cartItem.Amount, &bookId)
		if err != nil {
			return nil, err
		}
		//用bookId查到的值来获取对应的book来赋给购物项
		cartItem.Book = GetBookById(bookId)
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil
}

//根据购物项id来获取购物项
func GetCartItemByCartItemId(CartItemId string) (*model.CartItem, error) {
	cartItem := &model.CartItem{
		CartId: CartItemId,
	}
	stmt, err := DB.Prepare("select id,count,amount,book_id from cart_items where id =?")
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(CartItemId)
	var bookId int
	err = row.Scan(&cartItem.CartItemId, &cartItem.Count, &cartItem.Amount, &bookId)
	if err != nil {
		return nil, err
	}
	//用bookId查到的值来获取对应的book来赋给购物项
	cartItem.Book = GetBookById(bookId)

	return cartItem, nil
}

//根据购物车id删除购物项
func DeleteCartItemByCartId(CartId string) error {
	stmt, err := DB.Prepare("delete from cart_items where cart_id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(CartId)
	if err != nil {
		return err
	}
	return nil
}

//根据购物项id删除购物项
func DeleteCartItemByCartItemId(CartItemId string) error {
	stmt, err := DB.Prepare("delete from cart_items where id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(CartItemId)
	if err != nil {
		return err
	}
	return nil
}
