package gplib

import "testing"

func TestProxmoximpl_GetNodeList(t *testing.T) {
	_, err := c.GetNodeList()
	if err != nil {
		t.Fatal(err)
	}
}