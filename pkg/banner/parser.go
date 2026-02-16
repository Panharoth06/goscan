package banner

import (
	"bytes"
	"strings"
)

func DetectProtocol(data []byte, port int) string {
	s := strings.ToLower(string(data))

	switch {
	case strings.Contains(s, "http/"):
		return "http"
	case strings.Contains(s, "ssh-"):
		return "ssh"
	case strings.Contains(s, "smtp"):
		return "smtp"
	case port == 3306:
		return "mysql"
	default:
		return "unknown"
	}
}

func ParseBanner(data []byte) string {
	clean := cleanASCII(data)

	lines := strings.Split(clean, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			if len(line) > 120 {
				return line[:117] + "..."
			}
			return line
		}
	}

	return ""
}

func cleanASCII(data []byte) string {
	var buf bytes.Buffer
	for _, b := range data {
		if b >= 32 && b <= 126 || b == '\n' || b == '\r' || b == '\t' {
			buf.WriteByte(b)
		}
	}
	return buf.String()
}
