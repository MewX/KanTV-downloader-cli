package kantv

// TODO(#22): refactor those json parser codes into a neat and easy way.

import (
	"encoding/json"
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

	// Print as string.
	if util.VerboseMode {
		b, _ := json.MarshalIndent(obj, "", "  ")
		fmt.Println("==== Parsed Struct obj in JSON format ====")
		fmt.Println(string(b))
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

		// Print results.
		f := "%s: %s (%s)\n"
		fmt.Printf(f, "iso_code", "name_cn", "name_en")
		for _, v := range obj["data"].([]interface{}) {
			m := v.(map[string]interface{})
			fmt.Printf(f, m["iso_code"].(string), m["name_cn"].(string), m["name"].(string))
		}
		return nil
	},
}

// CmdDownload defines the downloading-related commands.
var CmdDownload = &cli.Command{
	Name:    "download",
	Aliases: []string{"d", "dl"},
	Usage: "Download a video with a tvid or specifying an URL, plus other options.\n" +
		"Find more using command: $ kantv d --help",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  FlagURL,
			Usage: "Specify the URL of the video to download.",
		},
		&cli.StringFlag{
			Name:  FlagTvid,
			Usage: "Specify the tvid of the video to download.",
		},
	},
	Action: func(c *cli.Context) error {
		// Must specify at least one of tvid or video URL.
		url := c.String(FlagURL)
		tvid := c.String(FlagTvid)
		if url == "" && tvid == "" {
			return fmt.Errorf("you must specify at least one of tvid or video URL")
		}

		// If both specified.
		if url != "" && tvid != "" {
			return fmt.Errorf("please do not specify both tvid and video URL")
		}

		// URL is specified
		if url != "" {
			fmt.Println("Working in progress")
			// TODO: need to extract and set tvid from URL
			// If unable to extract, through an error as well!

			// Here are some examples:
			// https://www.wekan.tv/tvdrama/302014371619001
			// https://www.wekan.tv/movie/302002655075001
		}

		// Tvid is specified
		var obj, err = api.SendRequest(api.NewGetVideoInfoRequest(tvid))
		err = checkError(err, obj)
		if err != nil {
			return err
		}

		// Print results.
		data := obj["data"].(map[string]interface{})
		generalInfo := data["info"].(map[string]interface{})
		playInfo := data["playinfo"].(map[string]interface{})
		fmt.Printf("Downloading: %s...\n", generalInfo["title"].(string))
		fmt.Printf("Will download from this link: %s\n", playInfo["url"].(string))

		// TODO(#8): add download logic and allow specify output dir (allow creating).
		return nil
	},
}
