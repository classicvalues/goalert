package app

import (
	"time"
)

type appConfig struct {
	ListenAddr  string
	Verbose     bool
	JSON        bool
	LogRequests bool
	APIOnly     bool

	DBMaxOpen int
	DBMaxIdle int

	MaxReqBodyBytes   int64
	MaxReqHeaderBytes int

	DisableHTTPSRedirect bool

	TwilioBaseURL string
	SlackBaseURL  string

	DBURL     string
	DBURLNext string

	JaegerEndpoint      string
	JaegerAgentEndpoint string

	StackdriverProjectID string

	TracingClusterName   string
	TracingPodNamespace  string
	TracingPodName       string
	TracingContainerName string
	TracingNodeName      string

	KubernetesCooldown time.Duration
	StatusAddr         string

	LogTraces        bool
	TraceProbability float64

	EncryptionPassphrases []string

	RegionName string

	StubNotifiers bool

	UIURL string
}
