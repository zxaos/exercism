// Package space provides utilities for handling dates on different planets.
package space

type Planet string

const EARTH_YEAR_SECONDS = 31557600

var orbitalPeriod = map[Planet]float64{
	"Earth":   EARTH_YEAR_SECONDS,
	"Mercury": 0.2408467 * EARTH_YEAR_SECONDS,
	"Venus":   0.61519726 * EARTH_YEAR_SECONDS,
	"Mars":    1.8808158 * EARTH_YEAR_SECONDS,
	"Jupiter": 11.862615 * EARTH_YEAR_SECONDS,
	"Saturn":  29.447498 * EARTH_YEAR_SECONDS,
	"Uranus":  84.016846 * EARTH_YEAR_SECONDS,
	"Neptune": 164.79132 * EARTH_YEAR_SECONDS,
}

// Age returns an age in "planet years" given a planet and an age in seconds.
func Age(seconds float64, planet Planet) float64 {
	return seconds / orbitalPeriod[planet]
}
