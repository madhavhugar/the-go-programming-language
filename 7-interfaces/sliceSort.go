package interfaces

type integers []int

func (ins integers) Len() int {
	return len(ins)
}

func (ins integers) Less(i, j int) bool {
	return ins[i] < ins[j]
}

func (ins integers) Swap(i, j int) {
	ins[i], ins[j] = ins[j], ins[i]
}
