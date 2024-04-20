package database

import (
	"time"
)

type SaleItem struct {
	ID     uint `gorm:"primaryKey" json:"id"`
	SaleID uint `json:"sale_id"`
	ItemID uint `json:"item_id"`
	Amount uint `json:"amount"`
}

type Sale struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	UserID    int        `json:"user_id"`
	Price     float64    `json:"price"`
	CreatedAt time.Time  `json:"created_at"`
	Item      []SaleItem `gorm:"foreignKey:SaleID" json:"items"`
}

type Item struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Cost      float64   `json:"cost"`
	UserID    uint      `json:"user_id"`
}

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	Sales     []Sale    `gorm:"foreignKey:UserID" json:"sales"`
	Items     []Item    `gorm:"foreignKey:UserID" json:"items"`
}
