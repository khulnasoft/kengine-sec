package report

const (
	// DateTime is an ISO timestamp of the format "2017-07-03T09:45:00.329067309Z"
	DateTime = "datetime"

	// Duration specified in seconds, e.g. "3600" means one hour
	Duration = "duration"

	// IP is a string in the format "182.43.147.201"
	IP = "ip"

	// Number as an integer or a floating point
	Number = "number"
)

const (
	KhulnasoftSystemLabelKey   = "khulnasoft.role"
	KhulnasoftSystemLabelValue = "system"

	KhulnasoftK8sAppLabelKey          = "app"
	KhulnasoftK8sConsoleAppLabelValue = "khulnasoft-console"
	KhulnasoftK8sAgentAppLabelValue   = "khulnasoft-agent"
)

const (
	PodNamespaceLabel = "io.kubernetes.pod.namespace"
	PodNameLabel      = "io.kubernetes.pod.name"
	PodIDLabel        = "io.kubernetes.pod.uid"
)
