package agentintegration

type ServerMonitorTimeLineRequestData struct {
	Category, SubCategory string
	FromTime, ToTime      int
}

type ServerMonitorTimeLinePoint struct {
	Value float32
	Time  int
}

type ServerMonitorTimeLineResponseData struct {
	Data map[string][]ServerMonitorTimeLinePoint
}
