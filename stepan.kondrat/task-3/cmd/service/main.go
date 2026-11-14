package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

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

type CurrencyResult struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
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

	results := make([]CurrencyResult, 0, len(rates.Currencies))

	for _, currency := range rates.Currencies {
		valueText := strings.ReplaceAll(currency.Value, ",", ".")

		value, err := strconv.ParseFloat(valueText, 64)
		if err != nil {
			panic(fmt.Errorf(
				"parsing value for currency %q: %w",
				currency.CharCode,
				err,
			))
		}

		results = append(results, CurrencyResult{
			NumCode:  currency.NumCode,
			CharCode: currency.CharCode,
			Value:    value,
		})
	}

	sort.SliceStable(results, func(i, j int) bool {
		return results[i].Value > results[j].Value
	})

	jsonData, err := json.MarshalIndent(results, "", "    ")
	if err != nil {
		panic(fmt.Errorf("encoding currencies as JSON: %w", err))
	}

	outputDir := filepath.Dir(config.OutputFile)

	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		panic(fmt.Errorf("creating output directory: %w", err))
	}

	if err := os.WriteFile(config.OutputFile, jsonData, 0o644); err != nil {
		panic(fmt.Errorf("writing JSON file: %w", err))
	}
}
