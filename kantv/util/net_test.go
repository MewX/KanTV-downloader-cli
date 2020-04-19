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
