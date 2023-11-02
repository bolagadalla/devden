package helpers

import (
	"encoding/json"
	"log"
	"os"
)

func ReadJsonFile[R comparable](source string) (value R) {
	fileBytes, err := os.ReadFile(source)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(fileBytes, &value)
	if err != nil {
		log.Fatal(err)
	}

	return value
}

func WriteJsonFile[V comparable](destination string, object V) {
	objectBytes, err := json.Marshal(object)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(destination, objectBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
