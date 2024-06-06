package action

import (
	"os"
	"strings"
)

func GetInput(i string) string {
	e := strings.ReplaceAll(i, " ", "_")
	e = strings.ToUpper(e)
	e = "INPUT_" + e
	return strings.TrimSpace(os.Getenv(e))
}
