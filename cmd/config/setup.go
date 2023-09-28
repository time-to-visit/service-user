package config

import (
	"flag"
	"service-user/cmd/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	var configPath = ""
	configPath = *flag.String("config", "", "")

	if configPath == "" {
		configPath = "./data/config.yml"
	}

	setConfiguration(configPath)
}

func setConfiguration(configPath string) {
	Setup(configPath)
}

func Run(configPath string) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	conf := GetConfig()
	setupDB(conf)
	user, progress := genIoc()
	e = handler.NewHandlerUser(e, user, progress)
	e.Logger.Fatal(e.Start(":" + conf.Server.Port))
}
