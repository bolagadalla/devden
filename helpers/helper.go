package helpers

import (
	"os"
	"regexp"
)

func IsStringURL(str string) bool {
	var re *regexp.Regexp = regexp.MustCompile(`((git|ssh|http(s)?)|(git@[\w\.]+))(:(//)?)([\w\.@\:/\-~]+)(\.git)(/)?`)
	return len(re.FindAllString(str, -1)) >= 1
}

func IsStringPath(str string) bool {
	var re *regexp.Regexp = regexp.MustCompile(`(?i)(?:[\w]\:|(\/|\\))((\/|\\)[a-z_\-\s0-9\.]+)`)
	return len(re.FindAllString(str, -1)) >= 1
}

func DoesTemplateExist(location string) bool {
	_, err := os.Stat(location)
	if err == nil {
		return true
	}

	return false
}
