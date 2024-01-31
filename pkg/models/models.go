package models

type Record struct {
	Artist       string  `json:"id"`
	Record_Name  string  `json:"record_name"`
	Record_Price float32 `json:"record_price"`
}
