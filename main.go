package main

import (
	"go-study-admin/models"
	"go-study-admin/router"
)

func main() {
	models.NewGormDB()
	err := router.App().Run(":9000")
	if err != nil {
		return
	}
}
