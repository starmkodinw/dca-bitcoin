package models

type Response struct {
	Error  int           `json:"error"`
	Result OrderResponse `json:"result"`
}

type OrderResponse struct {
	Id      string  `json:"id,omitempty"`
	Hash    string  `json:"hash,omitempty"`
	Type    string  `json:"typ,omitempty"`
	Amount  float64 `json:"amt,omitempty"`
	Rate    float64 `json:"rat,omitempty"`
	Fee     float64 `json:"fee,omitempty"`
	Cred    float64 `json:"cre,omitempty"`
	Receive float64 `json:"rec,omitempty"`
	Ts      string  `json:"ts,omitempty"`
}
