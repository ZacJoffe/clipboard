package main

import (
	"log"
	"os/exec"
)

func checkCommand(command string) error {
	_, err := exec.LookPath("xclip")
	if err != nil {
		return err
	}

	return nil
}

func Write(text string) error {
	err := checkCommand("xclip")
	if err != nil {
		return err
	}

	cmd := exec.Command("xclip", "-selection", "clipboard")

	in, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	in.Write([]byte(text))
	if err != nil {
		return err
	}

	err = in.Close()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := Write("dawdawdwafajfghkadsjklhfgasldjkghlasjkdh")
	if err != nil {
		log.Fatal(err)
	}
}
