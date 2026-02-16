package banner

type ServiceInfo struct {
	Port     int
	Host     string
	Protocol string
	TLS      bool
	Banner   string
	Error    error
}
