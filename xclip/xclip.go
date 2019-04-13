package xclip

import (
	"os/exec"
)

func checkCommand(command string) error {
	_, err := exec.LookPath("xclip")
	if err != nil {
		return err
	}

	return nil
}

func Read() (string, error) {
	err := checkCommand("xclip")
	if err != nil {
		return "", err
	}

	cmd := exec.Command("xclip", "-out", "-selection", "clipboard")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}

func Write(text string) error {
	err := checkCommand("xclip")
	if err != nil {
		return err
	}

	cmd := exec.Command("xclip", "-in", "-selection", "clipboard")

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
