package components

import "github.com/a-h/templ"

type profileData interface {
	ProfTitle() string
	ProfSubTitle() string
	ProfDesc() string
	ProfImg() string
}

type IconData interface {
	GithubURL() templ.SafeURL
	LinkedinURL() templ.SafeURL
	EmailURL() templ.SafeURL
	CellURL() templ.SafeURL
	ProjectURL() templ.SafeURL
}
type data struct {
	Title       string        `json:"title"`
	SubTitle    string        `json:"subtitle"`
	Description string        `json:"description"`
	Image       string        `json:"image"`
	Github      templ.SafeURL `json:"github"`
	Linkedin    templ.SafeURL `json:"linkedin"`
	Email       templ.SafeURL `json:"email"`
	Cell        templ.SafeURL `json:"cell"`
	Project     templ.SafeURL `json:"project"`
}
