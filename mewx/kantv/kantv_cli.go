package kantv

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func Cli() {
	app := cli.NewApp()
	app.Name = "KanTV Downloader CLI"
	app.Version = "0.1.0"
	app.Usage = "The downloader for downloading KanTV videos via command line."
	app.Action = func(c *cli.Context) error {
		fmt.Println("To be done!")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
