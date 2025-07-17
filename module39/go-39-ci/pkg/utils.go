package pkg

import "strings"

const vscSeparator = ","

func someHelperFunction(s string) []string {
	cleaned := strings.ReplaceAll(s, ";", "")
	return strings.Split(cleaned, vscSeparator)
}
