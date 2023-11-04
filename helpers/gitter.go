package helpers

import "os/exec"

func PullTemplate(gitURL string, cloneDir string, templatesDir string) error {
	cmdStruct := exec.Command("git", "clone", gitURL, cloneDir)
	cmdStruct.Dir = templatesDir
	_, err := cmdStruct.Output()
	if err != nil {
		return err
	}

	return nil
}
