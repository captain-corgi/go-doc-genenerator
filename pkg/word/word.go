package word

// DocTemplate is an interface for doc format functionality
type DocTemplate interface {
	// OpenTemplate open existing doc template for render later
	OpenTemplate(path string) (DocTemplate, error)
	// RenderTemplate render data using given template
	RenderTemplate(data interface{}) error
	// Save rendered data to file
	Save(filePath string) error
	// ResetTemplate re-open template, if it cannot reset by calling RenderTemplate again
	ResetTemplate() error
	// CloseTemplate close opening template
	CloseTemplate() error
}

// DocxTemplate is an interface for docx format functionality
type DocxTemplate interface {
	// OpenTemplate open existing docx template for render later
	OpenTemplate(path string) (DocxTemplate, error)
	// RenderTemplate render data using given template
	RenderTemplate(data interface{}) error
	// Save rendered data to file
	Save(filePath string) error
	// ResetTemplate re-open template, if it cannot reset by calling RenderTemplate again
	ResetTemplate() error
	// CloseTemplate close opening template
	CloseTemplate() error
}

// NewDocLegionZverTemplate create doc template with LegionZverDoc library
func NewDocLegionZverTemplate() DocTemplate {
	return &LegionZverDoc{}
}

// NewDocxLegionZverTemplate create docx template with LegionZverDocx library
func NewDocxLegionZverTemplate() DocxTemplate {
	return &LegionZverDocx{}
}

// NewDocxTemplate create docx template with SimpleDocx library
func NewDocxTemplate() DocxTemplate {
	return &SimpleDocx{}
}
