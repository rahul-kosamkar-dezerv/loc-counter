package classifier

import (
	"strings"

	"loc-counter/syntax"
)

type LineType int

const (
	Blank LineType = iota
	Comment
	Code
	Import
	Declaration
)

type LineClassifier struct {
	Provider *syntax.Provider
}

func NewLineClassifier(p *syntax.Provider) LineClassifier {
	// Register built-in syntaxes
	p.Register(syntax.JavaSyntax{})
	p.Register(syntax.CSyntax{})
	p.Register(syntax.JSSyntax{})
	p.Register(syntax.PythonSyntax{})
	p.Register(syntax.PythonSyntax{})
	return LineClassifier{Provider: p}
}

func (c LineClassifier) Classify(line string, current syntax.Syntax, inBlock *bool) LineType {
	trimmed := strings.TrimSpace(line)

	if trimmed == "" {
		return Blank
	}

	// If already in a block comment, check for end
	if *inBlock {
		if current.IsMultiLineCommentEnd(line) {
			*inBlock = false
			if idx := strings.Index(line, currentEndToken(current)); idx >= 0 {
				after := strings.TrimSpace(line[idx+len(currentEndToken(current)):])
				if after == "" {
					return Comment
				}
				return Code
			}
			return Comment
		}
		return Comment
	}

	// Check block comment start
	if current.IsMultiLineCommentStart(line) {
		// if start and end on same line
		if current.IsMultiLineCommentEnd(line) && strings.Index(line, currentStartToken(current)) <= strings.Index(line, currentEndToken(current)) {
			startIdx := strings.Index(line, currentStartToken(current))
			endIdx := strings.Index(line, currentEndToken(current))
			before := strings.TrimSpace(line[:startIdx])
			after := strings.TrimSpace(line[endIdx+len(currentEndToken(current)):])
			if before == "" && after == "" {
				return Comment
			}
			return Code
		}
		*inBlock = true
		startIdx := strings.Index(line, currentStartToken(current))
		before := strings.TrimSpace(line[:startIdx])
		if before == "" {
			return Comment
		}
		return Code
	}

	// Single-line comment
	if current.IsSingleLineComment(trimmed) {
		return Comment
	}

	// import line
	if current.IsImportLine(line) {
		return Import
	}
	// declaration line (heuristic)
	if current.IsDeclarationLine(line) {
		return Declaration
	}

	if strings.Contains(line, "//") || (strings.Contains(line, "#") && !strings.HasPrefix(strings.TrimSpace(line), "#")) {
		return Code
	}

	return Code
}

// Helpers to detect token strings for multi-line comments per syntax
func currentStartToken(s syntax.Syntax) string {
	switch s.Name() {
	case "python":
		return `"""`
	default:
		return "/*"
	}
}
func currentEndToken(s syntax.Syntax) string {
	switch s.Name() {
	case "python":
		return `"""`
	default:
		return "*/"
	}
}
