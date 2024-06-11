package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type grepOptions struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
}

func parseFlags() grepOptions {
	var options grepOptions
	flag.IntVar(&options.after, "A", 0, "print +N lines after match")
	flag.IntVar(&options.before, "B", 0, "print +N lines before match")
	flag.IntVar(&options.context, "C", 0, "print ±N lines around match")
	flag.BoolVar(&options.count, "c", false, "print count of matching lines")
	flag.BoolVar(&options.ignoreCase, "i", false, "ignore case")
	flag.BoolVar(&options.invert, "v", false, "invert match")
	flag.BoolVar(&options.fixed, "F", false, "fixed string match")
	flag.BoolVar(&options.lineNum, "n", false, "print line number")
	flag.Parse()
	return options
}

func grep(lines []string, pattern string, options grepOptions) ([]string, error) {
	var result []string
	var matcher func(string) bool
	matchCount := 0

	if options.ignoreCase {
		pattern = "(?i)" + pattern
	}

	if options.fixed {
		matcher = func(line string) bool {
			return strings.Contains(line, pattern)
		}
	} else {
		re, err := regexp.Compile(pattern)
		if err != nil {
			return nil, err
		}
		matcher = re.MatchString
	}

	for i, line := range lines {
		matched := matcher(line)
		if options.invert {
			matched = !matched
		}

		if matched {
			matchCount++
			if options.count {
				continue
			}
			if options.lineNum {
				result = append(result, fmt.Sprintf("%d:%s", i+1, line))
			} else {
				result = append(result, line)
			}
			if options.after > 0 || options.context > 0 {
				afterCount := options.after
				if options.context > 0 {
					afterCount = options.context
				}
				for j := 1; j <= afterCount && i+j < len(lines); j++ {
					result = append(result, lines[i+j])
				}
			}
		} else {
			if options.before > 0 || options.context > 0 {
				beforeCount := options.before
				if options.context > 0 {
					beforeCount = options.context
				}
				if i-beforeCount >= 0 {
					result = append(result, lines[i-beforeCount:i]...)
				}
			}
		}
	}

	if options.count {
		result = append(result, fmt.Sprintf("Count: %d", matchCount))
	}
	return result, nil
}

func main() {
	options := parseFlags()
	pattern := flag.Arg(0)
	filename := flag.Arg(1)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	result, err := grep(lines, pattern, options)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing grep: %v\n", err)
		os.Exit(1)
	}

	for _, line := range result {
		fmt.Println(line)
	}
}
