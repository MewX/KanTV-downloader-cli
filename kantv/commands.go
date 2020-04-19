package kantv

// TODO(#22): refactor those json parser codes into a neat and easy way.

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/MewX/KanTV-downloader-cli/kantv/api"
	"github.com/MewX/KanTV-downloader-cli/kantv/util"
	"github.com/grafov/m3u8"
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
		m3u8URL := playInfo["url"].(string)
		fmt.Printf("Downloading: %s...\n", generalInfo["title"].(string))
		fmt.Printf("Will download from this link: %s\n", m3u8URL)

		// Download m3u8 playlist.
		b, errM3u8 := util.FetchLinkContent(m3u8URL)
		if errM3u8 != nil {
			return errM3u8
		}

		// Decode the m3u8 playlist.
		p, listType, errPlaylist := m3u8.DecodeFrom(bytes.NewReader(b), true)
		if errPlaylist != nil {
			return errPlaylist
		}

		// Download the video files.
		if util.VerboseMode {
			switch listType {
			case m3u8.MEDIA:
				mediapl := p.(*m3u8.MediaPlaylist)
				fmt.Printf("MediaPL:\n%+v\n", mediapl)
			case m3u8.MASTER:
				masterpl := p.(*m3u8.MasterPlaylist)
				fmt.Printf("MasterPL:\n%+v\n", masterpl)
			}
		}

		// Expect to receive Media Playlist.
		if listType != m3u8.MEDIA {
			return fmt.Errorf("please report this error, the server returns a Master playlist")
		}
		return nil
	},
}
