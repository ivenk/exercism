package list_ops

type IntList []int

type BinFunc func(x ,y int) int


func (list IntList) Foldr(fn BinFunc, initial int) int {

	for _, value := range list {
		initial += value
	}

	return initial
}




func (list IntList) Foldl(fn BinFunc, initial int) int {


	return initial
}
