package models

import "time"

type User struct{
	ID 				int 			`json:"id" gorm:"primary_key: auto_increment"`
	Fullname 	string 		`json:"fullname" gorm:"type: varchar(255)"`
	Password 	string 		`json:"-" gorm:"type: varchar(255)"`
	Email			string 		`json:"email" gorm:"type: varchar(255)"`
	Image			string		`json:"image" gorm:"type: varchar(255)"`
	Role 			string 		`json:"role"`
	CreateAt 	time.Time `json:"-"`
	UpdateAt 	time.Time `json:"-"`
}

type UserResponse struct {
	ID 				int 	 `json:"id"`
	Fullname 	string `json:"fullname"`
	Email			string `json:"email"`
	Image			string `json:"-"`
	Role			string `json:"-"`
}

func (UserResponse) TableName() string {
	return "users"
}