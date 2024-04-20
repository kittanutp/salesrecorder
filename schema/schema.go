package schema

type SaleItem struct {
	ID     int `json:"id"`
	SaleID int `json:"sale_id"`
	ItemID int `json:"item_id"`
	Amount int `json:"amount"`
}

type Sale struct {
	ID        int     `json:"id"`
	UserID    int     `json:"user_id"`
	Price     float64 `json:"price"`
	CreatedAt string  `json:"created_at"`
}

type Item struct {
	ID   int     `json:"id"`
	Name string  `json:"name"`
	Cost float64 `json:"cost"`
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AddSaleSchema struct {
	Price float64          `json:"price"`
	Sales []SaleItemSchema `json:"sales"`
}

type SaleItemSchema struct {
	ItemID int `json:"item_id"`
	Amount int `json:"amount"`
}
