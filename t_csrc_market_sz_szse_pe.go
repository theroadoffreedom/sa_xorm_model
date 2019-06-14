package models

type TCsrcMarketSzSzsePe struct {
	CsrcId           string  `json:"csrc_id" xorm:"unique(csrc_id) VARCHAR(128)"`
	StockCountSzMain int     `json:"stock_count_sz_main" xorm:"INT(10)"`
	StockAvgPeSzMain float32 `json:"stock_avg_pe_sz_main" xorm:"FLOAT(16,2)"`
	StockCountSzSme  int     `json:"stock_count_sz_sme" xorm:"INT(10)"`
	StockAvgPeSzSme  float32 `json:"stock_avg_pe_sz_sme" xorm:"FLOAT(16,2)"`
	StockCountSzGeb  int     `json:"stock_count_sz_geb" xorm:"INT(10)"`
	StockAvgPeSzGeb  float32 `json:"stock_avg_pe_sz_geb" xorm:"FLOAT(16,2)"`
	DataTime         int     `json:"data_time" xorm:"unique(csrc_id) INT(10)"`
	CreateTime       int     `json:"create_time" xorm:"INT(10)"`
}
