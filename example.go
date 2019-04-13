package main

import (
	"github.com/ZacJoffe/clipboard/xclip"
	"github.com/ZacJoffe/clipboard/xsel"

	"fmt"
	"log"
)

func main() {
	err := xclip.Write("Writing to the clipboard using xclip!")
	if err != nil {
		log.Fatal(err)
	}

	out, err := xclip.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)

	err = xsel.Write("Writing to the clipboard using xsel!")
	if err != nil {
		log.Fatal(err)
	}

	out, err = xsel.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
}
