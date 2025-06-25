package main

import (
	"fmt"
	"os"
	"strings"

	goreloaded "goreloaded/functions"
)

func main() {
	content, err := os.ReadFile("input/sampl.txt")
	if err != nil {
		fmt.Print("Error:", err)
		return
	}
	txt := string(content)
	clean := goreloaded.Cleaned(txt)
	clean = goreloaded.Convert(clean)
	clean = goreloaded.Punc(clean)
	clean = goreloaded.ConvN(clean)
	clean = goreloaded.AtoAn(clean)
	 hasQuote := false
	for i := 0; i < len(clean); i++ {
    if goreloaded.IsQ(clean[i]) {
        hasQuote = true
        break
    }
}
	if hasQuote {
    fmt.Print(strings.Join(clean, ""))
	}else {
    fmt.Print(strings.Join(clean, " "))
}

	
}
