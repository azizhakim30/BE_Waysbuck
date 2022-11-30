package models

import "time"

type Transaction struct {
	ID       int     `json:"id" gorm:"primary_key:auto_increment"`
	Name     string  `json:"name" form:"name" gorm:"type: varchar(255)"`
	Email    string  `json:"email" form:"email" gorm:"type: varchar(255)"`
	Phone    string  `json:"phone" form:"phone" gorm:"type: varchar(255)"`
	Address  string  `json:"address" form:"address" gorm:"type : text"`
	Total 	 int		 `json:"total" form:"total" gorm:"type : int"`
	Status   string  `json:"status" gorm:"type: varchar(255)"`
	OrderID  int		 `json:"-" `
	Order		 []Order `json:"order" gorm:"many2many:transaction_order;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	BuyerID  int		 `json:"buyer_id"`
	Buyer    UserResponse `json:"buyer" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreateAt time.Time `json:"-"`
}