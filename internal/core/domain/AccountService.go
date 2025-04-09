package domain

import "account-service/internal/core/port"

type AccountService interface {
	Execute(req AccountReq) (AccountRes, error)
}

type AccountReq struct {
	UserId string `json:"userId" validate:"required"`
}

type AccountRes struct {
	Accounts []Account `json:"accounts"`
}

type Account struct {
	AccountId     string  `json:"accountId"`
	Type          string  `json:"type"`
	Currency      string  `json:"currency"`
	AccountNumber string  `json:"accountNumber"`
	Issuer        string  `json:"issuer"`
	Amount        float64 `json:"amount"`
	Color         string  `json:"color"`
	IsMainAccount bool    `json:"isMainAccount"`
	Progress      int64   `json:"progress"`
	FlagId        int64   `json:"flagId"`
	FlagType      string  `json:"flagType"`
	FlagValue     string  `json:"flagValue"`
}

func (res *AccountRes) BuildAccountResponse(repoRes []port.AccountRepoRes) {
	var accounts []Account
	for _, ele := range repoRes {
		accounts = append(accounts, Account{
			AccountId:     ele.AccountId,
			Type:          ele.Type,
			Currency:      ele.Currency,
			AccountNumber: ele.AccountNumber,
			Issuer:        ele.Issuer,
			Amount:        ele.Amount,
			Color:         ele.Color,
			IsMainAccount: ele.IsMainAccount,
			Progress:      ele.Progress,
			FlagId:        ele.FlagId,
			FlagType:      ele.FlagType,
			FlagValue:     ele.FlagValue,
		})
	}
	res.Accounts = accounts
}
