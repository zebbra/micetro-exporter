package micetro

type DHCPScope struct {
	Ref           string `json:"ref"`
	Name          string `json:"name"`
	RangeRef      string `json:"rangeRef"`
	DhcpServerRef string `json:"dhcpServerRef"`
	Superscope    string `json:"superscope"`
	Description   string `json:"description"`
	Available     int    `json:"available"`
	Enabled       bool   `json:"enabled"`
}

type DHCPScopeList struct {
	Result struct {
		DHCPScopes []DHCPScope `json:"dhcpScopes"`
	} `json:"result"`
	TotalResults int `json:"totalResults"`
}

type DHCPServer struct {
	Ref              string                 `json:"ref"`
	Name             string                 `json:"name"`
	Proxy            string                 `json:"proxy"`
	Address          string                 `json:"address"`
	ResolvedAddress  string                 `json:"resolvedAddress"`
	Username         string                 `json:"username"`
	Password         string                 `json:"password"`
	EnablePassword   string                 `json:"enablePassword"`
	Type             string                 `json:"type"`
	HaMode           string                 `json:"haMode"`
	State            string                 `json:"state"`
	Security         string                 `json:"security"`
	CustomProperties map[string]interface{} `json:"customProperties"`
	Enabled          bool                   `json:"enabled"`
}

type DHCPServerList struct {
	Result struct {
		DHCPServers []DHCPServer `json:"dhcpServers"`
	} `json:"result"`
	TotalResults int `json:"totalResults"`
}

type Range struct {
	Ref               string                 `json:"ref"`
	Name              string                 `json:"name"`
	From              string                 `json:"from"`
	To                string                 `json:"to"`
	ChildRanges       []interface{}          `json:"childRanges"`
	DhcpScopes        []interface{}          `json:"dhcpScopes"`
	Subnet            bool                   `json:"subnet"`
	Locked            bool                   `json:"locked"`
	AutoAssign        bool                   `json:"autoAssign"`
	HasSchedule       bool                   `json:"hasSchedule"`
	HasMonitor        bool                   `json:"hasMonitor"`
	CustomProperties  map[string]interface{} `json:"customProperties"`
	InheritAccess     bool                   `json:"inheritAccess"`
	IsContainer       bool                   `json:"isContainer"`
	HasRogueAddresses bool                   `json:"hasRogueAddresses"`
}

type RangeList struct {
	Result struct {
		Ranges []Range `json:"ranges"`
	} `json:"result"`
	TotalResults int `json:"totalResults"`
}
