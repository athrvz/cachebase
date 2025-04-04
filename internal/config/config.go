package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type RDBConfig struct {
	// Save intervals
	SaveIntervals []struct {
		Seconds int `yaml:"seconds"` // Time window in seconds
		Changes int `yaml:"changes"` // Minimum writes to trigger save
	} `yaml:"save_intervals"`
	FilePath         string `yaml:"filepath"`           // eg: "/var/cachebase/dump.rdb"
	TempFilePattern  string `yaml:"temp_file_pattern"`  // eg: "temp-*.rdb"
	CdbFileExtension string `yaml:"cdb_file_extension"` // eg: "gz"
	// Optimization flags
	BackgroundSave     bool   `yaml:"background_save"`      // Non-blocking save
	Compression        string `yaml:"compression"`          // "none", "gzip"
	MaxConcurrentSaves int    `yaml:"max_concurrent_saves"` // Default: 1
}

type Config struct {
	RDB RDBConfig `yaml:"rdb"`
}

func LoadConfig(filename string) Config {
	// Read yaml file
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("error reading yaml file: %v", err)
	}
	// Unmarshal the yaml file into RDBConfig
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error parsing yaml file: %v", err)
	}
	return config
}
