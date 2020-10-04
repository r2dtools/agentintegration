package agentintegration

// VirtualHost represents webserver vhost
type VirtualHost struct {
	FilePath,
	ServerName,
	DocRoot string
	Aliases []string
	Ssl     bool
}
