package components

import "github.com/a-h/templ"

type Button struct {
	Href templ.SafeURL
	Text string
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
