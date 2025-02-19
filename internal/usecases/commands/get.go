package commands

import (
	"GIT/redis/cachebase/internal/entity"
	"GIT/redis/cachebase/internal/pkg/storage/cache"
	"fmt"
	"net"
)

func Get(c net.Conn, args entity.Array) error {
	if len(args) != 1 {
		return fmt.Errorf("incorrect number of arguments: %+v", args)
	}
	key, ok := args[0].(entity.BulkString)
	if !ok {
		return fmt.Errorf("key should be a string: %+v", args[0])
	}
	response := entity.BulkString{IsNull: true}.Enconde()
	if val, ok := cache.Get(key.Value); ok {
		response = entity.BulkString{
			Value:  val,
			IsNull: false,
		}.Enconde()
	}
	if _, err := c.Write([]byte(response)); err != nil {
		return fmt.Errorf("uanble to write response buffer: %w", err)
	}
	return nil
}
