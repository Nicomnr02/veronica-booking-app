package entity

type Sizes struct {
	Sizes []Size `json:"sizes"`
}

type Size struct {
	ID        int    `json:"id"`
	Size      string `json:"size"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
