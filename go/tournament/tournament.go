package tournament

import (
	"errors"
	"io"
	"strings"
)

const BufferSize = 1000

const TableHeader = "Team                           | MP |  W |  D |  L |  P"

// input looks like :
// ---
// Allegoric Alaskians;Blithering Badgers;win
// Devastating Donkeys;Courageous Californians;draw
// ---
// gets converted to:
// internal map structure: [name_of_the_team] = [wins, losses]
// Tally converts the given input string into a table and returns it
func Tally(in io.Reader, out io.Writer) error {
	m := make(map[string][2]int)

	buffer := make([]byte, BufferSize)
	n, e := in.Read(buffer)
	if e != nil {
		return errors.New("Could not read input")
	}
	entries := string(buffer[:n])

	// iterate over all lines
	for _, l := range strings.Split(entries, "\n") {
		// there are some empty lines in the input
		if l == "" {
			continue
		}

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
		case "loss":
			m[w[0]] = [2]int{m[w[0]][0], m[w[0]][1] + 1} // add one to loss for team 1
			m[w[1]] = [2]int{m[w[1]][0] + 1, m[w[1]][1]} // second team gets a win
		case "draw": // do nothing
			break
		default:
			return errors.New("game result could not be parsed for " + w[0] + " : " + w[1])
		}
	}

	// build output
	return nil
}
