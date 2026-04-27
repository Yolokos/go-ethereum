package validatorqueue

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

type Queue struct {
	mu   sync.Mutex
	data map[common.Hash]map[uint64][]common.Hash
}

func New() *Queue {
	return &Queue{
		data: make(map[common.Hash]map[uint64][]common.Hash),
	}
}

func (q *Queue) Add(parent common.Hash, timestamp uint64, keys []common.Hash) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.data[parent] == nil {
		q.data[parent] = make(map[uint64][]common.Hash)
	}
	q.data[parent][timestamp] = keys
}

func (q *Queue) Pop(parent common.Hash, timestamp uint64) []common.Hash {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.data[parent] == nil {
		return nil
	}
	res := q.data[parent][timestamp]
	delete(q.data[parent], timestamp)
	if len(q.data[parent]) == 0 {
		delete(q.data, parent)
	}
	return res
}
