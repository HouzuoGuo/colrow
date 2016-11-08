package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Print an error message to stderr, and exit the program. The function return value is useless.
func errorExit(msg string, exitCode int) int {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(exitCode)
	return 0
}

// Return a positive integer converted from the input string, or exit if there is no integer or it is negative.
func mustGetNonNegativeInt(str string) int {
	ret, err := strconv.Atoi(str)
	if err != nil {
		return errorExit(fmt.Sprintf("\"%s\" is not a positive integer", str), 1)
	}
	return ret
}

// Among the input lines (in a single string), pick a block of columns specified by the index and return.
func getColumns(in string, col0, row0, lastCol, lastRow int) string {
	if lastCol == 0 {
		lastCol = 4095
	}
	if lastRow == 0 {
		lastRow = 4095
	}
	lines := strings.Split(in, "\n")
	ret := make([]string, 0, lastRow-row0)
	for row := row0; row < len(lines) && row <= lastRow; row++ {
		// Trailing white spaces can be ignored
		txt := strings.TrimRightFunc(lines[row], unicode.IsSpace)
		if col0 >= len(txt) {
			// Column is wider than the line
			ret = append(ret, "")
			continue
		}
		// Retrieve columns from the current row
		lastColIndex := lastCol + 1
		if lenTxt := len(txt); lastColIndex > lenTxt {
			lastColIndex = lenTxt
		}
		// Train white spaces can be once again ignored
		trimmedColumns := strings.TrimRightFunc(txt[col0:lastColIndex], unicode.IsSpace)
		ret = append(ret, trimmedColumns)
	}
	return strings.Join(ret, "\n")
}

func main() {
	if len(os.Args) <= 1 {
		errorExit("Usage: colrow col0 [row0 [last_col [last_row]]]", 1)
		return
	}
	var col0, row0, lastCol, lastRow int

	if len(os.Args) >= 2 {
		col0 = mustGetNonNegativeInt(os.Args[1])
	}
	if len(os.Args) >= 3 {
		row0 = mustGetNonNegativeInt(os.Args[2])
	}
	if len(os.Args) >= 4 {
		lastCol = mustGetNonNegativeInt(os.Args[3])
	}
	if len(os.Args) >= 5 {
		lastRow = mustGetNonNegativeInt(os.Args[4])
	}
	if lastCol != 0 && lastCol <= col0 || lastRow != 0 && lastRow <= row0 {
		errorExit("Last column/row must be greater than first column/row", 1)
		return
	}
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		errorExit(fmt.Sprintf("Failed to read from stdin: %v", err), 1)
		return
	}
	fmt.Println(getColumns(string(in), col0, row0, lastCol, lastRow))
}
