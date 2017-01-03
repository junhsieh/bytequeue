package bytequeue

import (
	//"fmt"
	"encoding/binary"
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
		byteArr:      make([]byte, capacityMB),
		capacity:     capacityMB,
		headerBuffer: make([]byte, headerEntrySize),
	}
}

func (bq *ByteQueue) GetByteArr() []byte {
	return bq.byteArr
}

// Push ...
// return the number of bytes copied
func (bq *ByteQueue) Push(data []byte) (int, error) {
	dataLen := len(data)
	entryLen := headerEntrySize + dataLen

	if entryLen > bq.capacity {
		return 0, errors.New("Entry size is bigger than capacity.")
	}

	if (headerEntrySize + dataLen) > bq.availableSpaceAfterTail() {
		// pop some entries until the space is enough
		// also check do not exceed the size.
	}

	// copy header
	binary.BigEndian.PutUint32(bq.headerBuffer, uint32(dataLen))

	bq.copyByte(bq.headerBuffer)

	// copy data
	bq.copyByte(data)

	//
	bq.count++

	return bq.tail, nil
}

func (bq *ByteQueue) copyByte(data []byte) {
	for _, v := range data {
		bq.byteArr[bq.tail] = v
		bq.tail++
		if bq.tail == bq.capacity {
			bq.tail = 0
		}
	}
}

func (bq *ByteQueue) availableSpaceAfterTail() int {
	if bq.tail >= bq.head {
		//return bq.capacity - (bq.tail - bq.head)
		return bq.capacity - bq.tail + bq.head
	}

	return bq.head - bq.tail
}
