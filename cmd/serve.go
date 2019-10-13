package cmd

import (
	"context"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go-boilerplate-project/app"
	"go-boilerplate-project/config"
	"os"
	"os/signal"
	"time"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run application server",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve() {
	e := echo.New()

	application := InitializeApplication()

	e.Use(middleware.Logger())
	e.Use(middleware.Secure())

	app.InitApi(e, application)

	go func() {
		address := fmt.Sprintf("%s:%d", config.Const.Server.Host, config.Const.Server.Port)
		if err := e.Start(address); err != nil {
			logrus.Fatal(fmt.Sprintf("server error, quiting..., error: %s\n", err.Error()))
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
