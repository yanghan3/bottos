package node

import (
	"container/list"
	"fmt"
	"sync"

	"github.com/bottos-project/bottos/core/db"
)

type Trie struct {
	mu       sync.Mutex
	root     Node
	roothash []byte
	cache    *db.Cache

	revisions *list.List
}

var indices = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "[17]"}

type Node interface {
	Value() Node
	Copy(*Trie) Node // All nodes, for now, return them self
	Dirty() bool
	fstring(string) string
	Hash() interface{}
	RlpData() interface{}
	setDirty(dirty bool)
}

// Value node
func (self *FullNode) String() string { return self.fstring("") }

//func (self *HashNode) fstring(ind string) string  { return fmt.Sprintf("< %x > ", self.key) }
func (self *HashNode) fstring(ind string) string {
	//lmq return fmt.Sprintf("%v", self.trie.trans(self))
	return ""
}

// Full node
func (self *FullNode) fstring(ind string) string {
	resp := fmt.Sprintf("[\n%s  ", ind)
	for i, node := range self.nodes {
		if node == nil {
			resp += fmt.Sprintf("%s: <nil> ", indices[i])
		} else {
			resp += fmt.Sprintf("%s: %v", indices[i], node.fstring(ind+"  "))
		}
	}

	return resp + fmt.Sprintf("\n%s] ", ind)
}
