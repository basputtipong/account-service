package domain

import (
	"account-service/internal/core/port"
	"math"
)

type AccountService interface {
	Execute(req AccountReq) (AccountRes, error)
}

type AccountReq struct {
	UserId string `json:"userId" validate:"required"`
}

type AccountRes struct {
	Accounts     []Account `json:"accounts"`
	TotalBalance float64   `json:"totalBalance"`
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
	Flags         []Flag  `json:"flags"`
}

type Flag struct {
	FlagId    int64  `json:"flagId"`
	FlagType  string `json:"flagType"`
	FlagValue string `json:"flagValue"`
}

func (res *AccountRes) BuildAccountResponse(repoRes []port.AccountRepoRes, flagRes []port.Flag) {
	flagMap := make(map[string][]Flag)
	for _, f := range flagRes {
		flagMap[f.AccountId] = append(flagMap[f.AccountId], Flag{
			FlagId:    f.FlagId,
			FlagType:  f.FlagType,
			FlagValue: f.FlagValue,
		})
	}

	var totalBalance float64
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
			Flags:         flagMap[ele.AccountId],
		})
		totalBalance = totalBalance + ele.Amount
	}
	res.Accounts = accounts
	res.TotalBalance = math.Round(totalBalance*100) / 100
}
