package commands

import (
	"GIT/redis/cachebase/internal/entity"
	"fmt"
	"net"
)

func Echo(c net.Conn, args entity.Array) error {
	if len(args) != 1 {
		return fmt.Errorf("incorrect number of arguments: %+v", args)
	}
	echoed, ok := args[0].(entity.BulkString)
	if !ok {
		return fmt.Errorf("echoed should be a string: %+v", args[0])
	}
	response := entity.SimpleString(echoed.Value).Encode()
	if _, err := c.Write([]byte(response)); err != nil {
		return fmt.Errorf("unable to write repose buffer: %w", err)
	}
	return nil
}
