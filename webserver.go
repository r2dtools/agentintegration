package agentintegration

// VirtualHost represents webserver vhost
type VirtualHost struct {
	FilePath,
	ServerName,
	DocRoot,
	WebServer string
	Aliases   []string
	Ssl       bool
	Addresses []VirtualHostAddress
}

// VirtualHostAddress repesents webserver vhost address
type VirtualHostAddress struct {
	IsIpv6 bool
	Host,
	Port string
}
