package models

type TCsrcMarketSzSzseIndustry struct {
	CsrcId                  string  `json:"csrc_id" xorm:"unique(csrc_id) VARCHAR(128)"`
	StockCount              int     `json:"stock_count" xorm:"INT(10)"`
	TotalShareCapital       float32 `json:"total_share_capital" xorm:"FLOAT(16,2)"`
	TotalMarketValue        float32 `json:"total_market_value" xorm:"FLOAT(16,2)"`
	TradeVolume             float32 `json:"trade_volume" xorm:"FLOAT(16,2)"`
	TradeAmount             float32 `json:"trade_amount" xorm:"FLOAT(16,2)"`
	CirculatingShareCapital float32 `json:"circulating_share_capital" xorm:"FLOAT(16,2)"`
	CirculatingMarketValue  float32 `json:"circulating_market_value" xorm:"FLOAT(16,2)"`
	DataTime                int     `json:"data_time" xorm:"unique(csrc_id) INT(10)"`
	CreateTime              int     `json:"create_time" xorm:"INT(10)"`
}
