package commands

import (
	"devden/helpers"
	"devden/models"
	"flag"
	"os"
	"path/filepath"
)

func HandleList(list *flag.FlagSet, all *bool, pageNumber *int) error {
	list.Parse(os.Args[2:])
	var itemCount int = 10 * *pageNumber
	var startCount int = itemCount - 10

	gConfig := createGlobalConfig()
	var selectedConfigs []*models.TemplateConfig = []*models.TemplateConfig{}

	if len(gConfig.TemplatesLocations) < itemCount {
		itemCount = len(gConfig.TemplatesLocations)
	}

	if *all {
		itemCount = len(gConfig.TemplatesLocations)
		startCount = 0
	}

	for _, template := range gConfig.TemplatesLocations[startCount:itemCount] {
		_, templateName := filepath.Split(template)
		selectedConfigs = append(selectedConfigs, getTemplateConfig(templateName))
	}
	helpers.PrintTable(selectedConfigs)

	return nil
}
