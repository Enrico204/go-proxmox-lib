package gplib

import (
	"os"
	"testing"
)

func TestProxmoximpl_GetUserInfo(t *testing.T) {
	_, err := c.GetUserInfo(os.Getenv("USER"))
	if err != nil {
		t.Fatal(err)
	}
}