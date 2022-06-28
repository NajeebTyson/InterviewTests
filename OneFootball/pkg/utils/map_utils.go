package utils

// GetMapValues is a generic function which returns all the values in the map as a slice
func GetMapValues[K comparable, V any](aMap map[K]V) []V {
	values := make([]V, 0, len(aMap))
	for _, value := range aMap {
		values = append(values, value)
	}
	return values
}
