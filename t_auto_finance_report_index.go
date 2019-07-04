package sa_xorm_model

type TAutoFinanceReportIndex struct {
	Id         string `json:"id" xorm:"not null pk VARCHAR(128)"`
	DataTime   int    `json:"data_time" xorm:"not null pk INT(10)"`
	CheckTime  int    `json:"check_time" xorm:"INT(10)"`
	CreateTime int    `json:"create_time" xorm:"INT(10)"`
	State      string `json:"state" xorm:"VARCHAR(128)"`
	ReportType string `json:"report_type" xorm:"not null pk VARCHAR(128)"`
}
