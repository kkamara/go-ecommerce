package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func TestExample(t *testing.T) {

	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	godotenv.Load(currentDir + "/.env")

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "chrome"}
	chromeCaps := chrome.Capabilities{
		Path: "",
		Args: []string{
			"--headless",
			"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36",
		},
	}
	caps.AddChrome(chromeCaps)

	host := fmt.Sprintf(
		"http://127.0.0.1:%s/wd/hub",
		os.Getenv("SELENIUM_PORT"),
	)
	if os.Getenv("ENV") == "testing" {
		host = fmt.Sprintf("http://127.0.0.1:%s", os.Getenv("SELENIUM_PORT"))
	}

	wd, err := selenium.NewRemote(caps, host)
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	if err := wd.Get("http://google.com"); err != nil {
		panic(err)
	}

	_, err = wd.FindElement(selenium.ByXPATH, "//input[@value='Google Search']")
	if err != nil {
		panic(err)
	}

	img, err := wd.Screenshot()
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(
		fmt.Sprintf("%s/tests/screenshots/screenie.jpg", currentDir),
		img,
		0644,
	)
	if err != nil {
		panic(err)
	}

}
