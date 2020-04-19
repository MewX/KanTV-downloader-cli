package util

import (
	"testing"
)

func TestExtractM3u8BaseURL(t *testing.T) {
	var tests = []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "empty",
			input:  "",
			output: "", // Error out.
		},
		{
			name:   "normal",
			input:  "http://bny.do1byvision.com/cn/movie/1006783/1006783-1/play.m3u8?md5=XLuHvVdvSVXt6Pn-W0V-kg&expires=1587315419&token=1006783",
			output: "http://bny.do1byvision.com/cn/movie/1006783/1006783-1/",
		},
		{
			name:   "no m3u9 keyword",
			input:  "http://bny.do1byvision.com/cn/movie/1006783/1006783-1/play.m3u9?md5=XLuHvVdvSVXt6Pn-W0V-kg&expires=1587315419&token=1006783",
			output: "", // Error out.
		},
		{
			name:   "URL structure not correct",
			input:  "http:play.m3u8/blablabla?md5=XLuHvVdvSVXt6Pn-W0V-kg&expires=1587315419&token=1006783",
			output: "", // Error out.
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := ExtractM3u8BaseURL(tt.input)
			if tt.output == "" {
				if err == nil {
					t.Fatalf("both output and err are empty")
				}
			} else {
				if err != nil {
					t.Fatalf("both output and err are not empty")
				} else if ans != tt.output {
					t.Fatalf("%s != %s", ans, tt.output)
				}
			}
		})
	}
}

func TestExtractTvidPartidFromURL(t *testing.T) {
	var tests = []struct {
		name   string
		input  string
		tvid   string
		partid string
	}{
		{
			name:   "empty",
			input:  "",
			tvid:   "", // Error out.
			partid: "",
		},
		{
			name:   "anime",
			input:  "https://www.wekan.tv/anime/301813564723005",
			tvid:   "301813564723005",
			partid: "",
		},
		{
			name:   "anime with partid",
			input:  "https://www.wekan.tv/anime/301813564723005-101813564723007",
			tvid:   "301813564723005",
			partid: "101813564723007",
		},
		{
			name:   "tvdrama",
			input:  "https://www.wekan.tv/tvdrama/301940350879001",
			tvid:   "301940350879001",
			partid: "",
		},
		{
			name:   "movie",
			input:  "https://www.wekan.tv/movie/302002655075001",
			tvid:   "302002655075001",
			partid: "",
		},
		{
			name:   "nothing",
			input:  "https://www.wekan.tv/nothing/302002655075001",
			tvid:   "", // Error out.
			partid: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tv, part, err := ExtractTvidPartidFromURL(tt.input)
			if tt.tvid == "" {
				if err == nil {
					t.Fatalf("both output and err are empty")
				}
			} else {
				if err != nil {
					t.Fatalf("both output and err are not empty")
				} else if tv != tt.tvid && part != tt.partid {
					t.Fatalf("%s != %s || %s != %s", tv, tt.tvid, part, tt.partid)
				}
			}
		})
	}
}
