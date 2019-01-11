package gplib

import "testing"

func TestProxmoximpl_GetNodeVMs(t *testing.T) {
	nodelist, err := c.GetNodeList()
	if err != nil {
		t.Fatal(err)
	}

	_, err = c.GetNodeVMs(nodelist[0].Node)
	if err != nil {
		t.Fatal(err)
	}
}