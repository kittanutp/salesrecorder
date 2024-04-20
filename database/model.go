package database

import (
	"time"
)

type SaleItem struct {
	ID     uint `gorm:"primaryKey" json:"id"`
	SaleID uint `json:"sale_id"`
	ItemID uint `json:"item_id"`
	Amount int  `json:"amount"`
	Item   Item `json:"items"`
}

type Sale struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	UserID    uint       `json:"user_id"`
	Price     float64    `json:"price"`
	CreatedAt time.Time  `json:"created_at"`
	Items     []SaleItem `gorm:"foreignKey:SaleID" json:"items"`
}

type Item struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Name      string     `gorm:"index:uid_name,priority:1" json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	Cost      float64    `json:"cost"`
	UserID    uint       `gorm:"index:uid_name,priority:2" json:"user_id"`
	SaleItems []SaleItem `gorm:"foreignKey:ItemID" json:"sale_items"`
}

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"uniqueIndex" json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	Sales     []Sale    `gorm:"foreignKey:UserID" json:"sales"`
	Items     []Item    `gorm:"foreignKey:UserID" json:"items"`
}
