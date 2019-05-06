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
var used = make(map[string]bool)
var maxRobots = 26 * 26 * 10 * 10 * 10

// Name returns the unique robot name, if the robot has no name yet a new one will be created
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	if len(used) >= maxRobots {
		return "", errors.New("Maximum number of available robots reached")
	}

	r.name = generateUniqueName()
	return r.name, nil
}

// Reset resets the current robot name
func (r *Robot) Reset() { r.name = "" }

// generates a random name for the robot which is not yet in the namelist. Adds it to the list after generation
func generateUniqueName() string {
	newName := generateRandomName()
	for used[newName] {
		newName = generateRandomName()
	}
	used[newName] = true
	return newName
}

func generateRandomName() string {
	return fmt.Sprintf("%s%s%03d", string(('A' + rand.Intn(26))), string(('A' + rand.Intn(26))), rand.Intn(1000))
}
