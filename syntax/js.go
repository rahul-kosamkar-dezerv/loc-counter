package syntax

import "strings"

type JSSyntax struct{}

func (JSSyntax) Name() string         { return "js" }
func (JSSyntax) Extensions() []string { return []string{".js", ".jsx", ".ts", ".tsx"} }
func (JSSyntax) IsSingleLineComment(line string) bool {
	return strings.HasPrefix(strings.TrimSpace(line), "//")
}
func (JSSyntax) IsMultiLineCommentStart(line string) bool {
	return strings.Contains(line, "/*")
}
func (JSSyntax) IsMultiLineCommentEnd(line string) bool {
	return strings.Contains(line, "*/")
}
func (JSSyntax) IsImportLine(line string) bool {
	trim := strings.TrimSpace(line)
	return strings.HasPrefix(trim, "import ") || strings.Contains(trim, "require(")
}
func (JSSyntax) IsDeclarationLine(line string) bool {
	trim := strings.TrimSpace(line)
	return strings.HasPrefix(trim, "const ") || strings.HasPrefix(trim, "let ") || strings.HasPrefix(trim, "var ") || strings.HasPrefix(trim, "function ")
}
