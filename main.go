package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kkamara/go-ecommerce/commands"
	"github.com/kkamara/go-ecommerce/database"

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
				db, err := database.Open()
				if err != nil {
					return err
				}
				products := []database.Product{}
				result := db.Find(&products)
				if result.Error != nil {
					return result.Error
				}
				fmt.Printf("%+v", products)
				return c.Render("product/index", fiber.Map{
					"Title":    "Home",
					"Products": products,
				}, "layouts/master")
			})

			port := os.Getenv("APP_PORT")
			if len(port) < 1 {
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
		{
			Name:  "test",
			Usage: "test some stuff",
			Action: func(c *cli.Context) error {
				// do stuff
				_, err := database.Open()
				if err != nil {
					return err
				}

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}

}
