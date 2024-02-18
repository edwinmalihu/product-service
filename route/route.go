package route

import (
	"log"
	"product-service/middleware"
	"product-service/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) {
	httpRouter := gin.Default()
	httpRouter.Use(middleware.CORSMiddleware())

	httpRouter.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})

	productRepository := repository.NewProductRepo(db)
	if err := productRepository.Migrate(); err != nil {
		log.Fatal("Product migrate err", err)
	}

	httpRouter.Run(":8082")
}
