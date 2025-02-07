package handler

import (
	"GIT/redis/cachebase/internal/entity"
	"errors"
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

func ProcessMessage(c net.Conn, data []byte) error {
	fmt.Printf("Processing data: %s", string(data))
	parsedData, data, err := entity.Parse(data)
	if err != nil {
		return fmt.Errorf("unable to parse data: %w", err)
	}
	if len(data) != 0 {
		return fmt.Errorf("whole data not processed, data left: %b", data)
	}
	arr, ok := parsedData.(entity.Array)
	if !ok {
		return errors.New("parsed command should be an array")
	}
	command, ok := arr[0].(entity.BulkString)
	if !ok {
		return fmt.Errorf("command item should be a string: %+v", arr[0])
	}
	args := entity.Array{}
	if len(arr) > 1 {
		args = arr[1:]
	}
}
