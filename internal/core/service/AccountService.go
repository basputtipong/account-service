package service

import (
	"account-service/internal/core/domain"
	"account-service/internal/core/port"
)

type accountSvc struct {
	accountRepo port.AccountsRepo
}

func NewAccountSvc(accountRepo port.AccountsRepo) domain.AccountService {
	return &accountSvc{accountRepo}
}

func (s *accountSvc) Execute(req domain.AccountReq) (domain.AccountRes, error) {
	var res domain.AccountRes
	repoRes, err := s.accountRepo.GetByUserId(req.UserId)
	if err != nil {
		return res, err
	}

	accountIds := buildGetFlagReq(repoRes)
	flagRes, err := s.accountRepo.GetFlagByAccountId(accountIds)
	if err != nil {
		return res, err
	}

	res.BuildAccountResponse(repoRes, flagRes)
	return res, nil
}

func buildGetFlagReq(repoRes []port.AccountRepoRes) []string {
	var accountIds []string
	for _, ele := range repoRes {
		accountIds = append(accountIds, ele.AccountId)
	}

	return accountIds
}
