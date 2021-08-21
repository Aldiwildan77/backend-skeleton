package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Aldiwildan77/backend-skeleton/broker"
	"github.com/Aldiwildan77/backend-skeleton/config"
	"github.com/Aldiwildan77/backend-skeleton/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "boot",
	Short: "Boot service.",
	Run: func(cmd *cobra.Command, args []string) {
		config.Initialize()

		e := echo.New()

		// handle trailing slash
		e.Pre(middleware.RemoveTrailingSlash())

		// handle CORS
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowHeaders: []string{
				echo.HeaderOrigin,
				echo.HeaderContentType,
				echo.HeaderAccept,
				echo.HeaderAuthorization,
			},
		}))

		// handle Routes
		routes.Endpoints(e)

		go func() {
			if err := e.Start(":" + fmt.Sprint(config.Cfg.Port)); err != nil {
				e.Logger.Info("Shutting down the server.")
			}
		}()

		go broker.KafkaInstance()

		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)

		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
