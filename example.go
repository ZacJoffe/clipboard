package main

import (
	//"github.com/ZacJoffe/clipboard/xclip"
	//"github.com/ZacJoffe/clipboard/xsel"

	"./xsel"
	"fmt"
	"log"
)

func main() {
	err := xsel.Write("Writing to the clipboard!!!")
	if err != nil {
		log.Fatal(err)
	}

	out, err := xsel.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
}
