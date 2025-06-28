package goreloaded

import (
	"fmt"
	"strconv"
	"strings"
)

func Convert(s []string) []string {
	for i := 0; i < len(s); i++ {
		// word := strings.ToLower(s[i])

		if i > 0 && (s[i] == "(up)") {
			s[i-1] = strings.ToUpper(s[i-1])
			s[i] = ""
			s=Rmov(s)
			i--
		} else if i == 0 && (s[i] == "(up)"){
			s[i] = ""
			s=Rmov(s)
		i--
		} else if i > 0 && (s[i] == "(low)" ) {
			s[i-1] = strings.ToLower(s[i-1])
			s[i] = ""
			s=Rmov(s)
		i--
		}else if i == 0 && (s[i] == "(low)"){
			s[i] = ""
			s=Rmov(s)
			i--
		} else if i > 0 && (s[i] == "(cap)" ) {
			s[i-1] = Capitalize(s[i-1])
			s[i] = ""
			s=Rmov(s)
		i--
		}else if i == 0 && (s[i] == "(cap)"){
			s[i] = ""
			s=Rmov(s)
			i--
		}else if i == 0 && (s[i] == "(hex)"){
			s[i] = ""
			s=Rmov(s)
			i--
		} else if i > 0 && (s[i] == "(hex)" ) {
			decimal, err := strconv.ParseInt(s[i-1], 16, 64)
			if err != nil {
				fmt.Println("Error Case:", err)
			} else {
				s[i-1] = strconv.Itoa(int(decimal))
			}
			s[i] = ""
		s=Rmov(s)
		i--
		} else if i > 0 && (s[i] == "(bin)" ) {
			decimal, err := strconv.ParseInt(s[i-1], 2, 64)
			if err != nil {
				fmt.Println("Error Case:", err)
			} else {
				s[i-1] = strconv.Itoa(int(decimal))
			}
			s[i] = ""
			s=Rmov(s)
			i--
		}else if i == 0 && (s[i] == "(bin)"){
			s[i] = ""
			s=Rmov(s)
			i--
		}

	}
	return s
	
}
func Rmov(s []string)[]string{

	var result []string
	for _, word := range s {
		if word != "" {
			result = append(result, word)
		}
	}
	return result
}
