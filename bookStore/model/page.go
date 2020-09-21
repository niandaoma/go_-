package model

//分页
type Page struct {
	Books       []*Book //每页查询出的图片存放数组
	PageNo      int     //当前页
	PageSize    int     //每页显示的条数
	TotalPageNo int     //总页数，通过计算所得
	TotalRecord int     //总记录数，通过查询数据库得到
	MinPrice    float64 //最低价格
	MaxPrice    float64 //最高价格
	IsLogin     bool    //判断是否已经登陆
	Username    string
}

//判断是否有上一页
func (p *Page) IsHasPrev() bool {
	return p.PageNo > 1
}

//判断是否有下一页
func (p *Page) IsHasNext() bool {
	return p.PageNo < p.TotalPageNo
}

//判断是否有价格查询
func (p *Page) IsPrice() bool {
	return p.MinPrice == 0 || p.MaxPrice == 0
}

//获取上一页
func (p *Page) GetPrevPageNo() int {
	if p.IsHasPrev() {
		return p.PageNo - 1
	} else {
		return 1
	}
}

//获取下一页
func (p *Page) GetNextPageNo() int {
	if p.IsHasNext() {
		return p.PageNo + 1
	} else {
		return p.TotalPageNo
	}
}
