package kantv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/MewX/KanTV-downloader-cli/kantv/api"
	"github.com/urfave/cli/v2"
)

// Cli is the entry of the command line interface.
func Cli() {
	app := &cli.App{
		Name:    "KanTV Downloader CLI",
		Version: "0.1.0",
		Usage:   "The downloader for downloading KanTV videos via command line.",
		Action: func(c *cli.Context) error {
			fmt.Println("Interactive CLI be done!")
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "email",
				Usage: "Your email address registered in Kantv website. Alternatively, you could use mobile phone number.",
			},
			&cli.StringFlag{
				Name:  "phone_number",
				Usage: "Your phone number registered in Kantv website. Alternatively, you could use email address.",
			},
			&cli.StringFlag{
				Name:  "password",
				Usage: "Your password for logging in. If not specified, you will be prompted to type your password.",
			},
			&cli.StringFlag{
				Name:  "set",
				Usage: "Set a configurable item. e.g. outdir=~/kantv",
			},
			&cli.BoolFlag{
				Name:  "anonymous",
				Usage: "Add this flag to not use your login credentials and keeps anonymous.",
			},
			&cli.StringFlag{
				Name:  "proxy",
				Usage: "Specify a proxy. http/https/socks5://<address>.",
			},
			&cli.StringFlag{
				Name:  "outdir",
				Usage: "Output directory. File name will be generated based on the resource you are downloading.",
			},
			&cli.StringFlag{
				Name:  "cookies",
				Usage: "Specify the cookies.txt file path.",
			},
			&cli.StringFlag{
				Name:  "iso_code",
				Value: "AU",
				Usage: "Specify the country of the account.",
			},
		},

		Commands: []*cli.Command{
			{
				Name:  "country",
				Usage: "Get country list.",
				Action: func(c *cli.Context) error {
					var j, err = api.SendRequest(api.NewGetCountryRequest())

					// Print as json.
					var buf bytes.Buffer
					json.Indent(&buf, []byte(j), "", "  ")
					fmt.Println(buf.String())

					// Print as string.
					var obj map[string]interface{}
					json.Unmarshal([]byte(j), &obj)
					fmt.Printf("%+v\n", obj)
					return err
				},
			},
			{
				Name:  "register",
				Usage: "Register a new account.",
				Action: func(c *cli.Context) error {
					// TODO: jump to website for a lazy implementation
					fmt.Println("registering.")
					return nil
				},
			},
			{
				Name:  "login",
				Usage: "Use your email or phone number to log in.",
				Action: func(c *cli.Context) error {
					// TODO
					fmt.Println("logging in.")
					return nil
				},
			},
			{
				Name:    "info",
				Aliases: []string{"config"},
				Action: func(c *cli.Context) error {
					// TODO: support those flags - n_thread, user_agent, cache_size, max_parallel, max_download_file
					fmt.Println("configuring")
					return nil
				},
			},
			{
				Name:    "download",
				Aliases: []string{"d", "dl"},
				Action: func(c *cli.Context) error {
					// TODO
					fmt.Println("downloading")
					return nil
				},
			},
			{
				Name:    "play",
				Aliases: []string{"p", "pl"},
				Action: func(c *cli.Context) error {
					// TODO
					// This option provides the streaming link or link to a local player to play the video directly.
					fmt.Println("playing...")
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
