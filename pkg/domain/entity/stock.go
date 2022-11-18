package entity

type Stocks struct {
	Stocks []Stock `json:"stocks"`
}

type Stock struct {
	ID        int    `json:"id"`
	ProductID int    `json:"product_id"`
	Detail    string `json:"detail"`
	Quantity  int    `json:"quantity"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
