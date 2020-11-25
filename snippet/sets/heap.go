package sets

import "sort"

type HeapInterface interface {
	sort.Interface
	Append(v interface{})
}
