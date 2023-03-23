package athelper

import (
	"testing"
)

func TestAddObsToken(t *testing.T) {
	token := "Hello, World!"
	salt := "+"
	res := AddObsToken(token, salt)
	if res == "+Hello, World!" {
		t.Error("Predicted Obfuscator ", res)
	}
}

func TestDeobfuscateToken(t *testing.T) {
	token := "Hello, World!"
	salt := "helloworld"
	obsfuscated := AddObsToken(token, salt)
	res := ReplaceTokenObs(obsfuscated, salt)
	if res != token {
		t.Error("Result Not Same", res)
	}
}
