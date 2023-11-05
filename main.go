package main

import (
	"devden/commands"
	"flag"
	"fmt"
	"log"
	"os"
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
	pullTemplate := createTemplate.Bool("pull", false, "Whether to pull packages from git and save it locally or not.")

	// 'create' subcommand to create a project from template
	create := flag.NewFlagSet("create", flag.PanicOnError)
	projectName := create.String("pn", "", "The name of the project being create. Default to the template's name.")
	newLocation := create.String("nl", "", "A new location of where to create the template. Default is the location of where the command was called")

	// 'bring' subcommand to bring either the config or a template to where the CLI was called to possibly be updated
	bring := flag.NewFlagSet("bring", flag.PanicOnError)
	bringConfig := bring.Bool("config", false, "Brings the devden config to the directory the command is being called.")

	// 'update' subcommand to update either a template or global config file or update a template
	update := flag.NewFlagSet("update", flag.PanicOnError)
	updateConfig := update.String("config", "", "Updates either the template config if the command is called inside the template directory, or update the devden config file given a location.")
	updateName := update.String("name", "", "Updates the template's name to a new name.")
	updateDescription := update.String("desc", "", "The template's description.")

	// 'init' to initalize the global config file with the default or with the passed in values
	init := flag.NewFlagSet("init", flag.PanicOnError)
	allowDotFiles := init.String("allow-dot-files", "true", "Whether or not to include the .<files> in the templates. Defaults to \"true\"")

	switch os.Args[1] {
	case createTemplate.Name():
		commands.HandleCreateTemplate(createTemplate, templateName, templateDescription, pullTemplate)
		break
	case create.Name():
		commands.HandleCreate(create, projectName, newLocation)
		break
	case bring.Name():
		commands.HandleBring(bring, bringConfig)
		break
	case update.Name():
		commands.HandleUpdate(update, updateConfig, updateName, updateDescription)
		break
	case init.Name():
		commands.HandleInit(init, allowDotFiles)
	default:
		log.Fatal(fmt.Sprintf("Could not understand that command: %s", os.Args[1]))
	}
}
