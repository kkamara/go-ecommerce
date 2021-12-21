package main

import (
	"log"
	"os"

	"github.com/kkamara/go-ecommerce/commands"

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
			engine := html.New("./views", ".gohtml")

			if os.Getenv("ENV") != "production" {
				engine.Reload(true)
				engine.Debug(true)
			}

			engine.AddFunc("appName", func() string {
				return os.Getenv("APP_NAME")
			})

			engine.AddFunc("cartCount", func() string {
				return "0"
			})

			webApp := fiber.New(fiber.Config{
				Views: engine,
			})

			webApp.Static("/", "./resources")

			webApp.Get("/", func(c *fiber.Ctx) error {
				return c.Render("product/index", fiber.Map{
					"Title": "Home",
				}, "layouts/master")
			})

			port := os.Getenv("APP_PORT")
			if port == "" {
				port = os.Getenv("PORT")
				if port == "" {
					port = "5000"
				}
			}

			log.Fatal(webApp.Listen(":" + port))
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
		{
			Name:   "download",
			Usage:  "download test binaries",
			Action: commands.DownloadTestBinaries,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}

}
