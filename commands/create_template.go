package commands

import (
	"devden/helpers"
	"devden/models"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func HandleCreateTemplate(createTemplate *flag.FlagSet, name *string, desc *string) {
	if len(os.Args) < 3 {
		fmt.Println("Please provide the location of the template you want to create.\nie C:/some/path/to/template/directory or https://some-url.com/someproject/project.git")
		os.Exit(1)
	}
	// Parses the arguments
	createTemplate.Parse(os.Args[3:])

	// Gets the absolute path of the path passed in
	templatePath, err := filepath.Abs(os.Args[2])
	if err != nil {
		fmt.Printf("Couldn't resolve the given path: %s\n", os.Args[2])
	}

	// Gets the name at the end (ie path/tempalte/template-name) will return template-name
	_, templatePossibleName := filepath.Split(templatePath)

	// If there is a name provided overwrite the template name
	if *name != "" {
		templatePossibleName = *name
	}

	if templatePossibleName == "" {
		fmt.Println("You must provide a template name for this template, run [dendev create-template <location> -name <template-name>]")
		os.Exit(1)
	}

	// Get path of the executable to create a directory there
	execPath, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("You need to set a default home directory first.")
		os.Exit(1)
	}

	// Generate a path and create it if it doesnt exist
	var templatesDir string = filepath.Join(execPath, ".devden", "templates", templatePossibleName)
	err = os.MkdirAll(templatesDir, os.ModePerm)

	if err != nil {
		fmt.Println("Dont have permission to create the required directories.")
		os.Exit(1)
	}

	// Zip the file
	helpers.Zip(templatePath, templatesDir, templatePossibleName)

	// Write an object with the project configuration in it.
	templateConfig := models.TemplateConfig{
		Id:                   helpers.GenerateId(8),
		Name:                 templatePossibleName,
		Description:          *desc,
		URI:                  templatePath,
		IsCloud:              false,
		PreCreationCommands:  []string{},
		PostCreationCommands: []string{},
	}

	helpers.WriteJsonFile[*models.TemplateConfig](filepath.Join(templatesDir, "config.json"), &templateConfig)
}
