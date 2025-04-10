package handler

import (
	"account-service/internal/core/domain"
	"net/http"

	liberror "github.com/basputtipong/library/error"
	"github.com/gin-gonic/gin"
)

type updateAccountHandler struct {
	svc domain.UpdateAccountService
}

func NewUpdateAccountHandler(svc domain.UpdateAccountService) *updateAccountHandler {
	return &updateAccountHandler{svc: svc}
}

func (h *updateAccountHandler) Handle(c *gin.Context) {
	var req domain.UpdateAccReq

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
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(liberror.ErrorBadRequest("Invalid request", err.Error()))
		return
	}

	res, err := h.svc.Execute(req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, &res)
}
