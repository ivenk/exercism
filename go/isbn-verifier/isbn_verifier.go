package isbn

import (
  "strings"
  "strconv"
)

func IsValidISBN(number string) bool {
  // first we clean the input
  number = strings.TrimSpace(number)
  cleanedInput := strings.Replace(number, "-", "", -1)

  // first we check odd cases
  // we trimed spaces and "-" so it should not be larger then 10
  if len(cleanedInput) != 10 {
    return false
  }

  var factor = 10
  var sum = 0

  // for every letter
  for i := 0; i < len(cleanedInput); i++ {
    var sn int
    var err error

    // if the last letter is X
    if i == (len(cleanedInput) - 1) && cleanedInput[i] == "X"[0] {
      sn = 10
    } else {
      // convert to int number
      sn, err = strconv.Atoi(string(cleanedInput[i]))

      // if we have something non numerical in our first 9 letters
      if err != nil {
        return false
      }
    }

    // multiply current number with factor 10-1 and add to total sum
    sum = sum + (sn * factor)
    // decrease the factor for the next iteration
    factor = factor - 1
  }

  // finally we check for validity
  return (sum % 11) == 0

}
