package commands

import (
	"devden/helpers"
	"devden/models"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func HandleInit(create *flag.FlagSet, allowDotFiles *string) error {
	create.Parse(os.Args[2:])
	// Get path of the executable to create a directory there
	execPath, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("You need to set a default home directory first.")
		return err
	}

	gConfig := createGlobalConfig()
	gConfig.AllowDotFiles = *allowDotFiles == "true"

	// Write the global config
	helpers.WriteJsonFile[**models.GlobalConfig](filepath.Join(execPath, ".devden", "templates", "global-config.json"), &gConfig)
	return nil
}

func createGlobalConfig() *models.GlobalConfig {
	gConfig := getGlobalConfigIfExists()
	if gConfig == nil {
		gConfig = &models.GlobalConfig{
			AllowDotFiles:      true,
			TemplatesLocations: []string{},
		}
	}
	return gConfig
}

func getGlobalConfigIfExists() *models.GlobalConfig {
	// Get path of the executable to create a directory there
	execPath, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("You need to set a default home directory first.")
		return nil
	}
	var globalConfigLocation string = filepath.Join(execPath, ".devden", "templates", "global-config.json")

	if helpers.DoesFileExist(globalConfigLocation) {
		return helpers.ReadJsonFile[*models.GlobalConfig](globalConfigLocation)
	}
	return nil
}
