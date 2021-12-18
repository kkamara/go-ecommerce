package commands

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/urfave/cli/v2"
)

var chromeDriver string = "https://chromedriver.storage.googleapis.com/96.0.4664.45/chromedriver_linux64.zip"

func DownloadTestBinaries(c *cli.Context) error {

	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	out, err := os.Create(currentDir + "/vendor/chromedriver.zip")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	resp, err := http.Get(chromeDriver)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully downloaded chrome driver.")
	return nil
}
