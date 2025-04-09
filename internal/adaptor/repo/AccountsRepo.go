package repository

import (
	"account-service/internal/core/port"
	"fmt"

	"gorm.io/gorm"
)

type accountsRepo struct {
	db *gorm.DB
}

func NewAccountsRepo(db *gorm.DB) port.AccountsRepo {
	return &accountsRepo{db: db}
}

func (r *accountsRepo) GetByUserId(userId string) ([]port.AccountRepoRes, error) {
	var repoRes []port.AccountRepoRes
	err := r.db.Table(port.AccountsTbl+" AS a").
		Select(`a.account_id, a.type, a.currency, a.account_number, a.issuer,
            ab.amount,
            ad.color, ad.is_main_account, ad.progress`).
		Joins(fmt.Sprintf(`LEFT JOIN %s ab ON a.account_id = ab.account_id`, port.AccountBalancesTbl)).
		Joins(fmt.Sprintf(`LEFT JOIN %s ad ON ab.account_id = ad.account_id`, port.AccountDetailsTbl)).
		Where("a.user_id = ?", userId).
		Scan(&repoRes).Error

	if err != nil {
		return nil, err
	}

	return repoRes, nil
}

func (r *accountsRepo) GetFlagByAccountId(accountIds []string) ([]port.Flag, error) {
	var flagRes []port.Flag
	err := r.db.Where("account_id in ?", accountIds).Find(&flagRes).Error
	if err != nil {
		return nil, err
	}

	return flagRes, nil
}
