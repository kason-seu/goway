package worker

import (
	"goway/gograph/graph"
	"log"
)

// 点计算的引擎功能，这个代表一个worker上的所有点的执行引擎
type Engine struct {
	ID                 int
	vertexMap          map[int]graph.Vertex
	inactiveVertexChan chan graph.ActiveMessage
}

// create a new engine
func NewEngine(vertexMap map[int]graph.Vertex, ID int, inactiveVertexChan chan graph.ActiveMessage) *Engine {
	return &Engine{
		ID:                 ID,
		vertexMap:          vertexMap,
		inactiveVertexChan: inactiveVertexChan,
	}
}

func (e *Engine) GetVertices() map[int]graph.Vertex {
	return e.vertexMap
}

// SuperStep runs a superstep on each vertex of the engine.
func (e *Engine) SuperStep(workerMsgChan chan graph.VertexMessage, superStep int, done chan bool) {

	log.Printf("Worker %d: Starting superstep %d\n", e.ID, superStep)

	// 1. 遍历所有点，更新点的状态
	for _, vertex := range e.vertexMap {
		// 1.1 更新点的状态. 执行该点的超步计算, 做完工作后并返回该点是否还处于激活状态
		active := vertex.Update(superStep, workerMsgChan)

		if !active {
			// 1.2 如果该点不再激活，则将该点的状态发送到inactiveVertexChan
			e.inactiveVertexChan <- graph.ActiveMessage{
				VertexID:  vertex.GetID(),
				Active:    false,
				SuperStep: superStep,
			}
		}
	}
}

func (e *Engine) distributeMessage(msgs map[int][]graph.VertexMessage) {

	for vertexID, msg := range msgs {
		// 1.1 获取该点
		if vertex, ok := e.vertexMap[vertexID]; ok {

			// 1.2 将消息发送到该点
			vertex.ReceiveMessages(msg)
		} else {
			log.Printf("Worker %d: Vertex %d not found\n", e.ID, vertexID)
		}
	}
}
