package graph

/*
*
图里面的点传送到其他点的消息

通过network将msg在worker之间进行传递

*
*/
type VertexMessage struct {
	FromID    int
	Value     interface{}
	ToID      int
	SuperStep int
}

// 一个超步完成之后，表达该节点是否被激活
type ActiveMessage struct {
	VertexID  int
	Active    bool
	SuperStep int
}
