package sa_xorm_model

type TCsrcMarket struct {
	Plate           int     `json:"plate" xorm:"unique(plate) INT(10)"`
	CsrcId          string  `json:"csrc_id" xorm:"unique(plate) VARCHAR(128)"`
	TradeStockCount int     `json:"trade_stock_count" xorm:"INT(10)"`
	MarketValue     int     `json:"market_value" xorm:"INT(164)"`
	AvlPe           float32 `json:"avl_pe" xorm:"FLOAT(16,2)"`
	AvlPrice        float32 `json:"avl_price" xorm:"FLOAT(16,2)"`
	DataTime        int     `json:"data_time" xorm:"unique(plate) INT(10)"`
	CreateTime      int     `json:"create_time" xorm:"INT(10)"`
}
