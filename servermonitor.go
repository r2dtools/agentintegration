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
