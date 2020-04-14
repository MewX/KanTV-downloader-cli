package kantv

import (
	"fmt"
	"log"
	"os"

	"github.com/MewX/KanTV-downloader-cli/kantv/api"
	"github.com/urfave/cli"
)

// Cli is the entry of the command line interface.
func Cli() {
	app := cli.NewApp()
	app.Name = "KanTV Downloader CLI"
	app.Version = "0.1.0"
	app.Usage = "The downloader for downloading KanTV videos via command line."
	app.Action = func(c *cli.Context) error {
		fmt.Println("To be done!")
		return nil
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "email",
			Usage: "Your email address registered in Kantv website. Alternatively, you could use mobile phone number.",
		},
		cli.StringFlag{
			Name:  "phone_number",
			Usage: "Your phone number registered in Kantv website. Alternatively, you could use email address.",
		},
		cli.StringFlag{
			Name:  "password",
			Usage: "Your password for logging in. If not specified, you will be prompted to type your password.",
		},
		cli.StringFlag{
			Name:  "set",
			Usage: "Set a configurable item. e.g. outdir=~/kantv",
		},
		cli.BoolFlag{
			Name:  "anonymous",
			Usage: "Add this flag to not use your login credentials and keeps anonymous.",
		},
		cli.StringFlag{
			Name:  "proxy",
			Usage: "Specify a proxy. http/https/socks5://<address>.",
		},
		cli.StringFlag{
			Name:  "outdir",
			Value: "Output directory. File name will be generated based on the resource you are downloading.",
		},
		cli.StringFlag{
			Name:  "cookies",
			Value: "Specify the cookies.txt file path.",
		},
	}

	app.Commands = []cli.Command{
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
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: delete this stub
	s := api.NewSign()
	fmt.Println(s)
	api.SendRequest(api.NewGetCountryRequest())
}
