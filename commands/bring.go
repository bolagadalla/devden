package commands

import (
	"devden/helpers"
	"devden/models"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func HandleBring(bring *flag.FlagSet, config *bool) error {
	if len(os.Args) >= 4 {
		bring.Parse(os.Args[3:])
	} else {
		bring.Parse(os.Args[2:])
	}

	destination, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	if *config && len(os.Args) < 4 {
		gConfig := createGlobalConfig()
		helpers.WriteJsonFile[*models.GlobalConfig](filepath.Join(destination, "global-config.json"), gConfig)
		return nil
	}

	var templateName string = os.Args[2]
	// load the template config
	var templateConfig *models.TemplateConfig = getTemplateConfig(templateName)

	if templateConfig.IsCloud {
		err := handleCreateFromCloudTemplate(destination, "", templateConfig)
		if err != nil {
			log.Fatalf("Could not clone your cloud template because of [Error = %s]", err.Error())
		}
	} else {
		err := handleCreateFromLocalTemplate(destination, "", templateConfig)
		if err != nil {
			log.Fatalf("Could not unzip your template because of [Error = %s]", err.Error())
		}
	}

	if *config {
		helpers.WriteJsonFile[*models.TemplateConfig](filepath.Join(destination, "config.json"), templateConfig)
	}
	return nil
}
