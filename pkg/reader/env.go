package reader

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
)

const ProjectDirName = "beta"

type Env struct {
}

func GetEnvReader() *Env {
	env := new(Env)
	env.loadEnv()
	return env
}

func (env *Env) Get(key string) string {
	value, exist := os.LookupEnv(key)

	if exist == false {
		log.Fatalf("No .env %s value exist", key)
	}

	return value
}

func (env *Env) loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + ProjectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))
	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
