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

func TestUpdateAcc_Execute(t *testing.T) {
	t.Run("Should_Pass_When_Update_Account_Color", func(t *testing.T) {
		mockRepo := new(mocks.AccountsRepo)

		mockRepoReq := port.UpdateAccountRepoReq{
			UserId:        "user-test",
			AccountId:     "acc-test",
			IsMainAccount: false,
			Color:         "red",
		}
		mockRepo.On("UpdateAccountById", mockRepoReq).Return(nil)

		svc := service.NewUpdateAccountSvc(mockRepo)
		res, err := svc.Execute(domain.UpdateAccReq{
			UserId:        "user-test",
			AccountId:     "acc-test",
			IsMainAccount: false,
			Color:         "red",
		})

		expected := domain.UpdateAccRes{}
		assert.NoError(t, err)
		assert.Equal(t, expected, res)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Should_Pass_When_Update_New_Main_Account", func(t *testing.T) {
		mockRepo := new(mocks.AccountsRepo)

		mockCurrentMainRes := port.AccountRepoRes{
			AccountId: "old-main",
		}
		mockRepo.On("GetCurrentMainAccountByUserId", "user-test").Return(mockCurrentMainRes, nil)

		mockRepoReq := port.UpdateAccountRepoReq{
			UserId:           "user-test",
			AccountId:        "new-main",
			CurrentMainAccId: "old-main",
			IsMainAccount:    true,
		}
		mockRepo.On("UpdateAccountById", mockRepoReq).Return(nil)

		svc := service.NewUpdateAccountSvc(mockRepo)
		res, err := svc.Execute(domain.UpdateAccReq{
			UserId:        "user-test",
			AccountId:     "new-main",
			IsMainAccount: true,
		})

		expected := domain.UpdateAccRes{}
		assert.NoError(t, err)
		assert.Equal(t, expected, res)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Should_Error_When_Req_Not_Valid", func(t *testing.T) {
		mockRepo := new(mocks.AccountsRepo)
		svc := service.NewUpdateAccountSvc(mockRepo)
		res, err := svc.Execute(domain.UpdateAccReq{
			UserId:    "",
			AccountId: "",
		})

		expected := domain.UpdateAccRes{}
		assert.Error(t, err)
		assert.Equal(t, expected, res)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Should_Error_When_Get_Old_Main_Account", func(t *testing.T) {
		mockRepo := new(mocks.AccountsRepo)

		mockCurrentMainRes := port.AccountRepoRes{}
		mockRepo.On("GetCurrentMainAccountByUserId", "user-test").Return(mockCurrentMainRes, liberror.ErrorInternalServerError("", ""))

		svc := service.NewUpdateAccountSvc(mockRepo)
		res, err := svc.Execute(domain.UpdateAccReq{
			UserId:        "user-test",
			AccountId:     "new-main",
			IsMainAccount: true,
		})

		expected := domain.UpdateAccRes{}
		assert.Error(t, err)
		assert.Equal(t, expected, res)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Should_Error_When_Update_New_Main_Acc", func(t *testing.T) {
		mockRepo := new(mocks.AccountsRepo)

		mockCurrentMainRes := port.AccountRepoRes{
			AccountId: "old-main",
		}
		mockRepo.On("GetCurrentMainAccountByUserId", "user-test").Return(mockCurrentMainRes, nil)

		mockRepoReq := port.UpdateAccountRepoReq{
			UserId:           "user-test",
			AccountId:        "new-main",
			CurrentMainAccId: "old-main",
			IsMainAccount:    true,
		}
		mockRepo.On("UpdateAccountById", mockRepoReq).Return(liberror.ErrorInternalServerError("", ""))

		svc := service.NewUpdateAccountSvc(mockRepo)
		res, err := svc.Execute(domain.UpdateAccReq{
			UserId:        "user-test",
			AccountId:     "new-main",
			IsMainAccount: true,
		})

		expected := domain.UpdateAccRes{}
		assert.Error(t, err)
		assert.Equal(t, expected, res)
		mockRepo.AssertExpectations(t)
	})
}
