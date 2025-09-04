package main

import (
	"log"

	"github.com/mpr1255/mhtml-to-html/cmd"
)

func main() {
	var c cmd.MHTMLToHTML
	if e := c.Run(); e != nil {
		log.Fatal(e)
	}
}
