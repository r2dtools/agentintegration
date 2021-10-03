package agentintegration

type ServerMonitorTimeLineRequestData struct {
	Category, SubCategory string
	FromTime, ToTime      int
}

type ServerMonitorTimeLinePoint struct {
	Value map[string]string
	Time  string
}

type ServerMonitorTimeLineResponseData struct {
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
