package repositories

import (
	"waysbuck/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	FindTransaction() ([]models.Transaction, error)
	DeleteTransaction(transaction models.Transaction) (models.Transaction, error)
	UpdateTransaction(transaction models.Transaction) (models.Transaction, error)
	GetOrderByID(userID int) (models.Transaction, error)
	FindTransactionID() (models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("Buyer").Create(&transaction).Error
	return transaction, err
}

func (r *repository) GetTransaction(userID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Order").Preload("User").Where("user_id = ?", userID).First(&transaction).Error
	return transaction, err
}

func (r *repository) FindTransaction() ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Preload("User").Find(&transaction).Error
	return transaction, err
}

func (r *repository) DeleteTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Delete(&transaction).Error
	return transaction, err
}

func (r *repository) GetOrderByID(userID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").Where("user_id = ?", userID).Find(&transaction).Error
	return transaction, err
}

func (r *repository) UpdateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Save(&transaction).Error
	return transaction, err
}

func (r *repository) FindTransactionID() (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").Preload("Carts").Preload("Carts.Product").Preload("Carts.Topping").Find(&transaction, "status = ?", "waiting").Error

	return transaction, err
}