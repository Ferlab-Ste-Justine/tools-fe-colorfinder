package main

import (
    "bufio"
    "fmt"
    "os"
    "path/filepath"
    "regexp"
)

var colorList = []string{
    "gray", "blue", "red", "volcano", "orange", "gold", "yellow",
    "lime", "green", "cyan", "geekblue", "purple", "magenta",
}

func main() {
    // Compile regex pattern based on color list
    pattern := fmt.Sprintf(`--(%s)-([1-9]|10)`, joinColors(colorList))
    regex := regexp.MustCompile(pattern)

    // Specify the file extensions to search
    extensions := []string{".css", ".less", ".ts", ".tsx", ".js", ".jsx"}

    // Walk through the current directory and process files
    err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() && hasExtension(path, extensions) {
            scanFile(path, regex)
        }
        return nil
    })

    if err != nil {
        fmt.Println("Error walking the path:", err)
    }
}

// joinColors joins color names with a regex OR (|) separator
func joinColors(colors []string) string {
    return regexp.QuoteMeta(colors[0]) + "|" + regexp.QuoteMeta(colors[1]) // Join colors for regex
	}

// hasExtension checks if the file has one of the specified extensions
func hasExtension(path string, extensions []string) bool {
    ext := filepath.Ext(path)
    for _, e := range extensions {
        if e == ext {
            return true
        }
    }
    return false
}

// scanFile scans a file for lines matching the regex and prints matches
func scanFile(path string, regex *regexp.Regexp) {
    file, err := os.Open(path)
    if err != nil {
        fmt.Println("Error opening file:", path, err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    lineNumber := 1
    for scanner.Scan() {
        line := scanner.Text()
        if regex.MatchString(line) {
					fmt.Printf("found : %s at line %d: %s\n", path, lineNumber, line)
        }
        lineNumber++
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", path, err)
    }
}
