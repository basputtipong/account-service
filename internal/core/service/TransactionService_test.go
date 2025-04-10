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

func TestTransactionService_Execute(t *testing.T) {
	t.Run("Should_Pass_Return_Transactions", func(t *testing.T) {
		mockRepo := new(mocks.TransactionRepo)

		mockTxnRepoRes := []port.Transaction{
			{
				TransactionId: "123",
				UserId:        "user-test",
				IsBank:        true,
			},
			{
				TransactionId: "124",
				UserId:        "user-test",
				IsBank:        true,
			},
		}
		mockRepo.On("GetByUserId", "user-test").Return(mockTxnRepoRes, nil)

		svc := service.NewTransactionSvc(mockRepo)

		res, err := svc.Execute(domain.TransactionReq{UserId: "user-test"})

		expected := domain.TransactionRes{
			Transactions: []domain.Transaction{
				{
					TransactionId: "123",
					IsBank:        true,
				},
				{
					TransactionId: "124",
					IsBank:        true,
				},
			},
		}

		assert.NoError(t, err)
		assert.Equal(t, expected, res)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Should_Error_When_No_UserID", func(t *testing.T) {
		mockRepo := new(mocks.TransactionRepo)
		svc := service.NewTransactionSvc(mockRepo)

		res, err := svc.Execute(domain.TransactionReq{UserId: ""})

		expected := domain.TransactionRes{}

		assert.Error(t, err)
		assert.Equal(t, expected, res)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Should_Error_When_Fail_Get_Txn", func(t *testing.T) {
		mockRepo := new(mocks.TransactionRepo)

		mockTxnRepoRes := []port.Transaction{}
		mockRepo.On("GetByUserId", "user-test").Return(mockTxnRepoRes, liberror.ErrorInternalServerError("", ""))

		svc := service.NewTransactionSvc(mockRepo)

		res, err := svc.Execute(domain.TransactionReq{UserId: "user-test"})

		expected := domain.TransactionRes{}

		assert.Error(t, err)
		assert.Equal(t, expected, res)
		mockRepo.AssertExpectations(t)
	})
}
