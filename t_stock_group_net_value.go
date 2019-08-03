package sa_xorm_model

type TStockGroupNetValue struct {
	GroupId            string `json:"group_id" xorm:"not null pk VARCHAR(128)"`
	HoldId             string `json:"hold_id" xorm:"not null pk VARCHAR(128)"`
	NetValue           string `json:"net_value" xorm:"default 1 DECIMAL(10,2)"`
	NetValueUpdateTime int    `json:"net_value_update_time" xorm:"INT(10)"`
}
