package sa_xorm_model

type TStockGroupHolding struct {
	HoldId     string `json:"hold_id" xorm:"not null pk VARCHAR(128)"`
	GroupId    string `json:"group_id" xorm:"not null pk VARCHAR(128)"`
	StockId    string `json:"stock_id" xorm:"not null pk VARCHAR(128)"`
	Percent    string `json:"percent" xorm:"DECIMAL(10,2)"`
	CreateTime int    `json:"create_time" xorm:"INT(10)"`
}
