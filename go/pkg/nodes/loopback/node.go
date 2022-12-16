/** Simple loopback node-type written in Go code.
 *
 * @author Steffen Vogel <post@steffenvogel.de>
 * @copyright 2014-2022, Institute for Automation of Complex Power Systems, EONERC
 * @license Apache 2.0
 *********************************************************************************/

package nodes

import (
	"encoding/json"
	"fmt"

	"git.rwth-aachen.de/acs/public/villas/node/go/pkg/errors"
	"git.rwth-aachen.de/acs/public/villas/node/go/pkg/nodes"
)

type Node struct {
	nodes.BaseNode

	channel chan []byte

	Config LoopbackConfig
}

type LoopbackConfig struct {
	nodes.NodeConfig

	Value int `json:"value"`
}

func NewNode() nodes.Node {
	return &Node{
		BaseNode: nodes.NewBaseNode(),
		channel:  make(chan []byte, 1024),
	}
}

func (n *Node) Parse(c []byte) error {
	return json.Unmarshal(c, &n.Config)
}

func (n *Node) Check() error {
	return nil
}

func (n *Node) Prepare() error {
	return nil
}

func (n *Node) Start() error {
	return n.BaseNode.Start()
}

func (n *Node) Read() ([]byte, error) {
	select {
	case <-n.Stopped:
		return nil, errors.ErrEndOfFile

	case buf := <-n.channel:
		return buf, nil
	}
}

func (n *Node) Write(data []byte) error {
	n.channel <- data

	return nil
}

func (n *Node) PollFDs() ([]int, error) {
	return []int{}, nil
}

func (n *Node) NetemFDs() ([]int, error) {
	return []int{}, nil
}

func (n *Node) Details() string {
	return fmt.Sprintf("value=%d", n.Config.Value)
}

func init() {
	nodes.RegisterNodeType("go.loopback", "A loopback node implmented in Go", NewNode, nodes.NodeSupportsRead|nodes.NodeSupportsWrite|nodes.NodeHidden)
}
