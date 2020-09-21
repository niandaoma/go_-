package dao

import "demo/bookStore/model"

//添加订单到数据库
func AddOrder(order *model.Order)error{
	stmt, err := DB.Prepare("insert into orders (id,create_time,total_count,total_amount,state,user_id) values (?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.OrderId, order.CreateTime, order.TotalCount, order.TotalAmount, order.State, order.UserId)
	if err != nil {
		return err
	}
	return nil
}


//获取某一用户所有订单
func GetOrdersByUserId(userId int)([]*model.Order,error){
	stmt, err := DB.Prepare("select * from orders where user_id=?")
	if err != nil {
		return nil,err
	}
	rows, err := stmt.Query(userId)
	if err != nil {
		return nil, err
	}
	orders:=[]*model.Order{}
	for rows.Next() {
		order:=&model.Order{}
		err := rows.Scan(&order.OrderId, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserId)
		if err != nil {
			return nil, err
		}
		orders=append(orders,order)
	}
	return orders,nil
}

//获取数据库中所有订单
func GetOrders()([]*model.Order,error){
	stmt, err := DB.Prepare("select * from orders")
	if err != nil {
		return nil,err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	orders:=[]*model.Order{}
	for rows.Next() {
		order:=&model.Order{}
		err := rows.Scan(&order.OrderId, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserId)
		if err != nil {
			return nil, err
		}
		orders=append(orders,order)
	}
	return orders,nil
}

//通过订单id获取订单详情
func GetOrderItemsByOrderId(orderId string)([]*model.OrderItem,error){
	stmt, err := DB.Prepare("select count,amount,title,author,price,img_path from order_items where order_id=?")
	if err != nil {
		return nil,err
	}
	rows, err := stmt.Query(orderId)
	if err != nil {
		return nil, err
	}
	orderItems:=[]*model.OrderItem{}
	for rows.Next() {
		orderItem:=&model.OrderItem{}
		err := rows.Scan(&orderItem.Count, &orderItem.Amount, &orderItem.Title, &orderItem.Author, &orderItem.Price, &orderItem.ImgPath)
		if err != nil {
			return nil, err
		}
		orderItems=append(orderItems,orderItem)
	}
	return orderItems,nil
}

//更新订单状态
func UpdateOrderState(orderId string,state int)error{
	stmt, err := DB.Prepare("update orders set state=? where id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(state, orderId)
	if err != nil {
		return err
	}
	return nil
}