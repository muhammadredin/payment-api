package entity

type Transaction struct {
	Id           string  `json:"id"`
	FromWalletId string  `json:"from_wallet_id"`
	ToWalletId   string  `json:"to_wallet_id"`
	CreatedAt    string  `json:"created_at"`
	Amount       float64 `json:"amount"`
	Message      string  `json:"message"`
}
