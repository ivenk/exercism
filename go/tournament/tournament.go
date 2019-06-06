package tournament

import (
	"strconv"
	"strings"
)

// map structure: [name_of_the_team] = [wins, losses]

// Tally converts the given input string into a table and returns it
func Tally(results string) string {
	m := make(map[string][2]int)

	for _, l := range strings.Split(results, "\n") { // split on new lines
		w := strings.Split(l, ";")
		if v, b := m[w[0]]; b {
			m[w[0]] = [...]int{v[0] + strconv.Atoi(w[1]), v[1] + strconv.Atoi(w[2])}
		} else {
			// add entry with the described format
			m[w[0]] = [...]int{w[1], w[2]}
		}
	}

	return ""
}
