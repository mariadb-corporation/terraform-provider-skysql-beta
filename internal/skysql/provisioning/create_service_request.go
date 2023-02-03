package provisioning

type CreateServiceRequest struct {
	Name               string          `json:"name"`
	ProjectID          string          `json:"project_id"`
	ServiceType        string          `json:"service_type"`
	Provider           string          `json:"provider"`
	Region             string          `json:"region"`
	Version            string          `json:"version"`
	Nodes              uint            `json:"nodes"`
	Architecture       string          `json:"architecture"`
	Size               string          `json:"size"`
	Topology           string          `json:"topology"`
	Storage            uint            `json:"storage"`
	VolumeIOPS         uint            `json:"volume_iops"`
	SSLEnabled         bool            `json:"ssl_enabled"`
	NoSQLEnabled       bool            `json:"nosql_enabled"`
	VolumeType         string          `json:"volume_type"`
	AllowedAccounts    []string        `json:"endpoint_allowed_accounts,omitempty"`
	Mechanism          string          `json:"endpoint_mechanism,omitempty"`
	ReplicationEnabled bool            `json:"replication_enabled,omitempty"`
	PrimaryHost        string          `json:"primary_host,omitempty"`
	AllowList          []AllowListItem `json:"allow_list,omitempty"`
}
