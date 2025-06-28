package goreloaded

func Punc(s []string) []string {
	var result []string
	for i := 0; i < len(s); i++ {
		if Only(s[i]) && i > 0 {
			punctGroup := s[i]
			j := i + 1
			for j < len(s) && Only(s[j]) {
				punctGroup += s[j]
				s[j] = ""
				j++
			}
			if !StringFlag(s[i-1]){

				s[i-1] += punctGroup
				s[i] = ""
			}else{
				s[i-2] += punctGroup
				s[i] = ""
			}
		}
	}
	for _, word := range s {
		if word != "" {
			result = append(result, word)
		}
	}
	return result
}
