package domain

import "account-service/internal/core/port"

type TransactionService interface {
	Execute(req TransactionReq) (TransactionRes, error)
}

type TransactionReq struct {
	UserId string `json:"userId" validate:"required"`
}

type TransactionRes struct {
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	TransactionId string `json:"transactionId"`
	Name          string `json:"name"`
	Image         string `json:"image"`
	IsBank        bool   `json:"isBank"`
}

func (res *TransactionRes) BuildTransactionRes(repoRes []port.Transaction) {
	var transactions []Transaction
	for _, ele := range repoRes {
		transactions = append(transactions, Transaction{
			TransactionId: ele.TransactionId,
			Name:          ele.Name,
			Image:         ele.Image,
			IsBank:        ele.IsBank,
		})
	}
	res.Transactions = transactions
}
