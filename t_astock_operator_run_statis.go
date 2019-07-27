package sa_xorm_model

type TAstockOperatorRunStatis struct {
	StockId    string `json:"stock_id" xorm:"not null pk VARCHAR(128)"`
	OperatorId string `json:"operator_id" xorm:"not null pk VARCHAR(128)"`
	CheckTime  int    `json:"check_time" xorm:"INT(10)"`
	ShouldRun  string `json:"should_run" xorm:"not null VARCHAR(128)"`
}
