package main

import (
	"github.com/d-cryptic/crm-golang-backend/config"
	"github.com/d-cryptic/crm-golang-backend/routes"
	"github.com/d-cryptic/crm-golang-backend/utils"
)

func main() {
	utils.LoadEnv()
	config.ConnectDB()
	r := routes.SetupRouter()
	r.Run(":8080")
}