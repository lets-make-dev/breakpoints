package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Breakpoint structure to store filename and line number
type Breakpoint struct {
	File string
	Line int
}

// XDebuggerManager struct to represent the XML node
type XDebuggerManager struct {
	XMLName           xml.Name          `xml:"component"`
	Name              string            `xml:"name,attr"`
	BreakpointManager BreakpointManager `xml:"breakpoint-manager"`
}

// BreakpointManager struct to represent the breakpoint manager node
type BreakpointManager struct {
	Breakpoints []LineBreakpoint `xml:"breakpoints>line-breakpoint"`
}

// LineBreakpoint struct to represent each breakpoint
type LineBreakpoint struct {
	Enabled string `xml:"enabled,attr"`
	Type    string `xml:"type,attr"`
	URL     string `xml:"url"`
	Line    string `xml:"line"`
}

func clearBreakpoints() {
	filePath := getWorkspacePath()
	xmlContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("An error occurred while reading the file: %v\n", err)
		return
	}

	// Define the regex pattern to find the XDebuggerManager component
	pattern := `<component name="XDebuggerManager">[\s\S]*?<\/component>`

	// Compile the regex pattern
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Printf("An error occurred while compiling the regex: %v\n", err)
		return
	}

	// Replace the XDebuggerManager component with an empty string
	updatedContent := re.ReplaceAllString(string(xmlContent), "")

	// Write the updated content back to the file
	err = ioutil.WriteFile(filePath, []byte(updatedContent), 0644)
	if err != nil {
		fmt.Printf("An error occurred while writing the updated content to the file: %v\n", err)
		return
	}
}

// setBreakpoints sets the breakpoints in the XML file
func setBreakpoints(breakpoints []Breakpoint) {
	// Create XDebuggerManager
	xDebuggerManager := XDebuggerManager{
		Name: "XDebuggerManager",
		BreakpointManager: BreakpointManager{
			Breakpoints: []LineBreakpoint{},
		},
	}

	// Create line breakpoints for each breakpoint
	for _, breakpoint := range breakpoints {
		xDebuggerManager.BreakpointManager.Breakpoints = append(
			xDebuggerManager.BreakpointManager.Breakpoints,
			LineBreakpoint{
				Enabled: "true",
				Type:    "php",
				URL:     fmt.Sprintf("file://$PROJECT_DIR$/%s", breakpoint.File),
				Line:    strconv.Itoa(breakpoint.Line),
			},
		)
	}

	// Convert XDebuggerManager struct into xml
	managerData, err := xml.MarshalIndent(xDebuggerManager, "  ", "    ")
	if err != nil {
		fmt.Printf("Failed to marshal the XML file: %v\n", err)
		return
	}

	// Read the existing XML
	xmlPath := getWorkspacePath()
	data, err := ioutil.ReadFile(xmlPath)
	if err != nil {
		fmt.Printf("Failed to read the XML file: %v\n", err)
		return
	}

	// Replace "</project>" with the new component + "</project>"
	newData := strings.Replace(string(data), "</project>", "\n"+string(managerData)+"</project>", 1)

	// Write XML back to file
	err = ioutil.WriteFile(xmlPath, []byte(newData), 0644)
	if err != nil {
		fmt.Printf("Failed to write the XML file: %v\n", err)
		return
	}
}

func main() {
	if len(os.Args) < 2 {
		//printHelp()
		//return
	}

	if len(os.Args) == 2 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			printHelp()
			return
		}

		if os.Args[1] == "clear" {
			clearBreakpoints()
			return
		}
	}
	// Check if the XML file exists

	// Regular expression to match <filename>:<line>
	re := regexp.MustCompile(`([\\\/\w\-. ]+):(\d+)`)

	// Read input from arguments or standard input
	var input string
	if len(os.Args) > 1 {
		input = strings.Join(os.Args[1:], " ")
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input += scanner.Text() + "\n"
		}
	}

	// Extract matches and gather unique breakpoints
	uniqueBreakpoints := make(map[string]Breakpoint)
	matches := re.FindAllStringSubmatch(input, -1)

	// print matches
	for _, match := range matches {
		file := strings.TrimSpace(match[1])
		line, err := strconv.Atoi(match[2])
		if err != nil {
			fmt.Printf("Invalid line number: %s\n", match[2])
			continue
		}
		key := fmt.Sprintf("%s:%d", file, line-1)
		uniqueBreakpoints[key] = Breakpoint{File: file, Line: line - 1}
	}

	// Convert map values to a slice
	var uniqueBreakpointsSlice []Breakpoint
	for _, value := range uniqueBreakpoints {
		uniqueBreakpointsSlice = append(uniqueBreakpointsSlice, value)
	}

	// Clear existing breakpoints
	clearBreakpoints()

	// Send to "set breakpoints" function
	setBreakpoints(uniqueBreakpointsSlice)
}
