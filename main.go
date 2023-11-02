package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Validates arguments being passed
	if len(os.Args) < 2 {
		log.Println("Expected a subcommand")
		os.Exit(1)
	}
	// 'create-template' subcommand
	createTemplate := flag.NewFlagSet("create-template", flag.PanicOnError)
	templateName := createTemplate.String("name", "", "The template name which will be used when you create from the template. It will default to the template root directory name or git project name.")
	templateDescription := createTemplate.String("desc", "", "The template's description.")

	// 'create' subcommand to create a project from template
	create := flag.NewFlagSet("create", flag.PanicOnError)
	projectName := create.String("pn", "", "The name of the project being create. Default to the template's name.")

	// 'bring' subcommand to bring either the config or a template to where the CLI was called to possibly be updated
	bring := flag.NewFlagSet("bring", flag.PanicOnError)
	bringConfig := bring.Bool("config", false, "Brings the devden config to the directory the command is being called.")

	// 'update' subcommand to update either a template or global config file or update a template
	update := flag.NewFlagSet("update", flag.PanicOnError)
	updateConfig := update.String("config", "", "Updates either the template config if the command is called inside the template directory, or update the devden config file given a location.")
	updateName := update.String("name", "", "Updates the template's name to a new name.")
	updateDescription := update.String("desc", "", "The template's description.")

	switch os.Args[1] {
	case createTemplate.Name():
		HandleCreateTemplate(createTemplate, templateName, templateDescription)
		break
	case create.Name():
		HandleCreate(create, projectName)
		break
	case bring.Name():
		HandleBring(bring, bringConfig)
		break
	case update.Name():
		HandleUpdate(update, updateConfig, updateName, updateDescription)
		break
	default:
		log.Fatal(fmt.Sprintf("Could not understand that command: %s", os.Args[1]))
	}
}

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
	Zip(templatePath, templatesDir, templatePossibleName)

	// Write an object with the project configuration in it.
	templateConfig := TemplateConfig{
		Id:                   GenerateId(8),
		Name:                 templatePossibleName,
		Description:          *desc,
		URI:                  templatePath,
		IsCloud:              false,
		PreCreationCommands:  []string{},
		PostCreationCommands: []string{},
	}

	WriteJsonFile[*TemplateConfig](filepath.Join(templatesDir, "config.json"), &templateConfig)
}

func HandleCreate(create *flag.FlagSet, pn *string) {
	create.Parse(os.Args[2:])
}

func HandleBring(bring *flag.FlagSet, config *bool) {
	bring.Parse(os.Args[2:])
}

func HandleUpdate(update *flag.FlagSet, config *string, name *string, desc *string) {
	update.Parse(os.Args[2:])
}

func HandleHelp(help *flag.FlagSet) {
	help.Parse(os.Args[2:])
}
