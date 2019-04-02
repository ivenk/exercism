package tree

import (
//	"errors"
//	"fmt"
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
	if len(records) == 0 {
		return nil, nil
	}
	if len(records) == 1 {
		return &Node{0, nil}, nil
	}
	slice := records[:]

	// need to test for failures

	return buildNode(0, slice), nil
}

// builds a node with the given id
// browses the slice for children nods
// slice has to be passed by value !
func buildNode(id int, slice []Record) *Node {
	var childRecords []Record
	var children []*Node

	var ownIndex int
	// assemble children
	for i, x := range slice {
		if x.Parent == id {
			cutSlice := remove(slice, i)
			childRecords = append(childRecords, x)
		}
		// find own position
		if x.ID == id {
			ownIndex = i
		}
	}
	cutSlice = remove(cutSlice, ownIndex)

	//  build rest ...
	for _, record : = childRecords {
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
