package gplib

type Capabilities struct {
	VMs     map[string]int `json:"vms"`
	DC      map[string]int `json:"dc"`
	Nodes   map[string]int `json:"nodes"`
	Access  map[string]int `json:"access"`
	Storage map[string]int `json:"storage"`
}

func (c *Capabilities) Has(capability string) bool {
	_, vms := c.VMs[capability]
	_, dc := c.DC[capability]
	_, nodes := c.Nodes[capability]
	_, access := c.Access[capability]
	_, storage := c.Storage[capability]
	return vms || dc || nodes || access || storage
}
