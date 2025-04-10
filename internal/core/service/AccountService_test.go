package service_test

import (
	"account-service/internal/core/domain"
	"account-service/internal/core/port"
	"account-service/internal/core/port/mocks"
	"account-service/internal/core/service"
	"testing"

	liberror "github.com/basputtipong/library/error"
	"github.com/stretchr/testify/assert"
)

func TestAccountService_Execute(t *testing.T) {
	t.Run("Should_Pass_Return_Accounts", func(t *testing.T) {
		mockRepo := new(mocks.AccountsRepo)

		mockAccRepoRes := []port.AccountRepoRes{
			{
				AccountId:     "acc1",
				Type:          "saving",
				Currency:      "THB",
				AccountNumber: "1234567890",
				Issuer:        "system",
				Amount:        123.45,
				Color:         "",
				IsMainAccount: true,
				Progress:      50,
			},
			{
				AccountId:     "acc2",
				Type:          "saving",
				Currency:      "THB",
				AccountNumber: "1234567891",
				Issuer:        "system",
				Amount:        123.45,
				Color:         "",
				IsMainAccount: false,
				Progress:      50,
			},
		}
		mockRepo.On("GetByUserId", "user-test").Return(mockAccRepoRes, nil)

		mockFlagRes := []port.Flag{
			{
				FlagId:    123,
				AccountId: "acc1",
			},
			{
				FlagId:    456,
				AccountId: "acc2",
			},
		}
		mockRepo.On("GetFlagByAccountId", []string{"acc1", "acc2"}).Return(mockFlagRes, nil)

		svc := service.NewAccountSvc(mockRepo)
		res, err := svc.Execute(domain.AccountReq{
			UserId: "user-test",
		})

		expected := domain.AccountRes{
			Accounts: []domain.Account{
				{
					AccountId:     "acc1",
					Type:          "saving",
					Currency:      "THB",
					AccountNumber: "1234567890",
					Issuer:        "system",
					Amount:        123.45,
					Color:         "",
					IsMainAccount: true,
					Progress:      50,
					Flags: []domain.Flag{
						{
							FlagId: 123,
						},
					},
				},
				{
					AccountId:     "acc2",
					Type:          "saving",
					Currency:      "THB",
					AccountNumber: "1234567891",
					Issuer:        "system",
					Amount:        123.45,
					Color:         "",
					IsMainAccount: false,
					Progress:      50,
					Flags: []domain.Flag{
						{
							FlagId: 456,
						},
					},
				},
			},
			TotalBalance: 246.90,
		}

		assert.NoError(t, err)
		assert.Equal(t, expected, res)
		mockRepo.AssertExpectations(t)
	})
	t.Run("Should_Error_When_No_UserId", func(t *testing.T) {
		mockRepo := new(mocks.AccountsRepo)
		svc := service.NewAccountSvc(mockRepo)

		res, err := svc.Execute(domain.AccountReq{})
		expected := domain.AccountRes{}

		assert.Error(t, err)
		assert.Equal(t, expected, res)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Should_Error_When_Get_Account_Fail", func(t *testing.T) {
		mockRepo := new(mocks.AccountsRepo)

		mockAccRepoRes := []port.AccountRepoRes{}
		mockRepo.On("GetByUserId", "user-test").Return(mockAccRepoRes, liberror.ErrorInternalServerError("", ""))

		svc := service.NewAccountSvc(mockRepo)
		res, err := svc.Execute(domain.AccountReq{
			UserId: "user-test",
		})

		expected := domain.AccountRes{}

		assert.Error(t, err)
		assert.Equal(t, expected, res)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Should_Error_When_Get_Flag_Fail", func(t *testing.T) {
		mockRepo := new(mocks.AccountsRepo)

		mockAccRepoRes := []port.AccountRepoRes{
			{
				AccountId:     "acc1",
				Type:          "saving",
				Currency:      "THB",
				AccountNumber: "1234567890",
				Issuer:        "system",
				Amount:        123.45,
				Color:         "",
				IsMainAccount: true,
				Progress:      50,
			},
			{
				AccountId:     "acc2",
				Type:          "saving",
				Currency:      "THB",
				AccountNumber: "1234567891",
				Issuer:        "system",
				Amount:        123.45,
				Color:         "",
				IsMainAccount: false,
				Progress:      50,
			},
		}
		mockRepo.On("GetByUserId", "user-test").Return(mockAccRepoRes, liberror.ErrorInternalServerError("", ""))

		mockFlagRes := []port.Flag{}
		mockRepo.On("GetFlagByAccountId", []string{"acc1", "acc2"}).Return(mockFlagRes, liberror.ErrorInternalServerError("", ""))

		svc := service.NewAccountSvc(mockRepo)
		res, err := svc.Execute(domain.AccountReq{
			UserId: "user-test",
		})

		expected := domain.AccountRes{}

		assert.Error(t, err)
		assert.Equal(t, expected, res)
	})
}
