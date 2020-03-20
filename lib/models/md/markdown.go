package md

import (
	"github.com/iancoleman/strcase"
	"io/ioutil"
	"strings"
)

// md is a md file
type markdown struct {
	location string
	content string
	parsed bool
	base string
}

type Type interface {
	Content() string
	URL() string
}

// Content of the md file
func (md *markdown) Content() string {
	if md.parsed == false {
		md.parsed = true
		if content, err := ioutil.ReadFile(md.location); err == nil {
			md.content = string(content)
		}
	}
	return md.content
}

// URL path of the md file
func (md *markdown) URL() string {
	var result string
	result = strings.ReplaceAll(md.location, md.base + "/", "")
	result = strings.TrimSuffix(result, ".md")
	result = strcase.ToKebab(result)
	result = strings.ReplaceAll(result, "/-", "/")
	return result
}

// New returns a new md file
func New(path, base string) Type {
	return &markdown{
		location: path,
		base: base,
	}
}
