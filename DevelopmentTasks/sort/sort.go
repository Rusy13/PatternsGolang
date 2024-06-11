package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type sortOptions struct {
	column     int
	numeric    bool
	reverse    bool
	unique     bool
	month      bool
	ignoreTail bool
	check      bool
	human      bool
}

func parseFlags() sortOptions {
	var options sortOptions
	flag.IntVar(&options.column, "k", 0, "specify column for sorting")
	flag.BoolVar(&options.numeric, "n", false, "sort by numeric value")
	flag.BoolVar(&options.reverse, "r", false, "sort in reverse order")
	flag.BoolVar(&options.unique, "u", false, "do not output duplicate lines")
	flag.BoolVar(&options.month, "M", false, "sort by month name")
	flag.BoolVar(&options.ignoreTail, "b", false, "ignore trailing spaces")
	flag.BoolVar(&options.check, "c", false, "check if sorted")
	flag.BoolVar(&options.human, "h", false, "sort by numeric value with suffixes")
	flag.Parse()
	return options
}

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeLines(lines []string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	return nil
}

func sortLines(lines []string, options sortOptions) []string {
	if options.ignoreTail {
		for i := range lines {
			lines[i] = strings.TrimRight(lines[i], " ")
		}
	}

	if options.numeric {
		sort.Slice(lines, func(i, j int) bool {
			var a, b float64
			var err error
			if options.column > 0 {
				a, err = strconv.ParseFloat(strings.Fields(lines[i])[options.column-1], 64)
				if err != nil {
					a = 0
				}
				b, err = strconv.ParseFloat(strings.Fields(lines[j])[options.column-1], 64)
				if err != nil {
					b = 0
				}
			} else {
				a, err = strconv.ParseFloat(lines[i], 64)
				if err != nil {
					a = 0
				}
				b, err = strconv.ParseFloat(lines[j], 64)
				if err != nil {
					b = 0
				}
			}
			if options.reverse {
				return a > b
			}
			return a < b
		})
	} else {
		sort.Slice(lines, func(i, j int) bool {
			var a, b string
			if options.column > 0 {
				a = strings.Fields(lines[i])[options.column-1]
				b = strings.Fields(lines[j])[options.column-1]
			} else {
				a = lines[i]
				b = lines[j]
			}
			if options.reverse {
				return a > b
			}
			return a < b
		})
	}

	if options.unique {
		uniqLines := make([]string, 0, len(lines))
		prev := ""
		for _, line := range lines {
			if line != prev {
				uniqLines = append(uniqLines, line)
				prev = line
			}
		}
		return uniqLines
	}

	return lines
}

func main() {
	options := parseFlags()
	inputFile := flag.Arg(0)
	outputFile := flag.Arg(1)

	lines, err := readLines(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	if options.check {
		sortedLines := sortLines(lines, options)
		for i, line := range lines {
			if line != sortedLines[i] {
				fmt.Fprintf(os.Stderr, "File is not sorted\n")
				os.Exit(1)
			}
		}
		fmt.Println("File is sorted")
		return
	}

	sortedLines := sortLines(lines, options)

	err = writeLines(sortedLines, outputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %v\n", err)
		os.Exit(1)
	}
}
