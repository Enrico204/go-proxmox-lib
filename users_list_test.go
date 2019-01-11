package gplib

import "testing"

func TestProxmoximpl_GetUserList(t *testing.T) {
	_, err := c.GetUserList()
	if err != nil {
		t.Fatal(err)
	}
}