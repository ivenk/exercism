package tournament

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

const BufferSize = 1000

const TableHeader = "Team                           | MP |  W |  D |  L |  P"

type entry struct {
	name   string
	wins   int
	losses int
	draws  int
}

// input looks like :
// ---
// Allegoric Alaskians;Blithering Badgers;win
// Devastating Donkeys;Courageous Californians;draw
// ---
// gets converted to:
// internal map structure: [name_of_the_team] = [wins, losses]
// Tally converts the given input string into a table and returns it
func Tally(in io.Reader, out io.Writer) error {
	m := make(map[string]entry)

	buffer := make([]byte, BufferSize)
	n, e := in.Read(buffer)
	if e != nil {
		return errors.New("Could not read input")
	}
	entries := string(buffer[:n])

	// iterate over all lines
	for _, l := range strings.Split(entries, "\n") {
		// there are some empty lines in the input
		if l == "" || string(l[0]) == "#" {
			continue
		}

		w := strings.Split(l, ";")
		if len(w) != 3 {
			return errors.New("Input data is formated incorrectly")
		}

		// make sure both teams are already in our map
		for i := 0; i < 2; i++ {
			if _, b := m[w[0]]; !b { // if team does not yet exist
				m[w[0]] = entry{w[0], 0, 0, 0} // create entry
			}
		}

		// add points to map
		switch w[2] {
		case "win":
			m[w[0]] = entry{w[0], m[w[0]].wins + 1, m[w[0]].losses, m[w[0]].draws} // add one to wins for team 1
			m[w[1]] = entry{w[1], m[w[1]].wins, m[w[1]].losses + 1, m[w[1]].draws} // second team gets a loss
		case "loss":
			m[w[0]] = entry{w[0], m[w[0]].wins, m[w[0]].losses + 1, m[w[0]].draws} // add one to loss for team 1
			m[w[1]] = entry{w[1], m[w[1]].wins + 1, m[w[1]].losses, m[w[1]].draws} // second team gets a win
		case "draw": // do nothing
			m[w[0]] = entry{w[0], m[w[0]].wins, m[w[0]].losses, m[w[0]].draws + 1} // add one to loss for team 1
			m[w[1]] = entry{w[1], m[w[1]].wins, m[w[1]].losses, m[w[1]].draws + 1} // second team gets a win
		default:
			return errors.New("game result could not be parsed for " + w[0] + " : " + w[1])
		}
	}

	// build output
	//	separatorIndex := strings.Index(TableHeader, "|")
	out.Write([]byte(TableHeader))
	out.Write([]byte("\n"))

	numberOfTeams := len(m)

	for i := 0; i < numberOfTeams; i++ {
		bestIndex := ""
		// search highest scoring team
		for k, v := range m {
			// first iteration
			if bestIndex == "" {
				bestIndex = k
				continue
			}

			if v.wins > m[bestIndex].wins {
				bestIndex = k
			} else if v.wins == m[bestIndex].wins {
				if v.losses < m[bestIndex].losses {
					bestIndex = k
				}
			}
		}

		bk := m[bestIndex]

		// pads each team name to 25 characters, added after the name (-)
		totalPoints := 3*bk.wins + bk.draws
		out.Write([]byte(fmt.Sprintf("%-31v|  %v |  %v |  %v |  %v |  %v\n", bk.name, (bk.wins + bk.losses + bk.draws), bk.wins, bk.draws, bk.losses, totalPoints)))
		delete(m, bestIndex)
	}

	return nil
}
