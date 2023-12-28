package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	_ "github.com/swaggo/echo-swagger/example/docs"
	"tesk-task-betera/models"
	"tesk-task-betera/pkg/apodWorker"
	"tesk-task-betera/repository"
	"tesk-task-betera/routes"
	"tesk-task-betera/service"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	if err := initConfig(); err != nil {
		e.Logger.Fatalf("error initializing configs: %v\n", err)
	}

	pool, err := repository.NewPostgresDB(context.Background(), repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	})

	if err != nil {
		e.Logger.Fatalf("unable to connect to database: %v\n", err)
	}

	postgresConnection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		viper.GetString("db.username"), viper.GetString("db.password"), viper.GetString("db.host"),
		viper.GetString("db.port"), viper.GetString("db.dbname"), viper.GetString("db.sslmode"))

	err = repository.RunMigrations(postgresConnection, "file://./migrations")
	if err != nil {
		e.Logger.Fatalf("Unable to connect to database: %v\n", err)
	}

	apodRepository := repository.NewApodRepository(pool)
	apodService := service.NewApodService(apodRepository)
	apodHandler := routes.NewApodHandler(apodService)

	nasaApodWorker, apodChan, err := createApodWorker(viper.GetInt("apiNasa.periodicitySeconds"), viper.GetString("apiNasa.url"))
	if err != nil {
		e.Logger.Errorf("Unable to create apodWorker: %v", err)
	}

	go startApodDataFetcher(e.Logger, nasaApodWorker)
	go startDataProcessing(e.Logger, apodChan, apodService)

	routes.SetupRoutes(apodHandler, e)

	e.Logger.Fatal(e.Start(":8080"))

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func createApodWorker(Periodicity int, NasaUrl string) (*apodWorker.ApodWorker, *chan models.ApodChanelType, error) {
	apodChan := make(chan models.ApodChanelType)
	nasaApodWorker := apodWorker.NewApodWorker(Periodicity, viper.GetString("apiNasa.key"), NasaUrl, apodChan)

	return nasaApodWorker, &apodChan, nil
}

func startApodDataFetcher(logger echo.Logger, nasaApodWorker *apodWorker.ApodWorker) {
	err := nasaApodWorker.GetApodData()
	if err != nil {
		logger.Errorf("Error fetching APOD data: %v", err)
	}
}

func startDataProcessing(logger echo.Logger, apodChan *chan models.ApodChanelType, apodService *service.Apod) {
	for astronomyData := range *apodChan {
		if astronomyData.Error != nil {
			logger.Errorf("Error fetching APOD data: %v", astronomyData.Error)
		} else {
			err := apodService.CreateApod(context.Background(), astronomyData.ApodData)
			if err != nil {
				logger.Errorf("Error creating APOD data: %v", err)
			}
		}
	}
}
