package commands

import (
	"GIT/redis/cachebase/internal/entity"
	"GIT/redis/cachebase/internal/pkg/storage/cache"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func Set(c net.Conn, args entity.Array) error {
	if len(args) < 2 {
		return fmt.Errorf("incorrect number of arguments: %+v", args)
	}
	key, ok := args[0].(entity.BulkString)
	if !ok {
		return fmt.Errorf("key should be a string: %+v", args[0])
	}
	val, ok := args[1].(entity.BulkString)
	if !ok {
		return fmt.Errorf("value should be a string: %+v", args[1])
	}
	options, err := parseSetOptions(args[2:])
	if err != nil {
		return fmt.Errorf("unable to parse set option: %w", err)
	}
	cache.Set(key.Value, val.Value, options)
	response := entity.SimpleString("OK").Encode()
	if _, err := c.Write([]byte(response)); err != nil {
		return fmt.Errorf("unable to write response buffer: %w", err)
	}
	return nil
}

func parseSetOptions(args entity.Array) (cache.SetOptions, error) {
	options := cache.SetOptions{}
	for len(args) > 0 {
		cmdArg, ok := args[0].(entity.BulkString)
		if !ok {
			return options, fmt.Errorf("option name should be bulk string: %+v", args[0])
		}
		switch strings.ToLower(cmdArg.Value) {
		case "px":
			expiryStr, ok := args[1].(entity.BulkString)
			if !ok {
				return options, fmt.Errorf("expiry value should be bulk string: %+v", args[1])
			}
			expiry, err := strconv.Atoi(expiryStr.Value)
			if err != nil {
				return options, fmt.Errorf("expiry value should be able to convert to integer: %+v", args[1])
			}
			options.Expiry = &expiry
			args = args[2:]
		default:
			return options, fmt.Errorf("unknown set option name: %s", cmdArg.Value)
		}
	}
	return options, nil
}
