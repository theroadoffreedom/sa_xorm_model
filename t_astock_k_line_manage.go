package sa_xorm_model

type TAstockKLineManage struct {
	StockId      string `json:"stock_id" xorm:"not null pk VARCHAR(128)"`
	DayKLineInit int    `json:"day_k_line_init" xorm:"default 0 INT(10)"`
}
