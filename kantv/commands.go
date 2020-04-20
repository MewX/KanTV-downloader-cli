package kantv

// TODO(#22): refactor those json parser codes into a neat and easy way.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"

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
		&cli.StringFlag{
			Name:  FlagPartid,
			Usage: "Specify the partid of the video to download.",
		},
		&cli.StringFlag{
			Name:  FlagOutDir,
			Usage: "Specify the output dir.",
		},
	},
	Action: func(c *cli.Context) error {
		// Must specify at least one of tvid or video URL.
		url := c.String(FlagURL)
		tvid := c.String(FlagTvid)
		partid := c.String(FlagPartid)
		outdir := c.String(FlagOutDir)
		if url == "" && tvid == "" {
			return fmt.Errorf("you must specify at least one of tvid or video URL")
		}

		// If both specified.
		if url != "" && tvid != "" {
			return fmt.Errorf("please do not specify both tvid and video URL")
		}

		// URL is specified
		if url != "" {
			var err error
			tvid, partid, err = util.ExtractTvidPartidFromURL(url)
			if err != nil {
				return err
			}
			fmt.Printf("Parsed tvid: %s; partid: %s\n", tvid, partid)
		}

		// Tvid is specified
		var obj, err = api.SendRequest(api.NewGetVideoInfoRequest(tvid, partid))
		err = checkError(err, obj)
		if err != nil {
			return err
		}

		// Print results.
		data := obj["data"].(map[string]interface{})
		generalInfo := data["info"].(map[string]interface{})
		videoTitle := generalInfo["title"].(string)
		playInfo := data["playinfo"].(map[string]interface{})
		partID := int(playInfo["part"].(float64))
		m3u8URL := playInfo["url"].(string)
		fmt.Printf("Downloading: %s...\n", videoTitle)
		fmt.Printf("Will download from this link: %s\n", m3u8URL)

		// Generate base URL.
		baseURL, errBaseURL := util.ExtractM3u8BaseURL(m3u8URL)
		if errBaseURL != nil {
			return errBaseURL
		}

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

		// Download all videos.
		// TODO(#8): Need to refactor this codes. Allow specifying file name.
		var wd string
		if outdir == "" {
			// Using working directory by default.
			wd, err = os.Getwd()
			if err != nil {
				return err
			}
		} else {
			wd = outdir
			_ = os.MkdirAll(wd, 0744)
		}
		// TODO(#8): Added support for part ID. They should be saved in the same folder.
		folderName := util.SanitizeFileName(videoTitle)
		_ = os.Mkdir(path.Join(wd, folderName), 0744)
		if util.VerboseMode {
			fmt.Println(path.Join(wd, folderName))
		}

		fmt.Println("Saving all files to folder: " + folderName)
		playlist := p.(*m3u8.MediaPlaylist)
		for i, segment := range playlist.Segments {
			if segment != nil {
				fullURL := baseURL + segment.URI
				if util.VerboseMode {
					fmt.Println("\nDownloading from: " + fullURL)
				}

				startPercentage := float64(i) / float64(len(playlist.Segments)) * 100.0
				finishPercentage := float64(i+1) / float64(len(playlist.Segments)) * 100.0
				fileName := strconv.Itoa(partID) + "-" + util.SanitizeFileName(segment.URI)
				filePath := path.Join(wd, folderName, fileName)
				fileInfo, errStat := os.Stat(filePath)
				fmt.Printf("\r(%.2f%%) Downloading: %s", startPercentage, fileName)
				if errStat != nil && os.IsNotExist(errStat) {
					// Need to download.
					fmt.Printf("\r(%.2f%%) Downloading: %s", finishPercentage, fileName)
					b, e := util.FetchLinkContentWithRetry(fullURL)
					if e != nil {
						return fmt.Errorf("unable to download %s even after %d times of retry",
							fullURL, util.RetryTimes)
					}
					err := ioutil.WriteFile(filePath, b, 0744)
					if err != nil {
						return err
					}
				} else if errStat != nil {
					return errStat
				} else if fileInfo.IsDir() {
					return fmt.Errorf("file name is taken by a dir: %s", filePath)
				} else {
					// Downloaded already. Do nothing here.
				}
			} else {
				// Should be an exception?
				if util.VerboseMode {
					fmt.Printf("Saw a nil segment at index: %d\n", i)
				}
			}
		}
		fmt.Println()
		return nil
	},
}
