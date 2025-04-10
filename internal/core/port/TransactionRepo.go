package port

type TransactionRepo interface {
	GetByUserId(userId string) ([]Transaction, error)
}

const (
	TransactionsTbl = "transactions"
)

type Transaction struct {
	TransactionId string
	UserId        string
	Name          string
	Image         string
	IsBank        bool
}
