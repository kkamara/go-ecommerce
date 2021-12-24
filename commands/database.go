package commands

import (
	"database/sql"
	"fmt"

	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/urfave/cli/v2"
)

// enable tls support
// https://stackoverflow.com/questions/42163732/how-to-use-the-go-mysql-driver-with-ssl-on-aws-with-a-mysql-rds-instance

func DbMigrate(c *cli.Context) (err error) {
	fmt.Print("test\n")
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	var (
		host string = os.Getenv("DB_HOST")
		name string = os.Getenv("DB_NAME")
		user string = os.Getenv("DB_USER")
		pass string = os.Getenv("DB_PASS")
	)
	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?multiStatements=true",
		user,
		pass,
		host,
		name,
	))
	if err != nil {
		return err
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s/database/migrations", currentDir),
		"mysql",
		driver,
	)
	if err != nil {
		return err
	}

	m.Up()

	return err
}
