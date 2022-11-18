package entity

type Sales struct {
	Sales []Sale `json:"sales"`
}

type Sale struct {
	ID         int    `json:"id"`
	TotalPrice int    `json:"total_price"`
	Discount   int    `json:"discount"`
	FinalPrice int    `json:"final_price"`
	Cash       int    `json:"cash"`
	Remaining  int    `json:"remaining"`
	Note       string `json:"note"`
	UserID     int    `json:"user_id"`
	Items      string `json:"items"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
