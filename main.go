package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	http2 "github.com/klepon46/edot-user-service/delivery/http"
	"github.com/klepon46/edot-user-service/model"
	"github.com/klepon46/edot-user-service/repository"
	"github.com/klepon46/edot-user-service/service"
	"net/http"

	//"github.com/klepon46/edot-user-service/config"
	//"github.com/klepon46/edot-user-service/service"
	//"gorm.io/gorm/logger"
	//"time"
	_ "time/tzdata"

	// Import third parties here
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	_ "github.com/spf13/viper/remote"
)

func main() {
	users := []model.User{}

	userRepo := repository.NewUserRepository(&users)
	repoRegistry := repository.NewRegistryRepository(userRepo)

	userSvc := service.NewUserService(repoRegistry)
	svcRegistry := service.NewRegistry(userSvc)

	delivery := http2.NewUserDelivery(svcRegistry)

	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	engine.POST("/login", delivery.Login)
	engine.POST("/register", delivery.Registry)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", 4001),
		Handler: engine,
	}

	err := srv.ListenAndServe()
	if err != nil {
		return
	}

}
