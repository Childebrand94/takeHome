package models

type PrefixInfo struct {
	Prefix      int    `json:"prefix"`
	Operator    string `json:"operator"`
	CountryCode int    `json:"country_code"`
	Region      string `json:"region"`
	Country     string `json:"country"`
}

type Resp struct {
	PrefixInfo `json:"prefix_info"`
	Message    string `json:"message"`
}
