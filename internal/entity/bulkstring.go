package entity

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

func ParseBulkString(data []byte) (BulkString, []byte, error) {
	bulkString := BulkString{
		Value:  "",
		IsNull: false,
	}
	if len(data) < 6 {
		return bulkString, data, errors.New("bulk string needs atleast 6 characters")
	}
	data = data[1:]
	bulkData := bytes.SplitN(data, []byte(CLRF), 2)
	if len(bulkData) != 2 {
		return bulkString, data, errors.New("Bulk string needs start deliminator")
	}
	data = bulkData[1]
	length, err := strconv.Atoi(string(bulkData[0]))
	if err != nil {
		return bulkString, data, fmt.Errorf("length is not an integer: %s", bulkData[0])
	}
	if length == -1 {
		bulkString.IsNull = true
		return bulkString, data, nil
	}
	bulkData = bytes.SplitN(data, []byte(CLRF), 2)
	if len(bulkData) != 2 {
		return bulkString, data, errors.New("bulk string needs an end deliminator")
	}
	data = bulkData[1]
	bulkString.Value = string(bulkData[0])
	return bulkString, data, nil
}

func (d BulkString) Enconde() string {
	if d.IsNull {
		return "$-1" + CLRF
	}
	return fmt.Sprintf("$%d\r\n%s\r\n", len(d.Value), d.Value)
}
