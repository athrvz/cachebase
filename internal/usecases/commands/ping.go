package commands

import (
	"GIT/redis/cachebase/internal/entity"
	"fmt"
	"net"
)

func Ping(c net.Conn) error {
	response := entity.SimpleString("PONG").Encode()
	if _, err := c.Write([]byte(response)); err != nil {
		return fmt.Errorf("unable to write response buffer: %w", err)
	}
	return nil
}
