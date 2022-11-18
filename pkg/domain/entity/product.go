package entity

type Products struct {
	Products []Product `json:"products"`
}

type Product struct {
	ID            int    `json:"id"`
	CategoryID    int    `json:"category_id"`
	SizeID        int    `json:"size_id"`
	Barcode       string `json:"Barcode"`
	Name          string `json:"name"`
	Price         int    `json:"price"`
	Stock         int    `json:"stock"`
	DetailProduct string `json:"detail_product"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
