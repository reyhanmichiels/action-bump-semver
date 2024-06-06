package action

import (
	"log"
	"os"
	"strings"
)

func GetInput(i string) string {
	e := strings.ReplaceAll(i, " ", "_")
	e = strings.ToUpper(e)
	e = "INPUT_" + e
	return strings.TrimSpace(os.Getenv(e))
}

func SetOutput(key, value string) error {
	filepath := os.Getenv("GITHUB_OUTPUT")
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Println("failed close file:", err.Error())
		}
	}()

	if _, err := f.WriteString("\n" + key + "=" + value); err != nil {
		return err
	}

	return nil
}
