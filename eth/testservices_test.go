package eth

import "testing"

func handleFatalError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
