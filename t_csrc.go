package sa_xorm_model

type TCsrc struct {
	CsrcCnName   string `json:"csrc_cn_name" xorm:"unique VARCHAR(1024)"`
	IndustryCode string `json:"industry_code" xorm:"VARCHAR(32)"`
	CsrcId       string `json:"csrc_id" xorm:"not null pk VARCHAR(128)"`
}
