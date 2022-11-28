package orderdto

type CreateOrder struct {
	ID        int `json:"id"`
	Price     int `json:"price" form:"price"`
	ProductID int `json:"product_id" form:"product_id"`
	ToppingID []int `json:"topping_id" form:"topping_id"`
	BuyyerID  int   `json:"buyyer_id" form:"buyyer_id"`
	Qty       int   `json:"qty" form:"qty"`
}

type UpdateOrder struct {
	Qty int `json:"qty" form:"qty"`
}
