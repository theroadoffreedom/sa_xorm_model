package sa_xorm_model

type TXjllbByYear struct {
	Id       string `json:"id" xorm:"not null pk VARCHAR(128)"`
	DataTime int    `json:"data_time" xorm:"not null pk INT(10)"`
	Data     string `json:"data" xorm:"JSON"`
}
