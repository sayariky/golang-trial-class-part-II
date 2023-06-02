package entity

import "time"

type Order struct {
	ID           uint
	BuyerEmail   string
	BuyerAddress string
	ProductID    uint
	OrderDate    time.Time
	Product      Product
}
