package util

import (
	"testing"
)

func TestSanitizeFileName(t *testing.T) {
	var tests = []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "empty",
			input:  "",
			output: "",
		},
		{
			name:   "normal",
			input:  "http://1/play.m3u8?md5=X",
			output: "http___1_play.m3u8_md5=X",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := SanitizeFileName(tt.input)
			if ans != tt.output {
				t.Fatalf("%s != %s", ans, tt.output)
			}
		})
	}
}
