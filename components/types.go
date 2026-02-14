package components

import "github.com/a-h/templ"

const (
	btnClassRed     = "btn btn-danger btn-lg active"
	btnClassPrimary = "btn btn-primary btn-lg active"
)

type Button struct {
	Href     templ.SafeURL
	cssClass string
	Text     string
	disabled bool
}
type Header struct {
	Title    string
	Subtitle string
}

type Body struct {
	Paragraph string
}
type Image struct {
	SrcURL templ.SafeURL
	Alt    string
}
