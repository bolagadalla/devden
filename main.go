package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Validates arguments being passed
	if len(os.Args) < 2 {
		fmt.Println("Expected a subcommand. Run [devden help]")
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
		panic(fmt.Sprintf("Could not understand that command: %s", os.Args[1]))
	}
}

func HandleCreateTemplate(createTemplate *flag.FlagSet, name *string, desc *string) {
	createTemplate.Parse(os.Args[2:])
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
