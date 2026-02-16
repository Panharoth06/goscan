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

	// if TLS port → handshake immediately
	if IsTLSLikely(port) {

		tlsConn, state, err := upgradeToTLS(conn, host)
		if err != nil {
			info.Error = err
			return info
		}
		defer tlsConn.Close()

		info.TLS = true
		populateTLSInfo(&info, state)

		probe := SelectProbe(port)
		response, err := sendProbe(tlsConn, probe)
		if err == nil && len(response) > 0 {
			info.Banner = ParseBanner(response)
			info.Protocol = DetectProtocol(response, port)
		}

		return info
	}

	// Non-TLS → try passive read
	initial, _ := readAll(conn)
	if len(initial) > 0 {
		info.Banner = ParseBanner(initial)
		info.Protocol = DetectProtocol(initial, port)
		return info
	}

	// Send probe
	probe := SelectProbe(port)
	response, err := sendProbe(conn, probe)
	if err != nil {
		info.Error = err
		return info
	}

	info.Banner = ParseBanner(response)
	info.Protocol = DetectProtocol(response, port)

	return info
}
