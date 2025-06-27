package goreloaded

func Quote(s []string)[]string{
for i:=0 ; i < len(s) ; i++{
	if IsQ(s[i]) && i != len(s)-1{
		s[i] = s[i] + s[i+1]
		s[i+1] = ""
	}
}
var Result []string
for _,word:= range s{
	if word != ""{
		Result = append(Result,word)
	}
}
	return Result
}