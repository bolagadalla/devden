package commands

import (
	"devden/helpers"
	"devden/models"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func HandleCreateTemplate(createTemplate *flag.FlagSet, name *string, desc *string, pull *bool) {
	if len(os.Args) < 3 {
		fmt.Println("Please provide the location of the template you want to create.\nie C:/some/path/to/template/directory or . (for current directory) or https://some-url.com/someproject/project.git")
		os.Exit(1)
	}
	// Parses the arguments
	createTemplate.Parse(os.Args[3:])

	// Gets the absolute path of the path passed in
	var location string = os.Args[2]

	if helpers.IsStringURL(location) {
		err := handleCloudTemplates(name, desc, pull, location)
		if err != nil {
			fmt.Printf("Could not save your cloud template because of [Error = %s]", err.Error())
			os.Exit(1)
		}
	} else {
		err := handleLocalTemplate(name, desc, location)
		if err != nil {
			fmt.Printf("Could not save your local template because of [Error = %s]", err.Error())
			os.Exit(1)
		}
	}
}

func handleLocalTemplate(name *string, desc *string, location string) error {
	// Gets the name at the end (ie path/template/template-name) will return template-name
	templateSourcePath, err := filepath.Abs(location)
	if err != nil {
		fmt.Printf("Couldn't resolve the given path: %s\n", location)
		return err
	}

	_, templatePossibleName := filepath.Split(templateSourcePath)

	// If there is a name provided overwrite the template name
	if *name != "" {
		templatePossibleName = *name
	}

	if templatePossibleName == "" {
		fmt.Println("You must provide a template name for this template, run [dendev create-template <location> -name <template-name>]")
		return fmt.Errorf("There was no a template name provided.")
	}

	// Get path of the executable to create a directory there
	execPath, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("You need to set a default home directory first.")
		return err
	}

	// Generate a path and create it if it doesnt exist
	var templateLocation string = filepath.Join(execPath, ".devden", "templates", templatePossibleName)
	err = os.MkdirAll(templateLocation, os.ModePerm)

	if err != nil {
		fmt.Println("Dont have permission to create the required directories.")
		return err
	}

	// Zip the template directory
	helpers.Zip(templateSourcePath, templateLocation, templatePossibleName)
	log.Println("Finished Zipping the template")

	templateConfig := models.TemplateConfig{
		Id:                   helpers.GenerateId(8),
		Name:                 templatePossibleName,
		Description:          *desc,
		URI:                  templateSourcePath,
		CurrentLocation:      templateLocation,
		IsCloud:              false,
		PreCreationCommands:  []string{},
		PostCreationCommands: []string{},
	}

	// Write an object with the project configuration in it.
	helpers.WriteJsonFile[*models.TemplateConfig](filepath.Join(templateLocation, "config.json"), &templateConfig)

	// Write this into the global config
	gConfig := createGlobalConfig()
	if !slices.Contains(gConfig.TemplatesLocations, templateLocation) {
		gConfig.TemplatesLocations = append(gConfig.TemplatesLocations, templateLocation)
	}

	// Write the global config
	helpers.WriteJsonFile[**models.GlobalConfig](filepath.Join(execPath, ".devden", "templates", "global-config.json"), &gConfig)

	return nil
}

func handleCloudTemplates(name *string, desc *string, pull *bool, location string) error {
	// Get path of the executable to create a directory there
	execPath, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("You need to set a default home directory first.")
		return err
	}

	locationSplit := strings.Split(location, "/")
	gitPackageName := strings.ReplaceAll(locationSplit[len(locationSplit)-1], ".git", "")
	var templateName string = gitPackageName
	// If there is a name provided overwrite the template name
	if *name != "" {
		templateName = *name
	}

	if templateName == "" {
		fmt.Println("You must provide a template name for this template, run [dendev create-template <location> -name <template-name>]")
		return fmt.Errorf("There was no a template name provided.")
	}

	// Generate a path and create it if it doesnt exist
	var templateLocation string = filepath.Join(execPath, ".devden", "templates", templateName)
	err = os.MkdirAll(templateLocation, os.ModePerm)

	templateConfig := models.TemplateConfig{
		Id:                   helpers.GenerateId(8),
		Name:                 templateName,
		Description:          *desc,
		URI:                  location,
		CurrentLocation:      templateLocation,
		IsCloud:              true,
		PreCreationCommands:  []string{},
		PostCreationCommands: []string{},
	}

	if *pull {
		var fullPath string = filepath.Join(templateLocation, gitPackageName)
		// Workflow to run the command to git clone this package, archive it, then remove the git package
		err = helpers.PullTemplate(location, gitPackageName, templateLocation)
		if err != nil {
			return err
		}
		helpers.Zip(fullPath, templateLocation, templateName)
		log.Println("Finished Zipping the template")
		// delete the git directory
		err := os.RemoveAll(fullPath)
		if err != nil {
			return err
		}
		// update the config
		templateConfig.IsCloud = false
		templateConfig.URI = templateLocation
	}

	// Write an object with the project configuration in it.
	helpers.WriteJsonFile[*models.TemplateConfig](filepath.Join(templateLocation, "config.json"), &templateConfig)
	return nil
}
