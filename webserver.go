package agentintegration

// VirtualHost represents webserver vhost
type VirtualHost struct {
	FilePath,
	ServerName,
	DocRoot,
	WebServer string
	Aliases []string
	Ssl     bool
}
