package sa_xorm_model

type TFinanceReportItemDefinition struct {
	Id       string `json:"id" xorm:"not null pk VARCHAR(128)"`
	Cn       string `json:"cn" xorm:"VARCHAR(128)"`
	En       string `json:"en" xorm:"VARCHAR(128)"`
	ItemType string `json:"item_type" xorm:"VARCHAR(128)"`
}
