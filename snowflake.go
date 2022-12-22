package main

import (
	"encoding/binary"
	"strconv"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

var (
	Epoch int64 = 1670628825675

	NodeBits uint8 = 3

	StepBits uint8 = 8
)

// A Node struct holds the basic information needed for a snowflake generator
// node
type Node struct {
	mu    sync.Mutex
	epoch time.Time
	time  int64
	node  int64
	step  int64

	stepMask  int64
	timeShift uint8
	stepShift uint8
}

// An ID is a custom type used for a snowflake ID.  This is used so we can
// attach methods onto the ID.
type ID uint64

// NewNode returns a new snowflake node that can be used to generate snowflake
// IDs
func NewNode(node int64) *Node {

	nodeMax := int64(-1 ^ (-1 << NodeBits))
	if node < 0 || node > nodeMax {
		panic("Node number must be between 0 and " + strconv.FormatInt(nodeMax, 10))
	}

	n := Node{
		node:      node,
		stepMask:  -1 ^ (-1 << StepBits),
		stepShift: NodeBits,
		timeShift: NodeBits + StepBits,
	}

	var curTime = time.Now()
	// add time.Duration to curTime to make sure we use the monotonic clock if available
	n.epoch = curTime.Add(time.Unix(Epoch/1000, (Epoch%1000)*1000000).Sub(curTime))

	return &n
}

func (n *Node) milliseconds() int64 { return time.Since(n.epoch).Milliseconds() }

// Generate creates and returns a unique snowflake ID
func (n *Node) Generate() ID {

	n.mu.Lock()
	defer n.mu.Unlock()

	now := n.milliseconds()

	if now == n.time {
		n.step = (n.step + 1) & n.stepMask

		if n.step == 0 {
			for now <= n.time {
				now = n.milliseconds()
			}

			log.Warn().Msg("mitigation triggered")
		}
	} else {
		n.step = 0
	}

	n.time = now

	return ID((now << n.timeShift) | (n.step << n.stepShift) | (n.node))
}

func (i ID) Uint64() uint64 { return uint64(i) }

func (i ID) UintBytes() []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}
