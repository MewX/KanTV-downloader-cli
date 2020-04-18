package kantv

import (
	"fmt"

	"github.com/MewX/KanTV-downloader-cli/kantv/api"
	"github.com/urfave/cli/v2"
)

// CmdCountry defines the country-related commands.
var CmdCountry = &cli.Command{
	Name:  "country",
	Usage: "Get country list.",
	Action: func(c *cli.Context) error {
		var obj, err = api.SendRequest(api.NewGetCountryRequest())
		if err != nil {
			return err
		}

		// Print as string.
		fmt.Printf("%+v\n", obj)
		return nil
	},
}
