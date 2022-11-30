package transactiondto

type TransactionRequest struct {
	ID      int    `json:"id"`
	Name    string `json:"name" form:"name" gorm:"type: varchar(255)"`
	Email   string `json:"email" form:"email" gorm:"type: varchar(255)"`
	Phone   string `json:"phone" form:"phone" gorm:"type: varchar(255)"`
	PosCode string `json:"pos_code" form:"pos_code" gorm:"type: varchar(255)"`
	Address string `json:"address" form:"address" gorm:"type : varchar(255)"`
	Total   int    `json:"total" gorm:"type : int"`
}

type TransactionResponse struct {
	ID      int      `json:"id"`
	Name    string   `json:"name" form:"name" gorm:"type: varchar(255)"`
	Email   string   `json:"email" form:"email" gorm:"type: varchar(255)"`
	Phone   string   `json:"phone" form:"phone" gorm:"type: varchar(255)"`
	PosCode string   `json:"pos_code" form:"pos_code" gorm:"type: varchar(255)"`
	Address string   `json:"address" form:"address" gorm:"type : varchar(255)"`
	Order   []string `json:"order"`
	Total   int      `json:"total" gorm:"type : int"`
}

type StatusTransaction struct {
	Status string `json:"status" gorm:"type: varchar(255)"`
}