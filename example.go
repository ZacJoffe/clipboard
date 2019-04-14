package main

import (
	"github.com/ZacJoffe/clipboard/xsel"
	"github.com/zacjoffe/clipboard/xclip"

	"fmt"
	"log"
	"os"
)

func main() {
	err := xclip.WriteText("Writing to the clipboard using xclip!")
	if err != nil {
		log.Fatal(err)
	}

	out, err := xclip.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)

	err = xsel.WriteText("Writing to the clipboard using xsel!")
	if err != nil {
		log.Fatal(err)
	}

	out, err = xsel.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)

	image, err := os.Open("/tmp/screenshot.png")
	if err != nil {
		log.Fatal(err)
	}

	err = xclip.WriteImage(image)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Wrote image to clipboard using xclip!")
}
