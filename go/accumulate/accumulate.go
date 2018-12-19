package accumulate

func Accumulate(collection []string, operation func (string) string) [] string {
	var newCollection = []string{}

	// the iteration with increasing indizes preserves the order
	for i:= 0; i< len(collection); i++ {
		var convertedValue = operation(collection[i])
		newCollection = append(newCollection, convertedValue)
	}

	return newCollection
}
