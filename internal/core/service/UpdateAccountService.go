package service

import (
	"account-service/internal/core/domain"
	"account-service/internal/core/port"
	"account-service/utils"

	liberror "github.com/basputtipong/library/error"
)

type updateAccountSvc struct {
	accountRepo port.AccountsRepo
}

func NewUpdateAccountSvc(accountRepo port.AccountsRepo) domain.UpdateAccountService {
	return &updateAccountSvc{accountRepo}
}

func (s *updateAccountSvc) Execute(req domain.UpdateAccReq) (domain.UpdateAccRes, error) {
	var emptyRes domain.UpdateAccRes
	var repoReq port.UpdateAccountRepoReq
	var currentMainAccId string

	if err := utils.Validate(req); err != nil {
		return emptyRes, liberror.ErrorBadRequest("Invalid request", err.Error())
	}

	if req.IsMainAccount {
		currentMainAcc, err := s.accountRepo.GetCurrentMainAccountByUserId(req.UserId)
		if err != nil {
			return emptyRes, err
		}

		currentMainAccId = currentMainAcc.AccountId
	}

	repoReq = buildUpdateAccountReq(currentMainAccId, req)
	err := s.accountRepo.UpdateAccountById(repoReq)
	if err != nil {
		return emptyRes, err
	}
	return emptyRes, nil
}

func buildUpdateAccountReq(currentMainAccId string, req domain.UpdateAccReq) port.UpdateAccountRepoReq {
	return port.UpdateAccountRepoReq{
		UserId:           req.UserId,
		AccountId:        req.AccountId,
		CurrentMainAccId: currentMainAccId,
		IsMainAccount:    req.IsMainAccount,
		Color:            req.Color,
	}
}
