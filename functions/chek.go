package goreloaded

func Check(s []string)bool{
	for i:=0 ; i < len(s) ; i++{
		if s[i]== "(up)" || s[i]== "(low)"  || s[i]== "(cap)"  || s[i]== "(hex)" || s[i]== "(bin)" {
			return true
		}
	}
	return false
}