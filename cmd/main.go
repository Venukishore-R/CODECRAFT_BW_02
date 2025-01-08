package main

import (
	"log"

	"github.com/Venukishore-R/CODECRAFT_BW_02/config"
	"github.com/Venukishore-R/CODECRAFT_BW_02/endpoints"
	"github.com/Venukishore-R/CODECRAFT_BW_02/internal/app/models"
	"github.com/Venukishore-R/CODECRAFT_BW_02/internal/app/services"
	"github.com/Venukishore-R/CODECRAFT_BW_02/internal/database"
	"github.com/Venukishore-R/CODECRAFT_BW_02/pkg/handlers"
	"github.com/gin-gonic/gin"
)

var (
	SERVER_PORT = ":5001"
)

func main() {
	r := gin.Default()

	config := config.LoadEnv()

	db, err := database.InitDB(config)
	if err != nil {
		log.Fatalf("error initializing database: %v", err)
		return
	}

	handler := handlers.NewHandler(services.NewService(models.NewUserModel(db)))

	endpoints.Endpoints(r, handler)

	r.Run(SERVER_PORT)
}
