package protocol

import (
	"flag"
	"go-hex-auth/configs"
	"go-hex-auth/database"
	"go-hex-auth/internal/core/services"
	"go-hex-auth/internal/handlers"
	"go-hex-auth/internal/repositories"
	"go-hex-auth/package/logger"
	"go-hex-auth/protocol/routes"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type config struct {
	Env string
}

func Start() error {
	app := fiber.New()
	app.Use(cors.New())

	var cfg config

	flag.StringVar(&cfg.Env, "env", "", "the environment to use")
	flag.Parse()
	configs.InitViper("./configs", cfg.Env)

	dbCon, err := database.ConnectToPostgreSQL(
		configs.GetViper().Postgres.Host,
		configs.GetViper().Postgres.Port,
		configs.GetViper().Postgres.Username,
		configs.GetViper().Postgres.Password,
		configs.GetViper().Postgres.DbName,
		configs.GetViper().Postgres.SSLMode,
	)
	if err != nil {
		logger.Error(err.Error())
		panic(err.Error())
	}
	// Graceful shutdown ...
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Println("Gracefull shut down ...")
			//TODO: close database or any connection before server has gone ...
			database.DisconnectPostgres(dbCon.Postgres)
			err := app.Shutdown()
			if err != nil {
				panic("Can't shutdown")
			}
		}
	}()

	memberRepository := repositories.NewRepository(dbCon.Postgres)
	memberService := services.NewMemberService(memberRepository)
	memberHandlers := handlers.NewMemberHandlers(memberService)

	app.Get("/healthz", func(c *fiber.Ctx) error { return c.SendStatus(fiber.StatusOK) })

	api := app.Group("/api/v1")
	{
		routes.MemberRoutes(api, memberHandlers)
	}

	err = app.Listen(":" + configs.GetViper().App.Port)
	if err != nil {
		return err
	}

	return nil
}
