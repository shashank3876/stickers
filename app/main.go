package main

import (
	"fmt"
	"log"

	_stickerDelivery "shashank-gusain-backend-onboarding/stickers/delivery"
	_stickerRepo "shashank-gusain-backend-onboarding/stickers/repository/sql"

	HttpDeliveryMiddleware "shashank-gusain-backend-onboarding/stickers/delivery/middleware"
	_stickerUse "shashank-gusain-backend-onboarding/stickers/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service running on debug mode ")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetInt(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	middL := HttpDeliveryMiddleware.InitMiddleware()
	e.Use(middL.CORS)
	ar := _stickerRepo.NewSQLStickerRepository(db)
	au := _stickerUse.NewStickerUsecase(ar)
	_stickerDelivery.NewStickerHandler(e, au)

	log.Fatal(e.Start(viper.GetString("server.address")))

}
