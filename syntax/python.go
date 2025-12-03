package syntax

import "strings"

type PythonSyntax struct{}

func (PythonSyntax) Name() string { return "python" }
func (PythonSyntax) Extensions() []string { return []string{".py"} }
func (PythonSyntax) IsSingleLineComment(line string) bool {
	return strings.HasPrefix(strings.TrimSpace(line), "#")
}
func (PythonSyntax) IsMultiLineCommentStart(line string) bool {
	trim := strings.TrimSpace(line)
	return strings.HasPrefix(trim, `"""`) || strings.HasPrefix(trim, "'''")
}
func (PythonSyntax) IsMultiLineCommentEnd(line string) bool {
	trim := strings.TrimSpace(line)
	return strings.HasSuffix(trim, `"""`) || strings.HasSuffix(trim, "'''")
}
func (PythonSyntax) IsImportLine(line string) bool {
	trim := strings.TrimSpace(line)
	return strings.HasPrefix(trim, "import ") || strings.HasPrefix(trim, "from ")
}
func (PythonSyntax) IsDeclarationLine(line string) bool {
	trim := strings.TrimSpace(line)
	// assignment at top-level or def/class
	return strings.Contains(trim, " = ") || strings.HasPrefix(trim, "@") || strings.HasPrefix(trim, "def ") || strings.HasPrefix(trim, "class ")
}
