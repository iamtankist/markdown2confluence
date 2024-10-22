package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	inputFile := flag.String("input", "", "Path to the input file")
	outputFile := flag.String("output", "", "Path to the output file")
	flag.Parse()

	if *inputFile == "" || *outputFile == "" {
		fmt.Println("Both input and output file paths are required")
		return
	}

	file, err := os.Open(*inputFile)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	var title string
	var content []string
	var attachments []string
	inFrontMatter := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "---" {
			inFrontMatter = !inFrontMatter
			continue
		}
		if inFrontMatter {
			if strings.HasPrefix(line, "title:") {
				filename := strings.TrimSuffix(filepath.Base(*inputFile), filepath.Ext(*inputFile))
				title = strings.TrimPrefix(strings.TrimSuffix(strings.TrimSpace(strings.TrimPrefix(line, "title:")), "\""), "\"")
				title = filename[:4] + " " + title
			}
		} else {
			content = append(content, line)
			// Find image references
			re := regexp.MustCompile(`!\[.*?\]\((.*?)\)`)
			matches := re.FindAllStringSubmatch(line, -1)
			for _, match := range matches {
				if len(match) > 1 {
					attachments = append(attachments, match[1])
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	output, err := os.Create(*outputFile)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return
	}
	defer output.Close()

	writer := bufio.NewWriter(output)

	// Write the metadata for mark files
	writer.WriteString(fmt.Sprintf("<!-- Title: %s -->\n", title))
	for _, attachment := range attachments {
		writer.WriteString(fmt.Sprintf("<!-- Attachment: %s -->\n", attachment))
	}
	writer.WriteString("\n")

	// Write the content
	for _, line := range content {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Printf("Error writing to output file: %v\n", err)
			return
		}
	}
	writer.Flush()

	fmt.Printf("Title: %s\n", title)
}
