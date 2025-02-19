package entity

import "fmt"

func (d SimpleString) Encode() string {
	return fmt.Sprintf("+%s\r\n", d)
}
