// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	var trimmedRemark = strings.TrimSpace(remark)
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	var isQuestion = strings.HasSuffix(trimmedRemark, "?")
	var isShouting = trimmedRemark == strings.ToUpper(trimmedRemark)
	var hasLetters, _ = regexp.MatchString(`[a-zA-Z]`, trimmedRemark)

	if trimmedRemark == "" {
		return "Fine. Be that way!"
	}

	if isQuestion && isShouting && hasLetters {
		return "Calm down, I know what I'm doing!"
	}

	if isQuestion {
		return "Sure."
	}

	if !hasLetters {
		return "Whatever."
	}

	if isShouting {
		return "Whoa, chill out!"
	}

	return "Whatever."
}
