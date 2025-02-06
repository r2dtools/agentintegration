package agentintegration

type VirtualHost struct {
	FilePath,
	ServerName,
	DocRoot,
	WebServer string
	Aliases     []string
	Ssl         bool
	Addresses   []VirtualHostAddress
	Certificate *Certificate
}

type VirtualHostAddress struct {
	IsIpv6 bool
	Host,
	Port string
}

type VirtualHostConfigResponseData struct {
	Content string
}

type VirtualHostConfigRequestData struct {
	WebServer  string
	ServerName string
}
