package port

type AccountsRepo interface {
	GetByUserId(userId string) ([]AccountRepoRes, error)
}

const (
	AccountsTbl        = "accounts"
	AccountFlagsTbl    = "account_flags"
	AccountDetailsTbl  = "account_details"
	AccountBalancesTbl = "account_balances"
)

type AccountRepoRes struct {
	AccountId     string
	Type          string
	Currency      string
	AccountNumber string
	Issuer        string
	Amount        float64
	Color         string
	IsMainAccount bool
	Progress      int64
	FlagId        int64
	FlagType      string
	FlagValue     string
}
