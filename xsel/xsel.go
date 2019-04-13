package xsel

import (
	"os/exec"
)

// checkCommand looks sees is a given command exists on the user's system, if the returned error is nil then it exists
func checkCommand(command string) error {
	_, err := exec.LookPath("xclip")
	if err != nil {
		return err
	}

	return nil
}

func Read() (string, error) {
	err := checkCommand("xsel")
	if err != nil {
		return "", err
	}

	cmd := exec.Command("xsel", "--output", "--clipboard")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}

func Write(text string) error {
	err := checkCommand("xsel")
	if err != nil {
		return err
	}

	cmd := exec.Command("xsel", "--input", "--clipboard")

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

	return cmd.Wait()
}
