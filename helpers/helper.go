package helpers

import (
	"devden/models"
	"fmt"
	"os"
	"regexp"
	"text/tabwriter"
)

func IsStringURL(str string) bool {
	var re *regexp.Regexp = regexp.MustCompile(`((git|ssh|http(s)?)|(git@[\w\.]+))(:(//)?)([\w\.@\:/\-~]+)(\.git)(/)?`)
	return len(re.FindAllString(str, -1)) >= 1
}

func IsStringPath(str string) bool {
	var re *regexp.Regexp = regexp.MustCompile(`(?i)(?:[\w]\:|(\/|\\))((\/|\\)[a-z_\-\s0-9\.]+)`)
	return len(re.FindAllString(str, -1)) >= 1
}

func DoesFileExist(location string) bool {
	_, err := os.Stat(location)
	if err == nil {
		return true
	}

	return false
}

func PrintTable(configs []*models.TemplateConfig) {
	w := tabwriter.NewWriter(os.Stdout, 1, 2, 2, ' ', 0)
	fmt.Fprintln(w, "Id\tName\tDescription\tIsCloud")
	for _, config := range configs {
		fmt.Fprintln(w, fmt.Sprintf("%s\t%s\t%s\t%t", config.Id, config.Name, config.Description, config.IsCloud))
	}
	w.Flush()
}
