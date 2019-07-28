package sa_xorm_model

import (
	"time"
)

type TAstockDayKLineData struct {
	StockId string    `json:"stock_id" xorm:"not null pk VARCHAR(128)"`
	Day     time.Time `json:"day" xorm:"not null pk TIMESTAMP"`
	Open    string    `json:"open" xorm:"DECIMAL(10,2)"`
	Close   string    `json:"close" xorm:"DECIMAL(10,2)"`
	High    string    `json:"high" xorm:"DECIMAL(10,2)"`
	Low     string    `json:"low" xorm:"DECIMAL(10,2)"`
	Volumn  int64     `json:"volumn" xorm:"BIGINT(24)"`
}
