package server

import (
	"fmt"
	"github.com/andrewshostak/booking-service/handler"
	"github.com/andrewshostak/booking-service/service"
	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type serverConfig struct {
	Port       string `env:"PORT" envDefault:"8080"`
	PgHost     string `env:"PG_HOST" envDefault:"localhost"`
	PgUser     string `env:"PG_USER" envDefault:"postgres"`
	PgPassword string `env:"PG_PASSWORD"`
	PgPort     string `env:"PG_PORT" envDefault:"5432"`
	PgDatabase string `env:"PG_DATABASE" envDefault:"postgres"`
}

func StartServer() {
	config := serverConfig{}
	if err := env.Parse(&config); err != nil {
		panic(err)
	}

	r := gin.Default()

	connectionParams := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s database=%s sslmode=disable",
		config.PgHost,
		config.PgUser,
		config.PgPassword,
		config.PgPort,
		config.PgDatabase,
	)
	_, err := gorm.Open(postgres.Open(connectionParams))
	if err != nil {
		panic(err)
	}

	bookingService := service.NewBookingService()
	bookingHandler := handler.NewBookingHandler(bookingService)

	r.POST("/bookings", bookingHandler.Create)
	r.GET("/bookings", bookingHandler.List)
	r.DELETE("/bookings/:id", bookingHandler.Delete)

	r.Run(fmt.Sprintf(":%s", config.Port))
}
