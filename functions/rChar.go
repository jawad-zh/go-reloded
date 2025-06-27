package goreloaded

import "strings"

func Rchar(clean []string) []string {
	for i := range clean {
		clean[i] = strings.ReplaceAll(clean[i], "\r", "")
	}
	var result []string
	for _, word := range clean {
		if word != "" {
			result = append(result, word)
		}
	}
	return result

}
