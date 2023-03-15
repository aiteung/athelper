package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func LoadEnvFile(filepath string) {
	err := godotenv.Load(filepath)
	if err != nil {
		fmt.Printf("Environtment Load is error. error log = %v", err)
		panic("Error Loading Environment")
	}
}

func CheckIsDevelopment() bool {
	return os.Getenv("DEV") == "true"
}
