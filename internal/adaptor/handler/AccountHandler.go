package handler

import (
	"account-service/internal/core/domain"
	"net/http"

	liberror "github.com/basputtipong/library/error"
	"github.com/gin-gonic/gin"
)

type accountHandler struct {
	svc domain.AccountService
}

func NewAccountHandler(svc domain.AccountService) *accountHandler {
	return &accountHandler{svc: svc}
}

func (h *accountHandler) Handle(c *gin.Context) {
	var req domain.AccountReq

	userIdRaw, ok := c.Get("user_id")
	if !ok {
		c.Error(liberror.ErrorBadRequest("Invalid request", "user_id missing from context"))
		return
	}

	userId, ok := userIdRaw.(string)
	if !ok {
		c.Error(liberror.ErrorBadRequest("Invalid request", "user_id must be string"))
		return
	}

	req.UserId = userId
	res, err := h.svc.Execute(req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, &res)
}
