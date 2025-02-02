package entity

import "fmt"

func (d SimpleString) Enconde() string {
	return fmt.Sprintf("+%s\r\n", d)
}
