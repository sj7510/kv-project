package index

import (
	"bytes"
	"github.com/google/btree"
	"go-kv/data"
)

// Indexer 抽象索引接口，接入其它数据结构，实现这个接口
type Indexer interface {
	// Put 向索引中存储 key 对应的数据位置信息
	Put(keu []byte, pos *data.LogRecordPos) bool

	// Get 根据 key 取出对应的索引位置信息
	Get(key []byte) *data.LogRecordPos

	// Delete 根据 key 删除对应的索引位置信息
	Delete(key []byte) bool
}

type Item struct {
	key []byte
	pos *data.LogRecordPos
}

func (ai *Item) Less(bi btree.Item) bool {
	return bytes.Compare(ai.key, bi.(*Item).key) == -1
}
