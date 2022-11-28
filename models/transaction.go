package models

import "time"

type Transaction struct {
	ID       int     `json:"id"`
	Name     string  `json:"name" form:"name" gorm:"type: varchar(255)"`
	Email    string  `json:"email" form:"email" gorm:"type: varchar(255)"`
	Phone    string  `json:"phone" form:"phone" gorm:"type: varchar(255)"`
	PosCode  string  `json:"pos_code" form:"pos_code" gorm:"type: varchar(255)"`
	Address  string  `json:"address" form:"address" gorm:"type : varchar(255)"`
	Total 	 int		 `json:"total" form:"total" gorm:"type : int"`
	OrderID  []int	 `json:"-" gorm:"-"`
	Order    []Order `json:"order" gorm:"many2many:transaction_order;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status 	 string 	 `json:"status" gorm:"type : varchar(255)"`
	CreateAt time.Time `json:"-"`	
}