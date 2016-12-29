package compromiseAnalyser

import (
	"encoding/hex"
	"testing"
)

func TestEnv(t *testing.T) {
	data, err := CheckEvairomentVariable()
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	ris1 := hex.EncodeToString(data)
	data, err = CheckEvairomentVariable()
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	ris2 := hex.EncodeToString(data)
	if ris1 != ris2 {
		t.Fatalf("Error: same command produce different result, this mean something went wrong or this computer is compromise.")
	}
}
