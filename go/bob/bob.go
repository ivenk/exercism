package bob

import (
	"strings"
	"unicode"
)

func Hey(remark string) string {
	//check if remark is empty
	if strings.TrimSpace(remark) == "" {
		return 	"Fine. Be that way!"
	}
	//make the remark upper
	var testUpper = strings.ToUpper(remark)
	//test whether that changed anything
	if testUpper == remark {
		var yelling = false;
		for _, r := range remark {
			if unicode.IsLetter(r) {
				yelling = true;
			}
		}
		if yelling {
			//are we yelling a question
			if strings.Contains(remark, "?") {
				return "Calm down, I know what I'm doing!"
			} else {
				return "Whoa, chill out!"
			}
		}
	}

	//no yelling just a question ?
	if strings.HasSuffix(strings.TrimSpace(remark), "?") {
		return "Sure."
	}

	return "Whatever."
}
