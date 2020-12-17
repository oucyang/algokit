package union

type Node struct {
	Value interface{}
	Id    int64
}

type UnionSet struct {
	Nodes   map[int64]*Node
	Parents map[int64]*Node
	Sizes   map[int64]int
	V       map[interface{}]int64
}

var lastId int64 = 0

func getId() int64 {
	id := lastId
	lastId++
	return id
}

func NewUnionSet(values []interface{}) *UnionSet {
	var us = new(UnionSet)
	us.Nodes = make(map[int64]*Node, 0)
	us.Parents = make(map[int64]*Node, 0)
	us.Sizes = make(map[int64]int, 0)
	for _, v := range values {
		n := &Node{Value: v, Id: getId()}
		us.Nodes[n.Id] = n
		us.Parents[n.Id] = n
		us.Sizes[n.Id] = 1
	}
	return us
}
