package gplib

import (
	"testing"
)

func TestProxmoximpl_GetVersion(t *testing.T) {
	_, err := c.GetVersion()
	if err != nil {
		t.Fatal(err)
	}
}