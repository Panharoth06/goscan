package banner

import (
	"bytes"
	"io"
	"net"
	"time"
)

func readAll(conn net.Conn) ([]byte, error) {
	var buf bytes.Buffer
	tmp := make([]byte, 4096)

	for {
		conn.SetReadDeadline(time.Now().Add(readTimeout))
		n, err := conn.Read(tmp)
		if n > 0 {
			buf.Write(tmp[:n])
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			break
		}

		if buf.Len() >= maxReadSize {
			break
		}
	}

	return buf.Bytes(), nil
}

func sendProbe(conn net.Conn, probe string) ([]byte, error) {
	if probe != "" {
		conn.SetWriteDeadline(time.Now().Add(readTimeout))
		_, err := conn.Write([]byte(probe))
		if err != nil {
			return nil, err
		}
	}
	return readAll(conn)
}
