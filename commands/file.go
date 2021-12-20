package commands

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"

	"github.com/urfave/cli/v2"
)

var chromeDriver string = "https://chromedriver.storage.googleapis.com/96.0.4664.45/chromedriver_linux64.zip"

func DownloadTestBinaries(c *cli.Context) error {

	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	path := currentDir + "/vendor"
	if _, err = os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, fs.ModePerm)
		if err != nil {
			return err
		}
	}

	out, err := os.Create(path + "/chromedriver.zip")
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(chromeDriver)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("Successfully downloaded chrome driver.")
	return nil
}
