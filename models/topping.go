package models

type Topping struct {
	ID     int          `json:"id"`
	Title  string       `json:"title" form:"title" gorm:"type: varchar(255)"`
	Price  int          `json:"price" form:"price" gorm:"type: varchar(255)"`
	Image  string       `json:"image" form:"image" gorm:"type: varchar(255)"`
	Qty    int          `json:"qty" form:"qty"`
	UserID int          `json:"user_id"`
	User   UserResponse `json:"user"`
}

type ToppingResponse struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Price  int    `json:"price"`
	Image  string `json:"image"`
	Qty    int    `json:"qty"`
	UserID int    `json:"-"`
}

func (ToppingResponse) TableName() string {
	return "topping"
}
