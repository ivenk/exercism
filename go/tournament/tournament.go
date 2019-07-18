package tournament

import (
	"errors"
	"io"
	"strconv"
	"strings"
)

// input looks like :
// ---
// Allegoric Alaskians;Blithering Badgers;win
// Devastating Donkeys;Courageous Californians;draw
// ---
// gets converted to:
// internal map structure: [name_of_the_team] = [wins, losses]
// Tally converts the given input string into a table and returns it
func Tally(io.Reader, io.Writer) error {
	var results string
	m := make(map[string][2]int)

	for _, l := range strings.Split(results, "\n") { // split on new lines
		w := strings.Split(l, ";")
		if len(w) != 3 {
			return errors.New("Input data is formated incorrectly")
		}

		// make sure both teams are already in our map
		for i := 0; i < 2; i++ {
			if _, b := m[w[0]]; !b { // if team does not yet exist
				m[w[0]] = [2]int{0, 0} // create entry
			}
		}

		// add points to map
		switch w[2] {
		case "win":
			m[w[0]] = [2]int{m[w[0]][0] + 1, m[w[0]][1]} // add one to wins for team 1
			m[w[1]] = [2]int{m[w[1]][0], m[w[1]][1] + 1} // second team gets a loss
		case "draw": // do nothing
			break
		case "loss":
			m[w[0]] = [2]int{m[w[0]][0], m[w[0]][1] + 1} // add one to loss for team 1
			m[w[1]] = [2]int{m[w[1]][0] + 1, m[w[1]][1]} // second team gets a win
		default:
			return errors.New("game result could not be parsed for " + w[0] + " : " + w[1])
		}
	}

	return ""
}
