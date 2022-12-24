package solver

type DataNode struct {
	ExpectedScore int64
	Score         int64
	MaxScore      int64
}

type OperatorNode struct {
	Operator string
}

type Node struct {
	Data     *DataNode
	Operator *OperatorNode
	RawData  string
	Left     *Node
	Right    *Node
}
