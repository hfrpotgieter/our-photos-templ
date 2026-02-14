package components

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/a-h/templ"
)

func (d data) ProfTitle() string {
	return d.Title
}
func (d data) ProfSubTitle() string {
	return d.SubTitle
}
func (d data) ProfDesc() string {
	return d.Description
}
func (d data) ProfImg() string {
	return d.Image
}

func (d data) GithubURL() templ.SafeURL {
	return d.Github
}
func (d data) LinkedinURL() templ.SafeURL {
	return d.Linkedin
}
func (d data) EmailURL() templ.SafeURL {
	return d.Email
}
func (d data) CellURL() templ.SafeURL {
	return d.Cell
}
func (d data) ProjectURL() templ.SafeURL {
	return d.Project
}

const (
	ProfileDataPath = "data/profile_info.json"
	IconsDataPath   = "data/icons_info.json"
)

func ReadJsonFile(path string) data {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("could not get current working directory %v", err)
	}
	b, err := os.ReadFile(strings.Join([]string{wd, path}, `/`))
	if err != nil {
		log.Fatalf("could not open json file %s, %v", path, err)
	}
	var d data
	err = json.Unmarshal(b, &d)
	if err != nil {
		log.Fatalf("could not unmarshal json data %s", err)
	}
	return d
}
