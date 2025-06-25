package goreloaded

import (
	"strings"
)

func Punc(s []string)[]string{

	for i := 0; i < len(s); i++ {
	if Only(s[i]) && i > 0 {
		if strings.HasSuffix(s[i-1], ")") {
			for j := i - 2; j >= 0; j-- {
				if s[j] != "" && !strings.HasPrefix(s[j], "(") && !strings.HasSuffix(s[j], ")") {
					s[j] = s[j] + s[i]
					s[i] = ""
					break
				}
			}
		} else {
			s[i-1] = s[i-1] + s[i]
			s[i] = ""
		}
	}
}
	var result []string
	for _, word := range s {
		if word != "" {
			result = append(result, word)
		}
	}

	 return result
}

