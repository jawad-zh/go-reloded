package goreloaded

func Cleaned(input string) []string {
	var word string
	var words []string
	var punct string
	runes := []rune(input)
for i := 0; i < len(runes); i++ {
    char := string(runes[i])
		if char != " " && char != "\n" && char != "\r"{
			if IsQ(char){
				if word!= ""{
					words = append(words, word)
					word=""
				}else if punct !=""{
					words = append(words, punct)
					punct=""
				}else{
					words = append(words,char)
					char  = ""
				}
			}
			if Only(char) {
				if word != "" {
					words = append(words, word)
					word = ""
				}
				punct += char
			} else {
				if punct != "" {
					words = append(words, punct)
					punct = ""
				}
				word += char
			}
		} else {
			if word != "" {
				words = append(words, word)
				word = ""
			}
			if punct != "" {
				words = append(words, punct)
				punct = ""
			}
			if char == "\n" || char == "\r" {
				words = append(words, char)
			}
		}
	}
	if word != "" {
		words = append(words, word)
	}
	if punct != "" {
		words = append(words, punct)
	}
	return words
}
