package entity

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func ParseArray(data []byte) (Array, []byte, error) {
	if len(data) < 4 {
		return nil, data, errors.New("array needs at least 4 characters")
	}
	data = data[1:]
	arrayData := bytes.SplitN(data, []byte(CLRF), 2)

	for _, line := range arrayData {
		fmt.Println("line: ", string(line))
	}
	if len(arrayData) != 2 {
		return nil, data, errors.New("Array needs start deliminator")
	}
	data = arrayData[1]
	stringNum := strings.Trim(string(arrayData[0]), " ")
	length, err := strconv.Atoi(stringNum)
	fmt.Println(err)
	if err != nil {
		return nil, data, fmt.Errorf("length is not an interger: %s", arrayData[0])
	}
	array := Array{}
	for i := 0; i < length; i++ {
		arrayItem, dataLeft, err := Parse(data)
		if err != nil {
			return nil, data, fmt.Errorf("failed to parse array item: %s, due to %w", arrayData[0], err)
		}
		array = append(array, arrayItem)
		data = dataLeft
	}
	return array, data, nil
}
