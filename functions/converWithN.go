package goreloaded

import (
	"strconv"
	"strings"
	"fmt"
)

func ConvN(s []string)[]string{
	for i := 0; i < len(s); i++ {
		if IsFlag(s){
			if s[i] == "(up," && i+1 < len(s) {
			nStr := strings.TrimSuffix(s[i+1], ")")
			n, err := strconv.Atoi(nStr)
			if err != nil {
				fmt.Println("Error Case:", err)
				continue
			}
			for j := 0; j <= n; j++ {
				if i-j >= 0 {
					s[i-j] = strings.ToUpper(s[i-j])
				}
			}
			s[i] = ""
			s[i+1] = ""
		}else if s[i] == "(low," && i+1 < len(s) {
			nStr := strings.TrimSuffix(s[i+1], ")")
			n, err := strconv.Atoi(nStr)
			if err != nil {
				fmt.Println("Error Case:", err)
				continue
			}
			for j := 0; j <= n; j++ {
				if i-j >= 0 {
					s[i-j] = strings.ToLower(s[i-j])
				}
			}
			s[i] = ""
			s[i+1] = ""
		}else if s[i] == "(cap," && i+1 < len(s) {
			nStr := strings.TrimSuffix(s[i+1], ")")
			n, err := strconv.Atoi(nStr)
			if err != nil {
				fmt.Println("Error Case:", err)
				continue
			}
			for j := 0; j <= n; j++ {
				if i-j >= 0 {
					s[i-j] = Capitalize(s[i-j])
				}
			}
			s[i] = ""
			s[i+1] = ""
		}else if s[i] == "(hex," && i+1 < len(s) {
			nStr := strings.TrimSuffix(s[i+1], ")")
			n, err := strconv.Atoi(nStr)
			if err != nil {
				fmt.Println("Error Case:", err)
				continue
			}
			for j := 0; j <= n; j++ {
				if i-j >= 0 {
					n,err := strconv.ParseInt(string(s[i-j]),16,64)
					if err != nil {
						fmt.Println("Error Case:",err)
					}else{
						s[i-j] = strconv.Itoa(int(n))
					}
				}
			}
			s[i] = ""
			s[i+1] = ""
		}else if s[i] == "(bin," && i+1 < len(s) {
			nStr := strings.TrimSuffix(s[i+1], ")")
			n, err := strconv.Atoi(nStr)
			if err != nil {
				fmt.Println("Error Case:", err)
				continue
			}
			for j := 0; j <= n; j++ {
				if i-j >= 0 {
					n,err := strconv.ParseInt(string(s[i-j]),2,64)
					if err != nil {
						fmt.Println("Error Case:",err)
					}else{
						s[i-j] = strconv.Itoa(int(n))
					}
				}
			}
			s[i] = ""
			s[i+1] = ""
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