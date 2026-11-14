package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"os"

	"go.yaml.in/yaml/v3"
	"golang.org/x/net/html/charset"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type Currency struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type CurrencyRates struct {
	XMLName    xml.Name   `xml:"ValCurs"`
	Currencies []Currency `xml:"Valute"`
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

	xmlData, err := os.ReadFile(config.InputFile)
	if err != nil {
		panic(fmt.Errorf("reading XML file: %w", err))
	}

	decoder := xml.NewDecoder(bytes.NewReader(xmlData))
	decoder.CharsetReader = charset.NewReaderLabel

	var rates CurrencyRates

	if err := decoder.Decode(&rates); err != nil {
		panic(fmt.Errorf("decoding XML: %w", err))
	}
}
