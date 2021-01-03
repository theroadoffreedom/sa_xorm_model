package sa_xorm_model

type TStock struct {
	Id                       string  `json:"id" xorm:"not null pk VARCHAR(128)"`
	Cn                       string  `json:"cn" xorm:"VARCHAR(128)"`
	CsrcId                   string  `json:"csrc_id" xorm:"VARCHAR(128)"`
	FullName                 string  `json:"full_name" xorm:"VARCHAR(512)"`
	RegisterAddress          string  `json:"register_address" xorm:"VARCHAR(1024)"`
	PublicTime               int     `json:"public_time" xorm:"INT(10)"`
	Region                   string  `json:"region" xorm:"VARCHAR(128)"`
	Province                 string  `json:"province" xorm:"VARCHAR(128)"`
	City                     string  `json:"city" xorm:"VARCHAR(128)"`
	Plate                    int     `json:"plate" xorm:"INT(10)"`
	State                    int     `json:"state" xorm:"INT(10)"`
	CreateTime               int     `json:"create_time" xorm:"INT(10)"`
	SearchTime               int     `json:"search_time" xorm:"INT(10)"`
	CbCheckTime              int     `json:"cb_check_time" xorm:"INT(10)"`
	CurrentPrice             float64 `json:"current_price" xorm:"DECIMAL(10,2)"`
	High52wPrice             float64 `json:"high52w_price" xorm:"DECIMAL(10,2)"`
	Low52wPrice              float64 `json:"low52w_price" xorm:"DECIMAL(10,2)"`
	InfoCheckTime            int     `json:"info_check_time" xorm:"INT(10)"`
	TotalAShareCapital       int64   `json:"total_a_share_capital" xorm:"BIGINT(20)"`
	FreeMarketCapital        int64   `json:"free_market_capital" xorm:"BIGINT(20)"`
	MarketCapital            int64   `json:"market_capital" xorm:"BIGINT(20)"`
	FreeAShareCapital        int64   `json:"free_a_share_capital" xorm:"BIGINT(20)"`
	CirculatingAShareCapital int64   `json:"circulating_a_share_capital" xorm:"BIGINT(20)"`
}
