package athelper

import (
	"os"
	"strings"
)

func ReplaceTokenObs(token string) string {
	obsString := os.Getenv("OBSSTRING")
	if obsString == "" {
		obsString = "+"
	}
	return strings.ReplaceAll(token, obsString, "")
}

func AddObsToken(token string) string {
	obsString := os.Getenv("OBSSTRING")
	if obsString == "" {
		obsString = "+"
	}
	return obsString + token
}
