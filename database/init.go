package database

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/bxcodec/faker/v3"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func randomBool() (random bool, err error) {
	tmp, err := faker.RandomInt(0, 1)
	if err != nil {
		return
	}
	if tmp[0] == 1 {
		random = true
	}
	return
}

type Product struct {
	gorm.Model
	CompanyID        uint    `faker:"-"`
	Name             string  `gorm:"size:100;not null" faker:"name"`
	ShortDescription string  `gorm:"size:100;not null" faker:"sentence"`
	LongDescription  string  `gorm:"size:1000" faker:"paragraph"`
	ProductDetails   string  `gorm:"size:1000" faker:"paragraph"`
	ImagePath        string  `gorm:"size:1000" faker:"-"`
	Cost             float64 `gorm:"not null" faker:amount`
	Shippable        bool    `faker:"-"`
	FreeDelivery     bool    `faker:"-"`
}

func seedDB(db *gorm.DB) (err error) {
	shippable, err := randomBool()
	if err != nil {
		return
	}
	freeDelivery, err := randomBool()
	if err != nil {
		return
	}
	for i := 0; i < 10; i++ {
		product := &Product{
			ImagePath:    "/image/products/default/not-found.jpg",
			Shippable:    shippable,
			FreeDelivery: freeDelivery,
		}
		err = faker.FakeData(&product)
		if err != nil {
			err = faker.FakeData(&product)
			if err != nil {
				return
			}
		}
		db.Create(&product)
	}
	return
}

func Open() (db *gorm.DB, err error) {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	fmt.Println(os.Getenv("DB_HOST"))
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	if os.Getenv("MIGRATE_DB") == "true" {
		err = db.AutoMigrate(&Product{})
		if err != nil {
			return
		}
		cmd := exec.Command(
			"sed",
			"-i",
			"s/MIGRATE_DB=true/MIGRATE_DB=false/g",
			".env",
		)
		err = cmd.Run()
		if err != nil {
			return
		}
	}
	if os.Getenv("SEED_DB") == "true" {
		err = seedDB(db)
		if err != nil {
			return
		}
		cmd := exec.Command(
			"sed",
			"-i",
			"s/SEED_DB=true/SEED_DB=false/g",
			".env",
		)
		err = cmd.Run()
		if err != nil {
			return
		}
	}
	if os.Getenv("MIGRATE_DB") == "true" || os.Getenv("SEED_DB") == "true" {
		var currentDir string
		currentDir, err = os.Getwd()
		if err != nil {
			return
		}
		godotenv.Load(currentDir + "/../.env")
	}

	return
}
