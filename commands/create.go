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

func HandleCreate(create *flag.FlagSet, pn *string, nl *string) {
	if len(os.Args) < 3 {
		fmt.Println("You need to provide a template name, Run [devden list]")
		os.Exit(1)
	}
	destination, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	var templateName string = os.Args[2]
	create.Parse(os.Args[3:])
	if *nl != "" {
		destination = *nl
	}
	// Get path of the executable to create a directory there
	execPath, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("You need to set a default home directory first. Could not get your home directory.")
		os.Exit(1)
	}

	// Make sure the template exists
	var templatesDir string = filepath.Join(execPath, ".devden", "templates", templateName)
	if !helpers.DoesTemplateExist(templatesDir) {
		fmt.Println("That is not a valid template, Run [devden list]")
		os.Exit(1)
	}
	// load the template config
	var templateConfig *models.TemplateConfig = helpers.ReadJsonFile[*models.TemplateConfig](filepath.Join(templatesDir, "config.json"))

	if templateConfig.IsCloud {
		err := handleCreateFromCloudTemplate(destination, *pn, templateConfig)
		if err != nil {
			log.Fatalf("Could not clone your cloud template because of [Error = %s]", err.Error())
		}
	} else {
		err := handleCreateFromLocalTemplate(templatesDir, destination, *pn, templateConfig)
		if err != nil {
			log.Fatalf("Could not unzip your template because of [Error = %s]", err.Error())
		}
	}
}

func handleCreateFromLocalTemplate(baseLocation string, destination string, newProjectName string, config *models.TemplateConfig) error {
	if newProjectName == "" {
		newProjectName = config.Name
	}
	err := helpers.Unzip(filepath.Join(baseLocation, config.Name+".zip"), filepath.Join(destination, newProjectName))
	if err != nil {
		return err
	}
	return nil
}

func handleCreateFromCloudTemplate(destination string, newProjectName string, config *models.TemplateConfig) error {
	if newProjectName == "" {
		newProjectName = config.Name
	}
	var filePath string = filepath.Join(destination, newProjectName)
	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		return err
	}
	err := helpers.PullTemplate(config.URI, filePath)
	if err != nil {
		return err
	}
	return nil
}
