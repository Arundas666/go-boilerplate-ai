package helper

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CreateFolderStructure(structure string, baseDir string) {
	paths := parseStructure(structure)
	for _, path := range paths {
		fullPath := filepath.Join(baseDir, path) // Join base directory with the intended path
		if strings.HasSuffix(path, "/") {
			// It's a directory
			err := os.MkdirAll(fullPath, os.ModePerm)
			if err != nil {
				fmt.Printf("Error creating directory %s: %v\n", fullPath, err)
			}
		} else {
			// It's a file
			_, err := os.Create(fullPath)
			if err != nil {
				fmt.Printf("Error creating file %s: %v\n", fullPath, err)
			}
		}
	}
}

func parseStructure(input string) []string {
    var stack []string
    var result []string

    lines := strings.Split(input, "\n")
    for _, line := range lines {
        line = strings.TrimSpace(line)
        // Skip empty lines or lines without intended structure chars
        if line == "" || !(strings.Contains(line, "├── ") || strings.Contains(line, "└── ")) {
            continue
        }

        // Count the prefixes to determine the depth
        prefixCount := strings.Count(line, "│   ")
        line = strings.TrimLeft(line, "│   ")
        line = strings.TrimPrefix(line, "├── ")
        line = strings.TrimPrefix(line, "└── ")

        // Pop from stack to get to correct depth (prefix count plus one for └── or ├──)
        currentDepth := len(stack)
        if prefixCount < currentDepth {
            stack = stack[:prefixCount]
        }

        // Append the current element to the path
        currentPath := ""
        if len(stack) > 0 {
            currentPath = strings.Join(stack, "/") + "/"
        }
        currentPath += line

        // Add / to directories (identified by trailing /)
        if strings.HasSuffix(currentPath, "/") && !strings.HasSuffix(line, "/") {
            currentPath += "/"
        }

        result = append(result, currentPath)

        // If it's a directory, push it onto stack
        if strings.HasSuffix(line, "/") {
            stack = append(stack, line)
        }
    }

    return result
}
