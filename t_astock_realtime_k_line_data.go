package sa_xorm_model

import (
	"time"
)

type TAstockRealtimeKLineData struct {
	StockId      string    `json:"stock_id" xorm:"not null pk VARCHAR(128)"`
	Moment       time.Time `json:"moment" xorm:"not null pk TIMESTAMP"`
	Open         string    `json:"open" xorm:"DECIMAL(10,2)"`
	LastdayClose string    `json:"lastday_close" xorm:"DECIMAL(10,2)"`
	Cur          string    `json:"cur" xorm:"DECIMAL(10,2)"`
	High         string    `json:"high" xorm:"DECIMAL(10,2)"`
	Low          string    `json:"low" xorm:"DECIMAL(10,2)"`
	BuyFirst     string    `json:"buy_first" xorm:"DECIMAL(10,2)"`
	SellFirst    string    `json:"sell_first" xorm:"DECIMAL(10,2)"`
	Volumn       int64     `json:"volumn" xorm:"BIGINT(24)"`
	Amount       string    `json:"amount" xorm:"DECIMAL(10,2)"`
	Turnoverrate string    `json:"turnoverrate" xorm:"DECIMAL(10,2)"`
	Percent      string    `json:"percent" xorm:"DECIMAL(10,2)"`
	BuyOne       string    `json:"buy_one" xorm:"DECIMAL(10,2)"`
	BuyTwo       string    `json:"buy_two" xorm:"DECIMAL(10,2)"`
	BuyThree     string    `json:"buy_three" xorm:"DECIMAL(10,2)"`
	BuyFour      string    `json:"buy_four" xorm:"DECIMAL(10,2)"`
	BuyFive      string    `json:"buy_five" xorm:"DECIMAL(10,2)"`
	SellOne      string    `json:"sell_one" xorm:"DECIMAL(10,2)"`
	SellTwo      string    `json:"sell_two" xorm:"DECIMAL(10,2)"`
	SellThree    string    `json:"sell_three" xorm:"DECIMAL(10,2)"`
	SellFour     string    `json:"sell_four" xorm:"DECIMAL(10,2)"`
	SellFive     string    `json:"sell_five" xorm:"DECIMAL(10,2)"`
}
