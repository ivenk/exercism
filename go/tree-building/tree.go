package tree

import (
	"errors"
	"fmt"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	// error checks
	if len(records) == 0 {
		return nil, nil
	}

	lastElements := 0
	// check for duplicates and non continuous
	for _, i := range records {
		occurence := 0
		hasSuccessor := false

		for _, j := range records {
			if i.ID == j.ID {
				occurence += 1
			}
			if i.ID == (j.ID + 1) {
				hasSuccessor = true
			}
		}

		if occurence > 1 {
			return nil, errors.New("fail")
		}

		// if it does not have a direct successor it must be the last element
		if !hasSuccessor {
			lastElements += 1
		}
	}

	// if there are more then one elments without successors we have non continiuous items
	if lastElements > 1 {
		return nil, errors.New("fail")
	}

	fmt.Println("Records:")
	workingParent := false
	for _, x := range records {
		fmt.Print("      ")
		fmt.Println(x)
		if (x.ID == x.Parent) && (x.ID == 0) { // node with id 0 and id = parentID
			workingParent = true
		}
	}
	if !workingParent {
		return nil, errors.New("Fail !")
	}

	slice := records[:]
	return buildNode(0, slice), nil
}

// builds a node with the given id
// browses the slice for children nods
// slice has to be passed by value !
func buildNode(id int, slice []Record) *Node {
	var childIDs []int
	var children []*Node

	var cutSlice []Record
	ownIndex := -1
	// assemble children
	for i, x := range slice {
		if x.Parent == id {
			if x.ID != id { // only add if not self
				if !contains(childIDs, x.ID) { // this check should not be requiered...
					childIDs = append(childIDs, x.ID)
				}
			}
		}
		// find own position
		if x.ID == id {
			ownIndex = i
		}
	}
	// only remove ourself if we found ourself
	if ownIndex > -1 {
		cutSlice = remove(slice, ownIndex)
	}

	fmt.Printf("childRecords for %d\n", id)
	fmt.Println(childIDs)

	// sorts the records by there id
	sort.Ints(childIDs)

	//  build rest ...
	for _, record := range childIDs {
		child := buildNode(record, cutSlice)
		children = append(children, child)
	}

	return &Node{id, children}
}

// Removes element of slice
// taken from ...
func remove(slice []Record, index int) []Record {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

// checks wether the int value is found within the given slice
func contains(slice []int, value int) bool {
	for _, x := range slice {
		if x == value {
			return true
		}
	}

	return false
}
