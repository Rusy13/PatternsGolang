package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type cutOptions struct {
	fields    string
	delimiter string
	separated bool
}

func parseFlags() cutOptions {
	var options cutOptions
	flag.StringVar(&options.fields, "f", "", "select fields (columns)")
	flag.StringVar(&options.delimiter, "d", "\t", "use different delimiter")
	flag.BoolVar(&options.separated, "s", false, "only lines with delimiter")
	flag.Parse()
	return options
}

func cutLine(line string, options cutOptions) string {
	parts := strings.Split(line, options.delimiter)
	fields := strings.Split(options.fields, ",")
	var result []string

	for _, field := range fields {
		index := 0
		fmt.Sscanf(field, "%d", &index)
		if index > 0 && index <= len(parts) {
			result = append(result, parts[index-1])
		}
	}
	return strings.Join(result, options.delimiter)
}

func main() {
	options := parseFlags()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, options.delimiter) && options.separated {
			continue
		}
		fmt.Println(cutLine(line, options))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}
