// mostly not my stuff
// looked at the amazing solution of luisrpp here on exercism.

package tree

import (
	"errors"
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
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	// create a slice of nodes
	nodes := make([]Node, len(records))

	for i, x := range records {
		// error checking
		if i != x.ID {
			return nil, errors.New("ids not continued")
		}
		if x.ID != 0 && x.ID <= x.Parent {
			return nil, errors.New("id has to be bigger then id of parent")
		}
		if x.ID == 0 && x.Parent != 0 {
			return nil, errors.New("root cannot have a parent")
		}
		// --------------

		nodes[x.ID] = Node{ID: x.ID}

		if i != 0 {
			// get reference from parent, Parent will already be created due to sorting and the lower indexed parents
			parent := &nodes[x.Parent]
			parent.Children = append(parent.Children, &nodes[i]) // append the node
		}
	}

	return &nodes[0], nil
}
