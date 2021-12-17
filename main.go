package main

import (
	"log"
	"os"

	"github.com/kkamara/ecommerce/commands"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

func main() {
	godotenv.Load()
	app := &cli.App{
		Name:  "ecommerce",
		Usage: "serve this app on the web",
		Action: func(c *cli.Context) error {
			engine := html.New("./views", ".html")

			if os.Getenv("ENV") != "production" {
				engine.Reload(true)
				engine.Debug(true)
			}

			engine.AddFunc("appName", func(name string) string {
				return os.Getenv("APP_NAME")
			})

			webApp := fiber.New(fiber.Config{
				Views: engine,
			})
			webApp.Get("/", func(c *fiber.Ctx) error {
				return c.Render("index", fiber.Map{
					"Title": "Hello, World!",
				})
			})

			log.Fatal(webApp.Listen(":8080"))
			return nil
		},
		Authors: []*cli.Author{{Name: "Kelvin Kamara", Email: "kelvinkamara@protonmail.com"}},
	}
	app.Commands = []*cli.Command{
		{
			Name:   "migrate",
			Usage:  "run database migrations",
			Action: commands.DbMigrate,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}

}
