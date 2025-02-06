package handler

import (
	"fmt"
	"net"
)

func HandleConnection(c net.Conn) error {
	for {
		var buf = make([]byte, 100)
		count, err := c.Read(buf)
		if err != nil || count == 0 {
			return fmt.Errorf("unable to read buffer: %w", err)
		}
		readBuf := buf[:count]
		if err := ProcessMessage(c, readBuf); err != nil {
			return fmt.Errorf("unable to process the message: %w", err)
		}
	}
}
