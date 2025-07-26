package agentintegration

type VirtualHost struct {
	FilePath    string
	ServerName  string
	DocRoot     string
	WebServer   string
	Aliases     []string
	Ssl         bool
	Addresses   []VirtualHostAddress
	Certificate *Certificate
}

type VirtualHostAddress struct {
	IsIpv6 bool
	Host   string
	Port   string
}

type VirtualHostConfigResponseData struct {
	Content string
}

type VirtualHostConfigRequestData struct {
	WebServer  string
	ServerName string
}

type ReloadWebServerRequestData struct {
	WebServer string
}
