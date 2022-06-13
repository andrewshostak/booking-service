package server

import (
	"fmt"
	"github.com/andrewshostak/booking-service/handler"
	"github.com/andrewshostak/booking-service/repository"
	"github.com/andrewshostak/booking-service/service"
	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

type serverConfig struct {
	Port            string `env:"PORT" envDefault:"8080"`
	PgHost          string `env:"PG_HOST" envDefault:"localhost"`
	PgUser          string `env:"PG_USER" envDefault:"postgres"`
	PgPassword      string `env:"PG_PASSWORD"`
	PgPort          string `env:"PG_PORT" envDefault:"5432"`
	PgDatabase      string `env:"PG_DATABASE" envDefault:"postgres"`
	LaunchpadApiUrl string `env:"LAUNCHPAD_API_URL" envDefault:"https://api.spacexdata.com"`
}

func StartServer() {
	config := serverConfig{}
	if err := env.Parse(&config); err != nil {
		panic(err)
	}

	r := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("date", handler.ValidateDate)
	}

	connectionParams := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s database=%s sslmode=disable",
		config.PgHost,
		config.PgUser,
		config.PgPassword,
		config.PgPort,
		config.PgDatabase,
	)
	db, err := gorm.Open(postgres.Open(connectionParams))
	if err != nil {
		panic(err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}

	driver, err := migratepg.WithInstance(sqlDb, &migratepg.Config{})
	m, err := migrate.NewWithDatabaseInstance("file://./migrations", config.PgDatabase, driver)
	err = m.Up()

	if err != nil && err != migrate.ErrNoChange {
		panic(err)
	}

	httpClient := http.Client{}

	launchpadRepository := repository.NewLaunchpadRepository(config.LaunchpadApiUrl, &httpClient)
	bookingRepository := repository.NewBookingRepository(db)
	bookingService := service.NewBookingService(bookingRepository, launchpadRepository)
	bookingHandler := handler.NewBookingHandler(bookingService)

	r.POST("/bookings", bookingHandler.Create)
	r.GET("/bookings", bookingHandler.List)
	r.DELETE("/bookings/:id", bookingHandler.Delete)

	r.Run(fmt.Sprintf(":%s", config.Port))
}
