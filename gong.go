package gong

type Node interface {
	Information() (*NodeInfo, error)
	Status() (*NodeStatus, error)

	NewApiAdmin() ApiAdmin
}

type ApiAdmin interface {
	Add(*ApiDefinition) (*ApiObject, error)
	Retrieve(name_or_id string) (*ApiObject, error)
	List(filters ...string) ([]*ApiObject, error)
	Update(name_or_id string, definition *ApiDefinition) (*ApiObject, error)
	Replace(*ApiDefinition) (*ApiObject, error)
	Delete(name_or_id string) error
}

type NodeInfo struct {
	Hostname   string `json:"hostname"`
	LuaVersion string `json:"lua_version"`
	Plugins    struct {
		AvailableOnServer []string `json:"available_on_server"`
		EnableInCluster   []string `json:"enable_in_cluster"`
	} `json:"plugins"`
	Configuration map[string]interface{} `json:"configuration"`
	Tagline       string                 `json:"tagline"`
	Version       string                 `json:"version"`
}

type NodeStatus struct {
	Server struct {
		TotalRequests       int `json:"total_requests"`
		ConnectionsActive   int `json:"connections_active"`
		ConnectionsAccepted int `json:"connections_accepted"`
		ConnectionsHandled  int `json:"connections_handled"`
		ConnectionsReading  int `json:"connections_reading"`
		ConnectionsWriting  int `json:"connections_writing"`
		ConnectionsWaiting  int `json:"connections_waiting"`
	} `json:"server"`
	Database struct {
		Reachable bool `json:"reachable"`
	} `json:"database"`
}

type ApiDefinition struct {

	// required
	Name        string   `json:"name"`
	Hosts       []string `json:"hosts"`
	Uris        []string `json:"uris"`
	Methods     []string `json:"methods"`
	UpstreamUrl string   `json:"upstream_url"`

	// optional
	StripUri               bool `json:"strip_uri"`
	PreserveHost           bool `json:"preserve_host"`
	Retries                int  `json:"retries"`
	UpstreamConnectTimeout int  `json:"upstream_connect_timeout"`
	UpstreamSendTimeout    int  `json:"upstream_send_timeout"`
	UpstreamReadTimeout    int  `json:"upstream_read_timeout"`
	HttpsOnly              bool `json:"https_only"`
	HttpIfTerminated       bool `json:"http_if_terminated"`
}

// NewApiDefinition generate ApiDefinition with some default values
func NewApiDefinition(
	name string, upstream string, uris ...string) *ApiDefinition {

	return &ApiDefinition{

		// required
		Name:        name,
		UpstreamUrl: upstream,

		// semi-optional
		Uris: uris,

		// optional with default value
		StripUri:               true,
		PreserveHost:           false,
		Retries:                5,
		UpstreamConnectTimeout: 60000,
		UpstreamSendTimeout:    60000,
		UpstreamReadTimeout:    60000,
		HttpsOnly:              false,
		HttpIfTerminated:       true,
	}
}

type ApiObject struct {
	CreatedAt              int64    `json:"created_at"`
	Hosts                  []string `json:"hosts"`
	HttpIfTerminated       bool     `json:"http_if_terminated"`
	HttpsOnly              bool     `json:"https_only"`
	Id                     string   `json:"id"`
	Name                   string   `json:"name"`
	PreserveHost           bool     `json:"preserve_host"`
	Retries                int      `json:"retries"`
	StripUri               bool     `json:"strip_uri"`
	UpstreamConnectTimeout int      `json:"upstream_connect_timeout"`
	UpstreamReadTimeout    int      `json:"upstream_read_timeout"`
	UpstreamSendTimeout    int      `json:"upstream_send_timeout"`
	UpstreamUrl            string   `json:"upstream_url"`
}
