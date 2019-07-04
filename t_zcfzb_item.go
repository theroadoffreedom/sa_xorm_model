package sa_xorm_model

type TZcfzbItem struct {
	Id     string `json:"id" xorm:"not null pk VARCHAR(128)"`
	CnText string `json:"cn_text" xorm:"VARCHAR(1024)"`
	EnText string `json:"en_text" xorm:"VARCHAR(512)"`
	Unit   string `json:"unit" xorm:"VARCHAR(128)"`
}
