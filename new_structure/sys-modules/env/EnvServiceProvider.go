package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"project-root/app"
)

func (s *EnvServiceProvider) Register() {

	env := os.Getenv("GO_ENV")
	fileName := ".env"
	if env == "test" {
		fileName = ".env.test"
	}

	// set PROJECT_ROOT while run test
	projectRoot := os.Getenv("PROJECT_ROOT")
	if projectRoot == "" {
		// default
		if err := godotenv.Load(".env"); err != nil {
			log.Fatalf("Error loading .env file")
			panic("Error loading .env file")
		}
	} else {
		// while testing - find from root project
		envPath := filepath.Join(projectRoot, fileName)
		if err := godotenv.Load(envPath); err != nil {
			log.Fatalf("Error loading %s file from %s: %v", fileName, projectRoot, err)
		}
	}

	// ----------------
	//env := os.Getenv("GO_ENV")
	//if env == "test" {
	//	if err := godotenv.Load(".env.test"); err != nil {
	//		log.Fatalf("Error loading .env.test file")
	//	}
	//} else {
	//	if err := godotenv.Load(".env"); err != nil {
	//		log.Fatalf("Error loading .env file")
	//		panic("Error loading .env file")
	//	}
	//}
	// -------------
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//	panic("Error loading .env file")
	//}
}

func (s *EnvServiceProvider) Boot() {

}

type EnvServiceProvider struct{}

var _ app.ServiceProviderInterface = &EnvServiceProvider{}
