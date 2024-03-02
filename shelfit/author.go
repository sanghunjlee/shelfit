package shelfit

import (
	"errors"
	"regexp"
	"strings"
)

type Author struct{}

func (a *Author) Write(text string, note string) (*Content, error) {
	if len(strings.Replace(text, " ", "", -1)) == 0 {
		return nil, errors.New("empty input: there is no args to parse from")
	}
	content := &Content{
		HasArchived: false,
		HasCategory: false,
		HasTag:      false,
		HasVolumes:  false,
		HasStatus:   false,
		HasLink:     false,
	}

	var title []string
	var re *regexp.Regexp

	for _, word := range strings.Split(text, " ") {
		match := false
		// Check for Category (qualified by `!`)
		re, _ = regexp.Compile(`^\![\p{L}\d_-]+`)
		if re.MatchString(word) {
			if !content.HasCategory {
				content.HasCategory = true
				content.Category = word[1:]
			} else {
				return nil, errors.New("multiple category: only one category is supported at the moment")
			}
			match = true
		}
		// Check for Tag (qualified by `.`)
		re, _ = regexp.Compile(`^\.[\p{L}\d_-]+`)
		if re.MatchString(word) {
			content.HasTag = true
			content.Tags = append(content.Tags, word[1:])
			match = true
		}

		if !match {
			title = append(title, word)
		}
	}

	content.Title = strings.Join(title, " ")
	content.Note = note
	if !content.HasCategory {
		return nil, errors.New("no category: there is no category (required)")
	}

	return content, nil
}
