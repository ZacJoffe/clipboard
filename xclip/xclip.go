package xclip

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

// Read reads whatever is in the clipboard, and returns it as a string
func Read() (string, error) {
	// check to see if xclip exists
	err := checkCommand("xclip")
	if err != nil {
		return "", err
	}

	// run the command "xclip -out -selection clipboard" to get what's in the clipboard
	cmd := exec.Command("xclip", "-out", "-selection", "clipboard")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// convert the output to a string, and return it
	return string(out), nil
}

// Write writes a given string to the clipboard
func Write(text string) error {
	// check to see if xclip exists
	err := checkCommand("xclip")
	if err != nil {
		return err
	}

	// create the command "xclip -in -selection clipboard"
	cmd := exec.Command("xclip", "-in", "-selection", "clipboard")

	// create new stdin pipe
	in, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	// start the command
	err = cmd.Start()
	if err != nil {
		return err
	}

	// write the given text parameter into the writer, which will act as the stdin
	in.Write([]byte(text))
	if err != nil {
		return err
	}

	// explicitly close the writer
	err = in.Close()
	if err != nil {
		return err
	}

	// wait for command to exit and stdin copying - will return nil if successful
	return cmd.Wait()
}
