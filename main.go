package main

import (
	"account-service/httpserv"
	"account-service/infrastructure"
)

func init() {
	infrastructure.InitConfig()
}

func main() {
	infrastructure.InitMiddleware()
	infrastructure.InitDB()
	httpserv.Run()
}
