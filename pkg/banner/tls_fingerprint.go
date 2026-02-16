package banner

import (
	"crypto/tls"
	"time"
)

func populateTLSInfo(info *ServiceInfo, state *tls.ConnectionState) {
	if state == nil {
		return
	}

	// TLS version
	info.TLSVersion = tlsVersionToString(state.Version)

	// Cipher suite
	info.CipherSuite = tls.CipherSuiteName(state.CipherSuite)

	// Certificate
	if len(state.PeerCertificates) > 0 {
		cert := state.PeerCertificates[0]

		info.CertCN = cert.Subject.CommonName
		info.CertIssuer = cert.Issuer.CommonName

		if len(cert.Subject.Organization) > 0 {
			info.CertOrg = cert.Subject.Organization[0]
		}

		info.CertExpiry = cert.NotAfter.Format(time.RFC3339)
	}
}

func tlsVersionToString(version uint16) string {
	switch version {
	case tls.VersionTLS10:
		return "TLS1.0"
	case tls.VersionTLS11:
		return "TLS1.1"
	case tls.VersionTLS12:
		return "TLS1.2"
	case tls.VersionTLS13:
		return "TLS1.3"
	default:
		return "Unknown"
	}
}
