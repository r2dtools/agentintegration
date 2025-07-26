package agentintegration

type ServerData struct {
	HostName,
	Os,
	Platform,
	PlatformFamily,
	PlatformVersion,
	AgentVersion,
	KernelVersion,
	KernelArch,
	Virtualization string
	Uptime, BootTime uint64
}
