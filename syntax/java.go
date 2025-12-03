package syntax

import "strings"

type JavaSyntax struct{}

func (JavaSyntax) Name() string { return "java" }
func (JavaSyntax) Extensions() []string { return []string{".java"} }
func (JavaSyntax) IsSingleLineComment(line string) bool {
	return strings.HasPrefix(strings.TrimSpace(line), "//")
}
func (JavaSyntax) IsMultiLineCommentStart(line string) bool {
	return strings.Contains(line, "/*")
}
func (JavaSyntax) IsMultiLineCommentEnd(line string) bool {
	return strings.Contains(line, "*/")
}
func (JavaSyntax) IsImportLine(line string) bool {
	return strings.HasPrefix(strings.TrimSpace(line), "import ")
}
func (JavaSyntax) IsDeclarationLine(line string) bool {
	trim := strings.TrimSpace(line)
	// simple heuristic: presence of access modifier or common types + identifier + ;
	if strings.HasPrefix(trim, "public ") || strings.HasPrefix(trim, "private ") || strings.HasPrefix(trim, "protected ") {
		return true
	}
	types := []string{"int ", "long ", "double ", "float ", "String ", "boolean ", "char ", "short "}
	for _, t := range types {
		if strings.Contains(trim, t) && (strings.HasSuffix(trim, ";") || strings.Contains(trim, " = ")) {
			return true
		}
	}
	return false
}
// TODO: Improve detection of generic types
