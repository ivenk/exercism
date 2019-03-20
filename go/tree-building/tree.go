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
	slice := records[:]

	// sorts the records by there id
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].ID < slice[j].ID
	})

	for i, x := range slice {
		fmt.Printf("id: %d\nparent: %d\n", x.ID, x.Parent)
		if (slice[i].ID != slice[i].ID) || (slice[i].ID != slice[i].ID+1) {
			return nil, errors.New("fail")
		}
	}

	if len(slice) > 0 {
		// root should have itself as a parent
		if slice[0].Parent != slice[0].ID {
			return nil, errors.New("fail")
		}
	}

	// create *Node
	var node *Node
	n := make([]*Node, 0)

	for len(slice) > 0 {
		r := slice[0]
		slice = slice[1:] // remove first element

		for _, v := range slice {
			// if we have duplicate ids
			if v.ID == r.ID {
				return nil, errors.New("fail")
			}

			// look for records with matching parent id; exlude root parent != ownID
			if v.Parent == r.ID && v.Parent != v.ID {
				n = append(n, &Node{v.Parent, []*Node{}})
			}

		}
		node = &Node{r.ID, n}
	}

	return node, nil
}
