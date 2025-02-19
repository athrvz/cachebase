package commands 

import (
  "fmt"
  "github.com/athrvz/cachebase/internal/entity"
  "net"
)

func Ping(c net.Conn) error {
  response := entity.SimpleString("PONG").Encode()
  if _, err := c.Write([]btye(response)); err != nil {
    return fmt.Errorf("unable to write response buffer: %w", err)
  }
  return nil 
}
