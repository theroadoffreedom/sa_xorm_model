package sa_xorm_model

type TAstockFinanceReportData struct {
	Data           string `json:"data" xorm:"VARCHAR(128)"`
	ItemId         string `json:"item_id" xorm:"not null pk VARCHAR(128)"`
	ReportId       string `json:"report_id" xorm:"not null pk VARCHAR(128)"`
	CheckTime      int    `json:"check_time" xorm:"INT(10)"`
	CreateTime     int    `json:"create_time" xorm:"INT(10)"`
	ReportType     int    `json:"report_type" xorm:"INT(10)"`
	ReportTimeType int    `json:"report_time_type" xorm:"INT(10)"`
}
