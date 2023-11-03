package helpers

import "os/exec"

func PullTemplate(location string, templatesDir string) error {
	cmdStruct := exec.Command("git", "clone", location)
	cmdStruct.Dir = templatesDir
	_, err := cmdStruct.Output()
	if err != nil {
		return err
	}

	return nil
}
