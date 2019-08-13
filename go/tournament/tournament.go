package tournament

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

const bufferSize int = 1000

// map structure: [name_of_the_team] = [wins, losses]

// Tally converts the given input string into a table and returns it
func Tally(reader io.Reader, buffer io.Writer) error {
	m := make(map[string][2]int)
	resultbytes := make([]byte, bufferSize) // dont know the size; but we want to read it all

	for hasMore := true; hasMore; {
		i, _ := reader.Read(resultbytes) // this is really redundant but i cant think of a bette way ...
		if i > 0 {

		}
	}

	fmt.Printf("%d\n", i)
	results := string(resultbytes)

	for _, l := range strings.Split(results, "\n") { // split on new lines
		w := strings.Split(l, ";")
		v1, _ := strconv.Atoi(w[1])
		v2, _ := strconv.Atoi(w[2])
		if v, b := m[w[0]]; b {
			m[w[0]] = [...]int{v[0] + v1, v[1] + v2}
		} else {
			// add entry with the described format
			m[w[0]] = [...]int{v1, v2}
		}
	}
	return nil
}
