package models

type Product struct {
	ID       uint    `json:"id" gorm:"primaryKey`
	Name     string  `json:"name`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}
