package sa_xorm_model

type TStockGroup struct {
	GroupId    string `json:"group_id" xorm:"not null pk VARCHAR(128)"`
	Name       string `json:"name" xorm:"not null VARCHAR(128)"`
	Desc       string `json:"desc" xorm:"TEXT"`
	Remark     string `json:"remark" xorm:"TEXT"`
	CreateTime int    `json:"create_time" xorm:"INT(10)"`
	DeleteTime int    `json:"delete_time" xorm:"INT(10)"`
	ModifyTime int    `json:"modify_time" xorm:"INT(10)"`
	HoldId     string `json:"hold_id" xorm:"not null VARCHAR(128)"`
}
