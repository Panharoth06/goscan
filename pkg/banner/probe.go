package banner

import "fmt"

func SelectProbe(port int) string {
	switch port {

	case 80, 8080, 8000, 8888:
		return httpProbe("")

	case 443, 8443, 9443:
		return httpProbe("")

	case 25:
		return "EHLO scanner\r\n"

	case 21:
		return "FEAT\r\n"

	default:
		return "\r\n"
	}
}

func httpProbe(host string) string {
	return fmt.Sprintf(
		"HEAD / HTTP/1.1\r\nHost: %s\r\nUser-Agent: GoScanner/1.0\r\nConnection: close\r\n\r\n",
		host,
	)
}
