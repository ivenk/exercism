package tree

import (
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
	slice := records[:]

	// sorts the records by there id
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].ID < slice[j].ID
	})

	for _, x := range slice {
		fmt.Printf("id: %d\n", x.ID)
	}

	// create *Node
	var node *Node
	n := make([]*Node, 0)

	for len(slice) > 0 {
		r := slice[0]
		for _, v := range slice {
			if v.Parent == r.ID {
				n = append(n, &Node{v.Parent, []*Node{}})
			}
		}
		node = &Node{r.ID, n}
		slice = slice[1:] // remove first element

	}
	return node, nil
}
