package agentintegration

type ServerMonitorStatisticsRequestData struct {
	Category, SubCategory string
	FromTime, ToTime      int
}

type ServerMonitorTimeLinePoint struct {
	Value map[string]string
	Time  string
}

type ServerMonitorStatisticsResponseData struct {
	Data map[string][]ServerMonitorTimeLinePoint
}

type ServerMonitorDiskResponseData struct {
	DiskUsageTimeLineData map[string][]ServerMonitorTimeLinePoint
	DiskIOTimeLineData    map[string][]ServerMonitorTimeLinePoint
	DiskInfo              map[string]map[string]string // mountpoint => {'total': ..., 'free': ..., 'used': ..., ...}
}

type ServerMonitorNetworkResponseData struct {
	TimeLineData   map[string][]ServerMonitorTimeLinePoint
	InterfacesInfo []map[string]string
}

type ProcessData struct {
	Name,
	User,
	Cmd string
	Pid,
	PPid int32
	Cpu       float64
	Memory    float32
	OpenFiles []string
	NetBytesRecv,
	NetBytesSent,
	NetPacketsRecv,
	NetPacketsSent,
	DiskReadBytes,
	DiskWriteBytes,
	DiskReadCount,
	DiskWriteCount uint64
}

type ServerMonitorProcessResponseData struct {
	ProcessesData []ProcessData
}
