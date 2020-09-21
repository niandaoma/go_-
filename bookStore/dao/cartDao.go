package dao

import "demo/bookStore/model"

//向数据库中插入购物车
func AddCart(cart *model.Cart)error{
	stmt, err := DB.Prepare("insert into carts (id,total_count,total_amount,user_id)values (?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(cart.CartId, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserId)
	if err != nil {
		return err
	}
	for _, v := range cart.CartItems {
		err := AddCartItem(v)
		if err != nil {
			return err
		}
	}
	return nil
}

//更新购物车信息
func UpdateCart(cart *model.Cart)error{
	stmt, err := DB.Prepare("update carts set total_count = ? , total_amount = ? where id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(cart.GetTotalCount(), cart.GetTotalAmount(),cart.CartId)
	if err != nil {
		return err
	}
	return nil
}

//根据用户id从数据库中查询对应的购物车
func GetCartByUserId(userId int)(*model.Cart,error){
	cart:=&model.Cart{}
	stmt, err := DB.Prepare("select id,total_count,total_amount,user_id from carts where user_id=?")
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(userId)
	err = row.Scan(&cart.CartId, &cart.TotalCount, &cart.TotalAmount, &cart.UserId)
	if err != nil {
		return nil, err
	}
	//有购物车，把IsNull设置为false
	cart.IsNull=false
	cart.CartItems,_=GetCartItemByCartId(cart.CartId)
	return cart, nil
}

//根据用户id清空购物车
func DeleteCartByUserId(userId int)error{
	//用用户id获取购物车id
	cart, err := GetCartByUserId(userId)
	if err != nil {
		return err
	}
	//用购物车id清空购物项
	err = DeleteCartItemByCartId(cart.CartId)
	if err != nil {
		return err
	}
	//清空购物车
	stmt, err := DB.Prepare("delete from carts where user_id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(userId)
	if err != nil {
		return err
	}
	return nil
}
