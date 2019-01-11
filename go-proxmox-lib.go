package gplib

type Proxmox interface {
	Login(username string, password string) (LoginResponse, error)
	GetVersion() (VersionInfo, error)

	GetUserList() ([]UserListItem, error)
	GetUserInfo(userid string) (UserInfo, error)
	NewUser(userid string, email string, expire int, firstname string, lastname string, groups []string, password string) error
	DeleteUser(userid string) error

	GetPoolList() ([]PoolListItem, error)
	NewPool(poolid string, comment string) error
	GetPoolInfo(poolid string) (PoolInfo, error)
	DeletePool(poolid string) error

	GetNodeList() ([]NodeListItem, error)

	GetNodeVMs(node string) ([]NodeVMListItem, error)

	GetNodeNetworks(node string, nictype string) ([]NodeNetworkListItem, error)
	ReloadNetworkConfig(node string) error
	RevertNetworkConfig(node string) error
}

type proxmoximpl struct {
	serverURL           string
	ticket              string
	csrfPreventionToken string
	username            string
	insecureSkipVerify  bool
}

func New(serverURL string, insecureTLS bool) Proxmox {
	return &proxmoximpl{
		serverURL: serverURL,
		insecureSkipVerify: insecureTLS,
	}
}
