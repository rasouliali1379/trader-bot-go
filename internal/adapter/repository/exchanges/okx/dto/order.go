package dto

type PlaceOrderRequest struct {
	InstID  string `json:"instId"`
	TdMode  string `json:"tdMode"`
	ClOrdID string `json:"clOrdId"`
	Side    string `json:"side"`
	OrdType string `json:"ordType"`
	Px      string `json:"px"`
	Sz      string `json:"sz"`
}
