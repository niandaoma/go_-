package dao

import "demo/bookStore/model"

//添加session
func AddSession(session *model.Session)error{
	stmt, err := DB.Prepare("insert into sessions values (?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(session.SessionId, session.UserName, session.UserId)
	if err != nil {
		return err
	}
	return nil
}

//删除session

func DelSession(id string)error{
	stmt, err := DB.Prepare("delete from sessions where session_id =?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

//根据sessionId获取session
func GetSession(sessionId string)(*model.Session,error){
	sess:=&model.Session{}
	stmt, err := DB.Prepare("select session_id,username,user_id from sessions where session_id = ?")
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(sessionId)
	row.Scan(&sess.SessionId,&sess.UserName,&sess.UserId)
	return sess,nil
}