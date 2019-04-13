package main

import (
	"log"
	"os/exec"
)

func main() {
	_, err := exec.LookPath("xclip")
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("xclip", "-selection", "clipboard")

	in, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	in.Write([]byte("dqasda"))
	if err != nil {
		log.Fatal(err)
	}

	err = in.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
