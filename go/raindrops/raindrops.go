package raindrops

import "fmt"

func Convert(input int) string {
	// the factors we are interested in
	var factors = []int{ 3, 5, 7 }

	result := ""
	for _, factor := range factors {
		// do we have a clean divide ?
		if input%factor == 0 {
			switch factor {
				case 3:
					result += "Pling"
				case 5:
					result += "Plang"
				case 7:
					result += "Plong"
			}
		}
	}

	// do we have any matches ?
	if result == "" {
		result = fmt.Sprintf("%d", input)
	}

	return result
}


