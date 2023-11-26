package dictionary

import (
	"beta/pkg/reader"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

const path = "/dictionary/validation/errors.json"

type ValidationDictionary struct {
	jsonData map[string]string
}

func (vd *ValidationDictionary) Init() *ValidationDictionary {
	projectName := regexp.MustCompile(`^(.*` + reader.ProjectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))
	fileData, err := ioutil.ReadFile(string(rootPath) + path)

	if err != nil {
		log.Println("dont read file")
	}

	data := make(map[string]string)
	err = json.Unmarshal(fileData, &data)
	if err != nil {
		fmt.Println("error:", err)
	}

	vd.jsonData = data

	return vd
}

func (vd *ValidationDictionary) Get(key string, defaultValue string) string {
	value := vd.jsonData[key]

	if value == "" {
		return defaultValue
	}

	return value
}
