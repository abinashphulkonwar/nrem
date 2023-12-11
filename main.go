package main

import (
	"flag"
	"strings"

	"golang.design/x/clipboard"
)

func main() {
	var text string
	flag.StringVar(&text, "t", "", "text")
	flag.Parse()
	println(text)
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	text = strings.ReplaceAll(text, "\n", " ")

	clipboard.Write(clipboard.FmtText, []byte(text))

}
