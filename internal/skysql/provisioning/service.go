package provisioning

type Service struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Region       string `json:"region"`
	Provider     string `json:"provider"`
	Tier         string `json:"tier"`
	Topology     string `json:"topology"`
	Version      string `json:"version"`
	Architecture string `json:"architecture"`
	Size         string `json:"size"`
	Nodes        int    `json:"nodes"`
	SslEnabled   bool   `json:"ssl_enabled"`
	NosqlEnabled bool   `json:"nosql_enabled"`
	FQDN         string `json:"fqdn"`
	Status       string `json:"status"`
	CreatedOn    int    `json:"created_on"`
	UpdatedOn    int    `json:"updated_on"`
	CreatedBy    string `json:"created_by"`
	UpdatedBy    string `json:"updated_by"`
	Endpoints    []struct {
		Name  string `json:"name"`
		Ports []struct {
			Name    string `json:"name"`
			Port    int    `json:"port"`
			Purpose string `json:"purpose"`
		} `json:"ports"`
	} `json:"endpoints"`
	StorageVolume struct {
		Size       int    `json:"size"`
		VolumeType string `json:"volume_type"`
		IOPS       int    `json:"iops"`
	} `json:"storage_volume"`
	OutboundIps        []string `json:"outbound_ips"`
	IsActive           bool     `json:"is_active"`
	ServiceType        string   `json:"service_type"`
	ReplicationEnabled bool     `json:"replication_enabled"`
	PrimaryHost        string   `json:"primary_host"`
}