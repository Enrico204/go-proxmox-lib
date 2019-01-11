package gplib

import "testing"

func TestProxmoximpl_GetNodeNetworks(t *testing.T) {
	nodelist, err := c.GetNodeList()
	if err != nil {
		t.Fatal(err)
	}

	_, err = c.GetNodeNetworks(nodelist[0].Node, "")
	if err != nil {
		t.Fatal(err)
	}
}