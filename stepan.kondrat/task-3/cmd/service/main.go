package main

import (
	"flag"
	"fmt"
	"os"

	"go.yaml.in/yaml/v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func main() {
	configPath := flag.String("config", "", "Path to YAML configuration")
	flag.Parse()

	if *configPath == "" {
		panic("configuration path is required: use --config")
	}

	data, err := os.ReadFile(*configPath)
	if err != nil {
		panic(fmt.Errorf("reading configuration: %w", err))
	}

	var config Config

	if err := yaml.Unmarshal(data, &config); err != nil {
		panic(fmt.Errorf("decoding configuration: %w", err))
	}

	if config.InputFile == "" || config.OutputFile == "" {
		panic("configuration must contain non-empty input-file and output-file")
	}

	fmt.Println("Input file:", config.InputFile)
	fmt.Println("Output file:", config.OutputFile)
}
