package robotname

// Want to hide the name from unchecked access
// Robot struct representing a robot with a unique name
type Robot struct {
	name string
}

// global nameslice
var nameList []string

// Name returns the unique robot name, if the robot has no name yet a new one will be created
func (r Robot) Name() (string, error) {
	if r.name != "" {
		return name, nil
	}

	name := generateName()
	r.name = name
	return name
}

// generates a random name for the robot which is not yet in the namelist. Adds it to the list after generation
func generateName() string {
	newName := ""
	for contains(nameList, newName) || (newName == "") {
		newName = fmt.Sprintf("%s%s%03d", ('A' + rand.Intn(24)), ('A' + rand.Intn(24)), rand.Intn(999))
	}
	nameList = append(nameList, newName)
	return newName
}

// contains checks wether the slice of strings contains the given key
func contains(slice []string, key string) bool {
	for _, x := range slice {
		if x == key {
			return true
		}
	}
	return false
}
