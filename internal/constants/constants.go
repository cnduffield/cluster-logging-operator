package constants

const (
	SingletonName = "instance"
	OpenshiftNS   = "openshift-logging"
	// global proxy / trusted ca bundle consts
	BearerTokenFileKey         = "token"
	ProxyName                  = "cluster"
	SharedKey                  = "shared_key"
	Passphrase                 = "passphrase"
	TrustedCABundleKey         = "ca-bundle.crt"
	SaslOverSSL                = "sasl_over_ssl"
	AWSSecretAccessKey         = "aws_secret_access_key" //nolint:gosec
	AWSAccessKeyID             = "aws_access_key_id"
	ClientCertKey              = "tls.crt"
	ClientPrivateKey           = "tls.key"
	ClientUsername             = "username"
	ClientPassword             = "password"
	InjectTrustedCABundleLabel = "config.openshift.io/inject-trusted-cabundle"
	TrustedCABundleMountFile   = "tls-ca-bundle.pem"
	TrustedCABundleMountDir    = "/etc/pki/ca-trust/extracted/pem/"
	TrustedCABundleHashName    = "logging.openshift.io/hash"
	SecretHashPrefix           = "logging.openshift.io/"
	KibanaTrustedCAName        = "kibana-trusted-ca-bundle"
	// internal elasticsearch FQDN to prevent to connect to the global proxy
	ElasticsearchFQDN          = "elasticsearch.openshift-logging.svc"
	ElasticsearchName          = "elasticsearch"
	ElasticsearchPort          = "9200"
	FluentdName                = "fluentd"
	KibanaName                 = "kibana"
	KibanaProxyName            = "kibana-proxy"
	CuratorName                = "curator"
	LogfilesmetricexporterName = "logfilesmetricexporter"
	LogStoreURL                = "https://" + ElasticsearchFQDN + ":" + ElasticsearchPort
	MasterCASecretName         = "master-certs"
	CollectorSecretName        = "collector"
	// Disable gosec linter, complains "possible hard-coded secret"
	CollectorSecretsDir     = "/var/run/ocp-collector/secrets" //nolint:gosec
	KibanaSessionSecretName = "kibana-session-secret"          //nolint:gosec

	CollectorName             = "collector"
	CollectorMetricSecretName = "collector-metrics"
	CollectorMonitorJobLabel  = "monitor-collector"
	CollectorTrustedCAName    = "collector-trusted-ca-bundle"

	LegacySecureforward = "_LEGACY_SECUREFORWARD"
	LegacySyslog        = "_LEGACY_SYSLOG"

	FluentdImageEnvVar            = "FLUENTD_IMAGE"
	LogfilesmetricImageEnvVar     = "LOGFILEMETRICEXPORTER_IMAGE"
	CertEventName                 = "cluster-logging-certs-generate"
	ClusterInfrastructureInstance = "cluster"

	//Google Cloud Logging Put all here just for testing
	Gdata_dir = "/data/vector"
	//[api]
	GAPIenabled    = true
	GAPIaddress    = "127.0.0.1:8686"
	GAPIplayground = false

	//[sources.demo_logs]
	GSourcestype   = "demo_logs"
	GSourcesformat = "shuffle"
	GSourceslines  = "[" + "Line 1" + "," + "Line 2" + "]"

	//[sources.internal_logs]
	GSourcesInternaltype     = "internal_logs"
	GSourcesInternalhost_key = "host"

	//[sinks.gcp_stackdriver_logs]
	GSyncsLogstype             = "gcp_stackdriver_logs"
	GSyncsLogsinputs           = "[" + "demo_logs" + "," + "internal_logs" + "]"
	GSyncsLogscredentials_path = "/var/run/secrets/google/credentials.json"
	GSyncsLogslog_id           = "vector-logs"
	GSyncsLogsproject_id       = "prj-caas-gcos-p-ac7a"

	//[sinks.gcp_stackdriver_logs.resource]
	GSyncsLogsResourcetype = "k8s_pod"
)

var ReconcileForGlobalProxyList = []string{CollectorTrustedCAName}
