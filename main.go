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
	err := clipboard.Init()
	if err != nil {

		panic(err)

	}

	if text == "" {
		text = string(clipboard.Read(clipboard.FmtText))
	}

	text = strings.ReplaceAll(text, "\n", " ")
	text = strings.ReplaceAll(text, "\r", " ")
	text = strings.ReplaceAll(text, "\t", " ")
	text = strings.ReplaceAll(text, "  ", " ")
	if text[0] == ' ' {
		text = text[1:]
	}
	clipboard.Write(clipboard.FmtText, []byte(text))
	println("Text copied to clip board")

}
