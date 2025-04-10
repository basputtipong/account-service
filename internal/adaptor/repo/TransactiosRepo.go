package repository

import (
	"account-service/internal/core/port"

	"gorm.io/gorm"
)

type transactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) port.TransactionRepo {
	return &transactionRepo{db}
}

func (r *transactionRepo) GetByUserId(userId string) ([]port.Transaction, error) {
	var repoRes []port.Transaction
	err := r.db.Table(port.TransactionsTbl+" AS t").
		Select("t.transaction_id, t.user_id, t.name, t.image, t.isBank").
		Where("t.user_id = ?", userId).
		Scan(&repoRes).Error
	if err != nil {
		return repoRes, err
	}
	return repoRes, nil
}
