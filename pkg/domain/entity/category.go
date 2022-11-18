package entity

type Categories struct {
	Categories []Category `json:"categories"`
}

type Category struct {
	ID        int    `json:"id"`
	Category  string `json:"category"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
