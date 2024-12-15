package furniture

import (
	"fmt"
	"strings"
)

type Furniture []*Item

func (fur Furniture) IsEmpty() bool {
	for _, item := range fur {
		if len(item.Contents) > 0 {
			return false
		}
	}
	return true
}

func (fur Furniture) ShowContents() (contents string) {
	var b strings.Builder
	furContents := make([]string, 0, len(fur))
	for _, item := range fur {
		if len(item.Contents) == 0 {
			continue
		}
		fmt.Fprintf(&b, "%s: %s", item.Name, strings.Join(item.IterationOrder, ", "))
		furContents = append(furContents, b.String())
		b.Reset()
	}
	return strings.Join(furContents, ", ")
}
