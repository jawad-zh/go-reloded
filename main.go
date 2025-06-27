package main

import (
	"fmt"
	goreloaded "goreloaded/functions"
	"os"
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
	clean = goreloaded.Quote(clean)
	// clean = goreloaded.Quote(clean)
	clean = goreloaded.Rchar(clean)
	// fmt.Print(strings.Join(clean, " "))
	fmt.Print(clean)
}
