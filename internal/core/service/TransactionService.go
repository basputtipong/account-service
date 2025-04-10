package service

import (
	"account-service/internal/core/domain"
	"account-service/internal/core/port"
	"account-service/utils"

	liberror "github.com/basputtipong/library/error"
)

type transactionSvc struct {
	transactionRepo port.TransactionRepo
}

func NewTransactionSvc(transactionRepo port.TransactionRepo) domain.TransactionService {
	return &transactionSvc{transactionRepo}
}

func (s *transactionSvc) Execute(req domain.TransactionReq) (domain.TransactionRes, error) {
	var res domain.TransactionRes

	if err := utils.Validate(req); err != nil {
		return res, liberror.ErrorBadRequest("Invalid request", err.Error())
	}

	repoRes, err := s.transactionRepo.GetByUserId(req.UserId)
	if err != nil {
		return res, err
	}

	res.BuildTransactionRes(repoRes)
	return res, nil
}
