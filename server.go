package agentintegration

type ServerData struct {
	HostName        string
	Os              string
	Platform        string
	PlatformFamily  string
	PlatformVersion string
	AgentVersion    string
	KernelVersion   string
	KernelArch      string
	Virtualization  string
	Uptime          uint64
	BootTime        uint64
	Settings        map[string]string
}
