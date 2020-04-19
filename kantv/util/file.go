package util

import (
	"regexp"
)

// SanitizeFileName replaces illegal characters in the file.
func SanitizeFileName(fn string) string {
	m := regexp.MustCompile(`[\\\\/:*?\"<>|]`)
	return m.ReplaceAllString(fn, "_")
}
