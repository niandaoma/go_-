package dao

import "demo/bookStore/model"

//添加订单项到数据库
func AddOrderItem(orderItem *model.OrderItem)error{
	stmt, err := DB.Prepare("insert into order_items (count ,amount,title,author,price,img_path,order_id) values (?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(orderItem.Count,orderItem.Amount,orderItem.Title,
		orderItem.Author,orderItem.Price,orderItem.ImgPath,orderItem.OrderId)
	if err != nil {
		return err
	}
	return nil
}
