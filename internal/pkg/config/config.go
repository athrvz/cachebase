package config 

import "os"

type Config struct {
  APP string
  ENV string
  SERVER struct {
    Host string
    Port string
  }
}

func NewConfig() (*Config, error) {
  var config Config 
  
  config.APP = getEnv("APP", "app")
  config.ENV = getEnv("ENV", "dev")
  
  config.SERVER.Host = getEnv("SERVER_HOST", "localhost")
  config.SERVER.Port = getEnv("SERVER_PORT", ":6379")
  
  return &config, nil
}

func getEnv(key string, defaultValue string) string {
  value, exists := os.LookupEnv(key)
  if exists {
    return value
  }
  return defaultValue
}
