package management

import (
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"hash/fnv"
	"sync"
)

type Node struct{
	id uint64
	host string
	port int

	isAlive bool

	closeFunc func()error
	conn *grpc.ClientConn
}

func (node *Node) Connect() error{
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(node.host +":"+fmt.Sprint(node.port), grpc.WithInsecure())
	if err != nil {
		return errors.New(fmt.Sprintf("did not connect: %s", err))
	}

	node.closeFunc =conn.Close
	node.conn = conn

	return nil
}

func (node *Node) GetConnection()  *grpc.ClientConn{
	return node.conn
}

func (node *Node) Close() {
	node.Close()
}

type  Nodes struct{
	Nodes []*Node
	mux sync.Mutex
}

func NewNodes() *Nodes{
	nodes := &Nodes {}
	nodes.Nodes = []*Node{}

	go nodes.HealthCycle()

	return nodes
}

func (nodes *Nodes) RegisterNode(host string, port int) error{
	h := fnv.New64()
	h.Write([]byte(host+fmt.Sprint(port)))
	nodeID := h.Sum64()

	for _, node := range nodes.Nodes{
		if node.id == nodeID{
			return errors.New(fmt.Sprintf("There's already a node registered at host:%s - port:%v",host,port))
		}
	}

	newNode := &Node{id:nodeID,host: host,port:port}
	err := newNode.Connect()
	if err != nil {
		return errors.New(fmt.Sprintf("There's a problem connecting to node at host:%s - port:%v with error:%s",host,port, err.Error()))
	}

	nodes.mux.Lock()
	nodes.Nodes = append(nodes.Nodes, newNode)
	nodes.mux.Unlock()

	return nil
}

func (nodes *Nodes) GetNode(key string) (*Node,error){
	if len(nodes.Nodes) == 0 {
		return nil, errors.New("There are 0 nodes registered.")
	}
	selectedNode := nodes.Nodes[0]
	return selectedNode,nil
}

func (nodes *Nodes) HealthCycle(){

}