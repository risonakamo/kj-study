// general utils

package utils

import "strings"

// trim common whitespace from target string
func TrimWhitespace(text string) string {
	return strings.Trim(text," \n\r")
}