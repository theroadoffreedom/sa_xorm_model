package sa_xorm_model

type TAutoFinanceReportIndex struct {
	ReportId       string `json:"report_id" xorm:"not null pk VARCHAR(128)"`
	Id             string `json:"id" xorm:"VARCHAR(128)"`
	DataTime       int    `json:"data_time" xorm:"INT(10)"`
	CheckTime      int    `json:"check_time" xorm:"INT(10)"`
	CreateTime     int    `json:"create_time" xorm:"INT(10)"`
	State          int    `json:"state" xorm:"INT(10)"`
	ReportType     int    `json:"report_type" xorm:"INT(10)"`
	ReportTimeType int    `json:"report_time_type" xorm:"INT(10)"`
}
