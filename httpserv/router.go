package httpserv

import (
	"account-service/infrastructure"
	"account-service/internal/adaptor/handler"
	repository "account-service/internal/adaptor/repo"
	"account-service/internal/core/service"

	libmiddleware "github.com/basputtipong/library/middleware"
	"github.com/gin-gonic/gin"
)

func bindGetAccountRoute(app *gin.Engine) {
	repo := repository.NewAccountsRepo(infrastructure.DB)
	svc := service.NewAccountSvc(repo)
	hdl := handler.NewAccountHandler(svc)

	app.GET("/account", libmiddleware.JWTVerify(), hdl.Handle)
}

func bindGetTransactionRoute(app *gin.Engine) {
	repo := repository.NewTransactionRepo(infrastructure.DB)
	svc := service.NewTransactionSvc(repo)
	hdl := handler.NewTransactionHandler(svc)

	app.GET("/transactions", libmiddleware.JWTVerify(), hdl.Handle)
}

func bindUpdateAccountRoute(app *gin.Engine) {
	repo := repository.NewAccountsRepo(infrastructure.DB)
	svc := service.NewUpdateAccountSvc(repo)
	hdl := handler.NewUpdateAccountHandler(svc)

	app.PUT("/update-account", libmiddleware.JWTVerify(), hdl.Handle)
}

func bindHelthRoute(app *gin.Engine) {
	app.GET("/health", handler.HealthHandle)
}
