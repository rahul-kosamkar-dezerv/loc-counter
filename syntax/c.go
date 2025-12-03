package syntax

import "strings"

type CSyntax struct{}

func (CSyntax) Name() string { return "c" }
func (CSyntax) Extensions() []string { return []string{".c", ".h", ".cpp", ".cc"} }
func (CSyntax) IsSingleLineComment(line string) bool {
	return strings.HasPrefix(strings.TrimSpace(line), "//")
}
func (CSyntax) IsMultiLineCommentStart(line string) bool {
	return strings.Contains(line, "/*")
}
func (CSyntax) IsMultiLineCommentEnd(line string) bool {
	return strings.Contains(line, "*/")
}
func (CSyntax) IsImportLine(line string) bool {
	trim := strings.TrimSpace(line)
	return strings.HasPrefix(trim, "#include")
}
func (CSyntax) IsDeclarationLine(line string) bool {
	trim := strings.TrimSpace(line)
	types := []string{"int ", "long ", "double ", "float ", "char ", "bool ", "void "}
	for _, t := range types {
		if strings.HasPrefix(trim, t) && (strings.HasSuffix(trim, ";") || strings.Contains(trim, "(")) {
			return true
		}
	}
	return false
}
