# LOC Counter

Lines of Code counter supporting multiple languages, multi-file scanning, multi-line comments and basic granular classification (imports, declarations).

## Features

- Multiple syntaxes: Java, C/C++, JS/TS, Python
- Multi-line (directory) scanning with totals
- Multi-line comment support (/_ ... _/ and Python triple-quotes)
- Granular categories: Imports, Declarations, Comments, Blank, Code
- CLI: choose a file or directory. Auto-detect language by extension.

## Quickstart

```bash
go run cmd/main.go fixtures/multi
```

Sample output:

```
Files scanned: 3
Blank: 7
Comments: 9
Code: 13
Imports: 4
Declarations: 3
Total lines: 32
```
// Final polishing
