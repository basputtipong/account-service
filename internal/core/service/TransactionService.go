package service

import (
	"account-service/internal/core/domain"
	"account-service/internal/core/port"
)

type transactionSvc struct {
	transactionRepo port.TransactionRepo
}

func NewTransactionSvc(transactionRepo port.TransactionRepo) domain.TransactionService {
	return &transactionSvc{transactionRepo}
}

func (s *transactionSvc) Execute(req domain.TransactionReq) (domain.TransactionRes, error) {
	var res domain.TransactionRes
	repoRes, err := s.transactionRepo.GetByUserId(req.UserId)
	if err != nil {
		return res, err
	}

	res.BuildTransactionRes(repoRes)
	return res, nil
}
