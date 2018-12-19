package luhn

import (
	"strings"
	"strconv"
)



func Valid(number string) bool{
	var double = false

	// we remove the spaces
	trimedNumber := strings.Replace(number, " ", "", -1)

	// length 1 or below is always false
	if len(trimedNumber) <= 1 {
		return false
	}

	// if we can parse it, it should not contain any special or alphabetical characters
	// this might fail with . or ,
	if _, err := strconv.Atoi(trimedNumber); err != nil {
		return false
	}

	var sum = 0
	//iterate in reverse
	for i := len(trimedNumber) -1; i >= 0; i-- {
		var helper, _  =  strconv.Atoi(string(trimedNumber[i]))

		// we double every second number
		if double {
			double = false

			// now we multiply
			helper = helper * 2

			//subtract 9 if the number is bigger then 9
			if helper > 9 {
				helper = helper - 9
			}
		} else {
			double = true
		}

		sum += helper
	}


	// can we evenly divide by 10
	if sum % 10 == 0 {
		return true
 	}

	return false
}
