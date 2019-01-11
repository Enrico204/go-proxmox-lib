package gplib

import "testing"

func TestProxmoximpl_GetPoolList(t *testing.T) {
	_, err := c.GetPoolList()
	if err != nil {
		t.Fatal(err)
	}
}