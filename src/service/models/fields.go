package models

type Field struct {
	Name       string
	Path       string
	ArrayPath  string
	Type       string
	ActionType string
	Value      string
	Error      error
}

type ExtractedFields struct {
	Fields []Field
}
