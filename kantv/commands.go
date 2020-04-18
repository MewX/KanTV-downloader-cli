package kantv

// TODO(#22): refactor those json parser codes into a neat and easy way.

import (
	"fmt"

	"github.com/MewX/KanTV-downloader-cli/kantv/api"
	"github.com/MewX/KanTV-downloader-cli/kantv/util"
	"github.com/urfave/cli/v2"
)

func checkError(err error, obj map[string]interface{}) error {
	if err != nil {
		return err
	}

	if obj == nil {
		return fmt.Errorf("receiving empty object")
	}

	if val, ok := obj["status_code"].(int); ok {
		if !ok {
			return fmt.Errorf("receiving broken object from server")
		}

		if val != 200 {
			return fmt.Errorf("server response code is not 200")
		}
	}

	if util.VerboseMode {
		fmt.Println("receiving 200 code from server: " + obj["msg"].(string))
	}
	return nil
}

// CmdCountry defines the country-related commands.
var CmdCountry = &cli.Command{
	Name:  "country",
	Usage: "Get country list.",
	Action: func(c *cli.Context) error {
		var obj, err = api.SendRequest(api.NewGetCountryRequest())
		err = checkError(err, obj)
		if err != nil {
			return err
		}

		// Print as string.
		if util.VerboseMode {
			fmt.Printf("%+v\n", obj)
		}
		f := "%s: %s (%s)\n"
		fmt.Printf(f, "iso_code", "name_cn", "name_en")
		for _, v := range obj["data"].([]interface{}) {
			m := v.(map[string]interface{})
			fmt.Printf(f, m["iso_code"].(string), m["name_cn"].(string), m["name"].(string))
		}
		return nil
	},
}
