package model

type Exchange struct {
	Usd  Usd  `json:"USD"`
	Eur  Eur  `json:"EUR"`
	Gbp  Gbp  `json:"GBP"`
	Btc  Btc  `json:"BTC"`
	Eth  Eth  `json:"ETH"`
	Gold Gold `json:"GA"`
}
type Usd struct {
	Satis   string `json:"satis"`
	Alis    string `json:"alis"`
	Degisim string `json:"degisim"`
}
type Eur struct {
	Satis   string `json:"satis"`
	Alis    string `json:"alis"`
	Degisim string `json:"degisim"`
}
type Gbp struct {
	Satis   string `json:"satis"`
	Alis    string `json:"alis"`
	Degisim string `json:"degisim"`
}
type Btc struct {
	Satis   string `json:"satis"`
	Alis    string `json:"alis"`
	Degisim string `json:"degisim"`
}
type Eth struct {
	Satis   string `json:"satis"`
	Alis    string `json:"alis"`
	Degisim string `json:"degisim"`
}
type Gold struct {
	Satis   string `json:"satis"`
	Alis    string `json:"alis"`
	Degisim string `json:"degisim"`
}

type Currency struct {
	Satis   string `json:"satis"`
	Alis    string `json:"alis"`
	Degisim string `json:"degisim"`
}
