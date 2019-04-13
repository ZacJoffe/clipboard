package main

import (
	//"github.com/ZacJoffe/clipboard/xclip"
	//"./xclip"
	"./xsel"
	"fmt"
	"log"
)

func main() {
	/*
		err := xclip.Write("Writing to the clipboard!!!")
		if err != nil {
			log.Fatal(err)
		}
	*/

	out, err := xsel.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
}
