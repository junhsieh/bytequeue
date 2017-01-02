package bytequeue

import (
	//"fmt"
	"errors"
)

// Size constants
const (
	KB = 1 << 10
	MB = 1 << 20
)

const (
	// Number of bytes used to keep the entry size information in the header.
	headerEntrySize = 4
)

// ByteQueue is a non-thread safe queue.
type ByteQueue struct {
	byteArr      []byte
	head         int
	tail         int
	count        int // number of entries
	capacity     int
	headerBuffer []byte
}

type queueError struct {
	message string
}

// NewByteQueue initializes new ByteQueue.
// capacity: in MB.
func NewByteQueue(capacityMB int) *ByteQueue {
	return &ByteQueue{
		byteArr:      make([]byte, capacityMB*MB),
		capacity:     capacityMB * MB,
		headerBuffer: make([]byte, headerEntrySize),
	}
}

// Push ...
func (bq *ByteQueue) Push(data []byte) (int, error) {
	dataLen := len(data)

	if (headerEntrySize + dataLen) > bq.availableSpaceAfterTail() {
		// pop some entries
	}

	index := bq.tail

	return index, nil
}

func (bq *ByteQueue) availableSpaceAfterTail() int {
	if bq.tail >= bq.head {
		return bq.capacity - (bq.tail - bq.head)
	}

	return bq.head - bq.tail
}
