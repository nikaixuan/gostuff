package binaryTree

// 341
type NestedInteger struct{}

func (this NestedInteger) IsInteger() bool           { return false }
func (this NestedInteger) GetInteger() int           { return 0 }
func (n *NestedInteger) SetInteger(value int)        {}
func (this *NestedInteger) Add(elem NestedInteger)   {}
func (this NestedInteger) GetList() []*NestedInteger { return []*NestedInteger{} }

type NestedIterator struct {
	value []int
	index int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	res := make([]int, 0)

	for _, l := range nestedList {
		traverseIterator(l, &res)
	}
	return &NestedIterator{value: res, index: 0}
}

func traverseIterator(nestedIt *NestedInteger, r *[]int) {
	if nestedIt.IsInteger() {
		*r = append(*r, nestedIt.GetInteger())
	}
	for _, l := range nestedIt.GetList() {
		traverseIterator(l, r)
	}
}

func (this *NestedIterator) Next() int {
	next := this.value[this.index]
	this.index++
	return next
}

func (this *NestedIterator) HasNext() bool {
	return this.index < len(this.value)
}
