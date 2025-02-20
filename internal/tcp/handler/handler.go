package handler

import (
	"GIT/redis/cachebase/internal/entity"
	"GIT/redis/cachebase/internal/usecases/commands"
	"errors"
	"fmt"
	"net"
)

func HandleConnection(c net.Conn) error {
	for {
		var buf = make([]byte, 100)
		count, err := c.Read(buf)
		if err != nil || count == 0 {
			return fmt.Errorf("uanble to read buffer: %w", err)
		}
		readBuf := buf[:count]
		if err := ProcessMessage(c, readBuf); err != nil {
			return fmt.Errorf("unable to process message: %w", err)
		}
	}
}

func ProcessMessage(c net.Conn, data []byte) error {
	fmt.Printf("Processing data: %s", string(data))
	parsedData, data, err := entity.Parse(data)
	if err != nil {
		return fmt.Errorf("unable to parse data: %w", err)
	}
	if len(data) != 0 {
		return fmt.Errorf("not all data are processed, data left: %b", data)
	}
	arr, ok := parsedData.(entity.Array)
	if !ok {
		return errors.New("parsed command data should be an array")
	}
	command, ok := arr[0].(entity.BulkString)
	if !ok {
		return fmt.Errorf("command item should be string: %+v", arr[0])
	}
	args := entity.Array{}
	if len(arr) > 1 {
		args = arr[1:]
	}
	fmt.Printf("Processing %s command with following args %+v", command.Value, args)
	switch command.Value {
	case "PING":
		err = commands.Ping(c)
	case "ECHO":
		err = commands.Echo(c, args)
	case "SET":
		err = commands.Set(c, args)
	case "GET":
		err = commands.Get(c, args)
	default:
		return fmt.Errorf("unknown command: %s", command.Value)
	}
	if err != nil {
		return fmt.Errorf("command %s: %+w", command.Value, err)
	}
	return nil
}
