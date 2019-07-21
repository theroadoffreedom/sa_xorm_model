package sa_xorm_model

type TAstockFinanceReportItemDefinition struct {
	Id         string `json:"id" xorm:"not null pk VARCHAR(128)"`
	Cn         string `json:"cn" xorm:"VARCHAR(128)"`
	Unit       string `json:"unit" xorm:"VARCHAR(128)"`
	ItemType   string `json:"item_type" xorm:"VARCHAR(128)"`
	CreateTime int    `json:"create_time" xorm:"INT(10)"`
}
