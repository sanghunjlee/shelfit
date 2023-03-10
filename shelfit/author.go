package shelfit

import (
	"regexp"
	"strings"
)

type Author struct{}

func (a *Author) Write(text string) (*Content, error) {
	content := &Content{
		HasArchived: false,
		HasCategory: false,
		HasGenres:   false,
		HasVolumes:  false,
		HasStatus:   false,
		HasLink:     false,
	}

	var title []string
	var re *regexp.Regexp

	for _, word := range strings.Split(text, " ") {
		match := false
		// Check for Category (quantified by "@")
		re, _ = regexp.Compile(`^\@[\p{L}\d_-]+`)
		if re.MatchString(word) {
			if !content.HasCategory {
				content.HasCategory = true
				content.Category = word[1:]
			}
			match = true
		}
		// Check for Genre (quantified by "!")
		re, _ = regexp.Compile(`^\.[\p{L}\d_-]+`)
		if re.MatchString(word) {
			content.HasGenres = true
			content.Genres = append(content.Genres, word[1:])
			match = true
		}
		// Check for Volumes (quantified by "+")
		re, _ = regexp.Compile(`^\+\p{L}*\d+.*$`)
		if re.MatchString(word) {
			content.HasVolumes = true
			content.VolumeTitles = append(content.VolumeTitles, word[1:])
			match = true
		}
		// Check for Status (quantified by "-")
		re, _ = regexp.Compile(`^\!\p{L}+.*$`)
		if re.MatchString(word) {
			content.Status = append(content.Status, a.parseStatus(word[1:]))
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

		if !match {
			title = append(title, word)
		}
	}

	content.Title = strings.Join(title, " ")

	if !content.HasStatus {
		content.Status = append(content.Status, Unread)
	}

	return content, nil
}

func (a *Author) parseStatus(text string) Status {
	if contain(text, Finished.RelatedStrings()) {
		return Finished
	} else if contain(text, Started.RelatedStrings()) {
		return Started
	} else {
		return Unread
	}
}
