package domain

type UpdateAccountService interface {
	Execute(req UpdateAccReq) (UpdateAccRes, error)
}

type UpdateAccReq struct {
	UserId        string `json:"userId" validate:"required"`
	AccountId     string `json:"accountId" validate:"required"`
	IsMainAccount bool   `json:"isMainAccount"`
	Color         string `json:"color"`
}

type UpdateAccRes struct{}
