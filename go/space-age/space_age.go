package space

// The tests require a Planet type but its just a fancy string
type Planet string

var planetInfo = map[Planet] float64 {
	// earth is the default
	"Earth": 1,
	// all other planets have modifiers
	"Mercury": 0.2408467,
	"Venus": 0.61519726,
	"Mars": 1.8808158,
	"Jupiter": 11.862615,
	"Saturn": 29.447498,
	"Uranus": 84.016846,
	"Neptune": 164.79132,
}

func Age(seconds float64, planet Planet) float64{
	//multiply the seconds per year with the planet modifier and then defied the given amount of seconds through our planet specific year duration in seconds
	//we then get the age
	return seconds/ (31557600 * planetInfo[planet])
}