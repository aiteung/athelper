package athelper

import (
	"github.com/goccy/go-json"
	"math/rand"
	"strings"
)

func ReplaceTokenObs(obfuscated string, salt string) (res string) {
	if salt == "" {
		salt = "+"
	}
	res = strings.ReplaceAll(obfuscated, salt, "")
	return
}

func AddObsToken(token string, salt string) (res string) {
	if salt == "" {
		salt = "+"
	}
	pos := rand.Intn(len(token))
	res = strings.Join([]string{token[:pos], salt, token[pos:]}, "")
	return
}

func ConvertWithJSON[req any, res any](data *req, result *res) (err error) {
	datamarshal, err := json.Marshal(data)
	if err != nil {
		return
	}
	err = json.Unmarshal(datamarshal, result)
	return
}
