package shelfit

import (
	"regexp"
	"strings"
)

type Author struct{}

func (a *Author) Write(text string) (*Content, error) {
	content := &Content{
		HasArchived: false,
		HasFinished: false,
		HasGenres:   false,
		HasVolumes:  false,
		HasStatus:   false,
		HasLink:     false,
	}

	re_category, _ := regexp.Compile(`\@[\p{L}\d_-]+`)
	re_genres, _ := regexp.Compile(`\![\p{L}\d_-]+`)

	content.Category = a.matchWords(text, re_category)[0]
	content.Genres = a.matchWords(text, re_genres)

	var re *regexp.Regexp

	for _, word := range strings.Split(text, " ") {
		match := false

		// Check for Volumes (quantified by "+")
		re, _ = regexp.Compile(`^\+\p{L}*\d+.*$`)
		if re.MatchString(word) {
			content.HasVolumes = true
			content.Volumes = append(content.Volumes, word)
			match = true
		}
		// Check for Status (quantified by "-")
		re, _ = regexp.Compile(`^\-\p{L}+.*$`)
		if re.MatchString(word) {
			content.Status = append(content.Status, a.parseStatus(word))
			if content.HasVolumes {
				if !content.HasStatus {
					content.Status = append(content.Status, content.PrevStatus())
				}
				if content.Status[0] < content.PrevStatus() {
					content.Status[0] = Started
				}
			}
			content.HasStatus = true
			match = true
		}

	}
}

func (a *Author) matchWords(text string, re *regexp.Regexp) []string {
	results := re.FindAllString(text, -1)
	ret := []string{}

	for _, val := range results {
		ret = append(ret, val[1:])
	}
	return ret
}

func (a *Author) parseStatus(text string) Status {

}
