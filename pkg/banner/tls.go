package banner

import (
	"crypto/tls"
	"net"
)

func upgradeToTLS(conn net.Conn, host string) (*tls.Conn, error) {
	tlsConf := &tls.Config{
		ServerName:         host,
		InsecureSkipVerify: true,
	}

	tlsConn := tls.Client(conn, tlsConf)

	if err := tlsConn.Handshake(); err != nil {
		return nil, err
	}

	return tlsConn, nil
}

func IsTLSLikely(port int) bool {
	switch port {
	case 443, 8443, 9443, 993, 995, 465:
		return true
	default:
		return false
	}
}
