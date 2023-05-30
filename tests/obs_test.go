package tests

import (
	"github.com/aiteung/athelper"
	"testing"
)

func TestAddObsToken(t *testing.T) {
	token := "Hello, World!"
	salt := "+"
	res := athelper.AddObsToken(token, salt)
	if res == "+Hello, World!" {
		t.Error("Predicted Obfuscator ", res)
	}
}

func TestDeobfuscateToken(t *testing.T) {
	token := "Hello, World!"
	salt := "helloworld"
	obsfuscated := athelper.AddObsToken(token, salt)
	res := athelper.ReplaceTokenObs(obsfuscated, salt)
	if res != token {
		t.Error("Result Not Same", res)
	}
}
