package validatorqueue

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

type Queue struct {
	mu   sync.Mutex
	data []common.Hash
}

func (q *Queue) Add(keys []common.Hash) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.data = append(q.data, keys...)
}

func (q *Queue) PopAll() []common.Hash {
	q.mu.Lock()
	defer q.mu.Unlock()

	res := q.data
	q.data = nil
	return res
}
