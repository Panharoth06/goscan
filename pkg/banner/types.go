package banner

type ServiceInfo struct {
	Port        int
	Host        string
	Protocol    string
	TLS         bool
	Banner      string
	Error       error

	// TLS Fingerprinting
	TLSVersion  string
	CipherSuite string
	CertCN      string
	CertIssuer  string
	CertOrg     string
	CertExpiry  string
}

