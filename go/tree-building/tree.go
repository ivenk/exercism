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

// special type sorting
type ByID []Record

func (rec ByID) Len() int {
	return len(rec)
}

func (rec ByID) Swap(i, j int) {
	rec[i], rec[j] = rec[j], rec[i]
}

func (rec ByID) Less(i, j int) bool {
	return rec[i].ID < rec[j].ID
}

func Build(records []Record) (*Node, error) {
	// error checks
	if len(records) == 0 {
		return nil, nil
	}
	if len(records) == 1 {
		return &Node{0, nil}, nil
	}
	workingParent := false
	for _, x := range records {
		fmt.Println("Records:")
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
	var childRecords []Record
	var children []*Node

	var cutSlice []Record
	ownIndex := -1
	// assemble children
	for i, x := range slice {
		if x.Parent == id {
			if x.ID != id { // only add if not self
				childRecords = append(childRecords, x)
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

	// sorts the records by there id
	sort.Sort(ByID(childRecords))

	//  build rest ...
	for _, record := range childRecords {
		child := buildNode(record.ID, cutSlice)
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
