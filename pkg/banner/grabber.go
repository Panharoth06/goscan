package banner

import (
	"fmt"
	"net"
	"time"
)

const readTimeout = 2 * time.Second
const maxReadSize = 16384

func GrabService(host string, port int) ServiceInfo {
	info := ServiceInfo{
		Port: port,
		Host: host,
	}

	address := net.JoinHostPort(host, fmt.Sprintf("%d", port))
	conn, err := net.DialTimeout("tcp", address, 3*time.Second)
	if err != nil {
		info.Error = err
		return info
	}
	defer conn.Close()

	conn.SetDeadline(time.Now().Add(readTimeout))

	// Passive read
	initial, err := readAll(conn)
	if err == nil && len(initial) > 0 {
		info.Banner = ParseBanner(initial)
		info.Protocol = DetectProtocol(initial, port)
		return info
	}

	// Decide probe
	probe := SelectProbe(port)

	// Handle TLS if likely
	if IsTLSLikely(port) {
		tlsConn, err := upgradeToTLS(conn, host)
		if err != nil {
			info.Error = err
			return info
		}
		defer tlsConn.Close()

		response, err := sendProbe(tlsConn, probe)
		if err != nil {
			info.Error = err
			return info
		}

		info.TLS = true
		info.Banner = ParseBanner(response)
		info.Protocol = DetectProtocol(response, port)
		return info
	}

	// Plain TCP probe
	response, err := sendProbe(conn, probe)
	if err != nil {
		info.Error = err
		return info
	}

	info.Banner = ParseBanner(response)
	info.Protocol = DetectProtocol(response, port)

	return info
}
