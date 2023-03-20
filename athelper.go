package athelper

import (
	"github.com/goccy/go-json"
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

func ConvertWithJSON[req any, res any](data *req, result *res) (err error) {
	datamarshal, err := json.Marshal(data)
	if err != nil {
		return
	}
	err = json.Unmarshal(datamarshal, result)
	return
}
