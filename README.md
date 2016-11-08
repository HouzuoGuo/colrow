# colrow
Read lines from stdin and print a block of columns specified by first and last of column and row indexes.

## Build
Simply check out the source code and do `go build`.

## Usage
`./colrow col0 [row0 [last_col [last_row]]]`

Parameters:
- `col0` \- the index of column to begin with, lowest is 0 (first column).
- `row0` \- the index of row to begin with, lowest is 0 (first row). Default to 0 if not specified.
- `last_col` \- the index of final column (inclusive), shorter rows will not trigger out-of-bound access. Default to 4095 if not specified.
- `last_row` \- the index of final row (inclusive). Default to 4095 if not specified.
