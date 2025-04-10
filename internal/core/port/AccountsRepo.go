package port

type AccountsRepo interface {
	GetByUserId(userId string) ([]AccountRepoRes, error)
	GetFlagByAccountId(accountIds []string) ([]Flag, error)
	GetCurrentMainAccountByUserId(userId string) (AccountRepoRes, error)
	UpdateAccountById(req UpdateAccountRepoReq) error
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
}

type Flag struct {
	FlagId    int64  `gorm:"column:flag_id"`
	AccountId string `gorm:"column:account_id"`
	FlagType  string `gorm:"column:flag_type"`
	FlagValue string `gorm:"column:flag_value"`
}

func (Flag) TableName() string {
	return AccountFlagsTbl
}

type UpdateAccountRepoReq struct {
	UserId           string
	AccountId        string
	CurrentMainAccId string
	IsMainAccount    bool
	Color            string
}
