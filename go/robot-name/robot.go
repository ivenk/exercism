package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

// Robot struct representing a robot with a unique name
// Want to hide the name from unchecked access
type Robot struct {
	name string
}

// global nameslice
var nameList []string
var maxRobots = 26 * 26 * 10 * 10 * 10

// Name returns the unique robot name, if the robot has no name yet a new one will be created
func (r *Robot) Name() (string, error) {
	//	fmt.Println(nameList)
	if r.name != "" {
		return r.name, nil
	}

	if len(nameList) >= maxRobots {
		return "", errors.New("Maximum number of available robots reached")
	}

	name := generateName()
	r.name = name
	return name, nil
}

// Reset resets the current robot name
func (r *Robot) Reset() { r.name = "" }

// generates a random name for the robot which is not yet in the namelist. Adds it to the list after generation
func generateName() string {
	newName := ""
	for {
		newName = fmt.Sprintf("%s%s%03d", string(('A' + rand.Intn(26))), string(('A' + rand.Intn(26))), rand.Intn(1000))
		if _, x := contains(nameList, newName); !x {
			nameList = append(nameList, newName)
			return newName
		}
	}
}

// contains checks wether the slice of strings contains the given key
func contains(slice []string, key string) (int, bool) {
	for i, x := range slice {
		if x == key {
			return i, true
		}
	}
	return -1, false
}
