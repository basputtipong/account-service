package repository

import (
	"account-service/internal/core/port"
	"fmt"

	liberror "github.com/basputtipong/library/error"
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
func (r *accountsRepo) GetCurrentMainAccountByUserId(userId string) (port.AccountRepoRes, error) {
	var repoRes port.AccountRepoRes
	err := r.db.Table(port.AccountsTbl+" AS a").
		Select(`a.account_id, a.type, a.currency, a.account_number, a.issuer,
            ab.amount,
            ad.color, ad.is_main_account, ad.progress`).
		Joins(fmt.Sprintf(`LEFT JOIN %s ab ON a.account_id = ab.account_id`, port.AccountBalancesTbl)).
		Joins(fmt.Sprintf(`LEFT JOIN %s ad ON ab.account_id = ad.account_id`, port.AccountDetailsTbl)).
		Where("a.user_id = ? AND ad.is_main_account = ?", userId, true).
		Scan(&repoRes).Error

	if err != nil {
		return repoRes, err
	}

	return repoRes, nil
}

func (r *accountsRepo) UpdateAccountById(req port.UpdateAccountRepoReq) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return liberror.ErrorInternalServerError("Failed to begin transaction", tx.Error.Error())
	}

	var err error
	if req.IsMainAccount {
		err = updateMainAccount(tx, req.AccountId, req.CurrentMainAccId)
	} else {
		err = updateColorOnly(tx, req.AccountId, req.Color)
	}

	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return liberror.ErrorInternalServerError("Failed to commit transaction", err.Error())
	}

	return nil
}

func updateMainAccount(tx *gorm.DB, accountId, currentMainAcc string) error {
	query := fmt.Sprintf(`
		UPDATE %s
		SET is_main_account = CASE account_id
			WHEN ? THEN true
			WHEN ? THEN false
		END
		WHERE account_id IN (?, ?)`, port.AccountDetailsTbl)

	err := tx.Exec(query, accountId, currentMainAcc, accountId, currentMainAcc).Error
	if err != nil {
		return liberror.ErrorInternalServerError("Failed to update main account", err.Error())
	}
	return nil
}

func updateColorOnly(tx *gorm.DB, accountID, color string) error {
	res := tx.Table(port.AccountDetailsTbl).
		Where("account_id = ?", accountID).
		Update("color", color)

	if res.Error != nil {
		return liberror.ErrorInternalServerError("Failed to update color account", res.Error.Error())
	}
	if res.RowsAffected == 0 {
		return liberror.ErrorInternalServerError("No rows were updated", "Account ID might be invalid")
	}
	return nil
}
