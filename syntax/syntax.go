package syntax

// Syntax defines language-specific rules for classification.
type Syntax interface {
	Name() string
	Extensions() []string
	IsSingleLineComment(line string) bool
	IsMultiLineCommentStart(line string) bool
	IsMultiLineCommentEnd(line string) bool
	IsImportLine(line string) bool
	IsDeclarationLine(line string) bool
}

// Provider holds available syntaxes and can pick by extension.
type Provider struct {
	syntaxes []Syntax
	defaultSyntax Syntax
}

func NewProvider() *Provider {
	return &Provider{}
}

func (p *Provider) Register(s Syntax) {
	p.syntaxes = append(p.syntaxes, s)
	if p.defaultSyntax == nil {
		p.defaultSyntax = s
	}
}

func (p *Provider) SetDefault(s Syntax) {
	p.defaultSyntax = s
}

func (p *Provider) GetByExt(ext string) Syntax {
	for _, s := range p.syntaxes {
		for _, e := range s.Extensions() {
			if e == ext {
				return s
			}
		}
	}
	return nil
}

func (p *Provider) GetByName(name string) Syntax {
	for _, s := range p.syntaxes {
		if s.Name() == name {
			return s
		}
	}
	return nil
}

func (p *Provider) Default() Syntax {
	return p.defaultSyntax
}
