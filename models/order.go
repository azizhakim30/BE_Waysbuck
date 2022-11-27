package models

type Order struct {
	ID    int `json:"id" gorm:"primary_key: auto_increment"`
	Qty   int `json:"qty" gorm:"type: int"`
	Price int `json:"price" gorm:"type: int"`

	ProductID int             `json:"product_id"`
	Product   ProductResponse `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// ToppingID []int     `json:"topping_id" form:"topping_id" gorm:"-"`
	Topping []Topping `json:"topping" gorm:"many2many:topping_order;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	BuyyerID int          `json:"buyyer_id"`
	Buyyer   UserResponse `json:"buyyer" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type OrderResponse struct {
	ID        int `json:"id"`
	Qty       int `json:"qty"`
	Price     int `json:"price"`
	ProductID int `json:"product_id"`
	TopingID  int `json:"toping_id"`
	BuyyerID  int `json:"buyyer_id"`
}

func (OrderResponse) TableName() string {
	return "orders"
}