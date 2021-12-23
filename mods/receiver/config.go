package receiver

type Config struct {
	LogName               string
	ClientId              string
	ConsumerGroup         string
	Brokers               []string
	SaslEnabled           bool
	SaslMechanism         string
	SaslUser              string
	SaslPassword          string
	TlsEnabled            bool
	TlsInsecureSkipVerify bool
	TlsClientAuth         int
}
