package strain

type Ints []int

type Lists [][]int

type Strings []string

func (ints Ints) Keep(operation func (int) bool) Ints {
	// nil check
	if ints == nil  {
		return nil
	}

	var newInts = Ints{}
	for _,x := range ints {
		if (operation(x)) {
			newInts = append(newInts, x)
		}
	}

	return newInts
}


func (ints Ints) Discard(operation func (int) bool) Ints {
	// nil check
	if ints == nil  {
		return nil
	}

	var newInts = Ints{}
	for _,x := range ints {
		if !operation(x) {
			newInts = append(newInts, x)
		}
	}
	return newInts
}


func (lists Lists) Keep(operation func([]int) bool) Lists {
	// nil check
	if lists == nil  {
		return nil
	}

	var newLists = Lists{}
	for _, xList := range lists {
		if operation(xList) {
			newLists = append(newLists, xList)
		}
	}

	return newLists
}

func (str Strings) Keep(operation func(string) bool) Strings {
	var newStrings = Strings{}
	for _,x := range str {
		if operation (x) {
			newStrings = append(newStrings, x)
		}
	}
	return newStrings
}

