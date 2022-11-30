package models

type Order struct {
	ID    int `json:"id" gorm:"primary_key: auto_increment"`
	Qty   int `json:"qty" gorm:"type: int"`
	Total int `json:"total" gorm:"type: int"`

	ProductID int     `json:"product_id"`
	Product   Product `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// ToppingID []int     `json:"topping_id" form:"topping_id" gorm:"-"`
	Topping []Topping `json:"topping" gorm:"many2many:topping_order;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	BuyerID int          `json:"buyer_id"`
	Buyer   UserResponse `json:"buyer" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type OrderResponse struct {
	ID        int   `json:"id"`
	Qty       int   `json:"qty"`
	Price     int   `json:"price"`
	ProductID int   `json:"product_id"`
	ToppingID []int `json:"topping_id"`
	BuyerID   int   `json:"buyer_id"`
}

func (OrderResponse) TableName() string {
	return "orders"
}