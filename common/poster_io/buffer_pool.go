package poster_io

import (
	"sync"
)

type bufferPool struct {
	Buffer map[string]map[int]([]byte)
	// 长度
	Size  map[string](int)
	mutex sync.RWMutex
}

var buffer_pool_instance *bufferPool

func GetBufferPoolInstance() *bufferPool {
	once.Do(func() {
		buffer_pool_instance = &bufferPool{
			Buffer: make(map[string]map[int]([]byte)),
		}
	})
	return buffer_pool_instance
}

func (bp *bufferPool) SetSize(key string, size int) {
	bp.mutex.Lock()
	defer bp.mutex.Unlock()
	bp.Size[key] = size
}

func (bp *bufferPool) GetSize(key string, size int) int {
	bp.mutex.RLock()
	defer bp.mutex.Unlock()
	v, ok := bp.Size[key]
	if !ok {
		return 0
	} else {
		return v
	}
}

func (bp *bufferPool) GetChunk(key string, idx int) []byte {
	bp.mutex.RLock()
	defer bp.mutex.Unlock()
	mm, ok := bp.Buffer[key]
	if !ok {
		return nil
	} else {
		m, ok := mm[idx]
		if !ok {
			return nil
		}
		return m
	}
}
func (bp *bufferPool) AddChunk(key string, idx int, data []byte) {
	bp.mutex.Lock()
	defer bp.mutex.Unlock()
	mm, ok := bp.Buffer[key]
	if !ok {
		mm := make(map[int][]byte)
		bp.Buffer[key] = mm
	}
	mm[idx] = data
}

func RemoveChunkByKey(key string) {

}
