package bytequeue

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strings"
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
	numOfEntries int // number of entries
	capacity     int
	headerBuffer []byte

	enableClearByte          bool // DEBUG: for testing purpose.
	enableByteArrDetail      bool // DEBUG: for testing purpose.
	enableNumOfPopBytesTrack bool // DEBUG: for testing purpose.
	numOfPopBytes            int  // DEBUG: for testing purpose.
	numOfAvailableBytes      int  // DEBUG: for testing purpose.
}

// NewByteQueue initializes new ByteQueue.
// capacity: in MB.
func NewByteQueue(capacityMB int) *ByteQueue {
	return &ByteQueue{
		byteArr:      make([]byte, capacityMB),
		capacity:     capacityMB,
		headerBuffer: make([]byte, headerEntrySize),

		numOfAvailableBytes: capacityMB, // DEBUG: for testing purpose.
	}
}

func (bq *ByteQueue) Pop() ([]byte, error) {
	if bq.numOfEntries == 0 {
		return nil, errors.New("Empty queue")
	}

	// get the header of the oldest entry.
	for i := 0; i < headerEntrySize; i++ {
		bq.headerBuffer[i] = bq.byteArr[bq.head]

		// DEBUG
		if bq.enableClearByte == true {
			bq.byteArr[bq.head] = 'X' // reset. Can be removed?
		}

		bq.head++

		// DEBUG
		if bq.enableNumOfPopBytesTrack == true {
			bq.numOfPopBytes++       // DEBUG: for testing purpose.
			bq.numOfAvailableBytes++ // DEBUG: for testing purpose.
		}

		if bq.head == bq.capacity {
			bq.head = 0
		}
	}

	// reset data bytes of the oldest entry and move head to the next entry.
	dataLen := int(binary.BigEndian.Uint32(bq.headerBuffer))
	data := make([]byte, headerEntrySize+dataLen)

	for i := 0; i < dataLen; i++ {
		data[i] = bq.byteArr[bq.head]

		// DEBUG
		if bq.enableClearByte == true {
			bq.byteArr[bq.head] = 'X' // reset. Can be removed?
		}

		bq.head++

		// DEBUG
		if bq.enableNumOfPopBytesTrack == true {
			bq.numOfPopBytes++       // DEBUG: for testing purpose
			bq.numOfAvailableBytes++ // DEBUG: for testing purpose
		}

		if bq.head == bq.capacity {
			bq.head = 0
		}
	}

	//
	bq.numOfEntries--

	return data, nil
}

// Push ...
// return the index of the pushed data
func (bq *ByteQueue) Push(data []byte) (int, error) {
	// DEBUG
	if bq.enableByteArrDetail == true {
		fmt.Printf(strings.Repeat("=", 130) + "\n")
	}

	// DEBUG: for testing purpose.
	if bq.enableNumOfPopBytesTrack == true {
		bq.numOfPopBytes = 0
	}

	dataLen := len(data)
	entryLen := headerEntrySize + dataLen

	if entryLen > bq.capacity {
		return -1, errors.New("Entry size is bigger than capacity.")
	}

	// save index for later use before pushing
	index := bq.tail

	for {
		if entryLen > bq.availableSpaceAfterTail() {
			if _, err := bq.Pop(); err != nil {
				return -1, err
			}

			// DEBUG
			if bq.enableByteArrDetail == true {
				fmt.Printf("info    (after pop):\t\tentryLen: %d\t\thead: %d\t\ttail: %d\t\tnumOfEntries: %d\t\tavailable: %d\n",
					entryLen, bq.head, bq.tail, bq.numOfEntries, bq.availableSpaceAfterTail())
				fmt.Printf("                   : %s\n", bq.debugGenByte())
				fmt.Printf("byteArr (after pop): %02v\n", bq.debugHighlightByteArr(bq.byteArr))
				fmt.Printf("\n")
			}
		} else {
			break
		}
	}

	// copy header
	binary.BigEndian.PutUint32(bq.headerBuffer, uint32(dataLen))

	bq.setByteArr(bq.headerBuffer)

	// copy data
	bq.setByteArr(data)

	//
	bq.numOfEntries++

	if bq.enableByteArrDetail == true {
		fmt.Printf("byteArr (afte push): %02v\n", bq.debugHighlightByteArr(bq.byteArr))
	}

	return index, nil
}

func (bq *ByteQueue) setByteArr(data []byte) {
	for _, v := range data {
		bq.byteArr[bq.tail] = v
		bq.tail++

		if bq.tail == bq.capacity {
			bq.tail = 0
		}

		// DEBUG
		if bq.enableNumOfPopBytesTrack == true {
			bq.numOfAvailableBytes-- // DEBUG: for testing purpose
		}
	}
}

func (bq *ByteQueue) availableSpaceAfterTail() int {
	if bq.tail > bq.head {
		//return bq.capacity - (bq.tail - bq.head)
		return bq.capacity - bq.tail + bq.head
	} else if bq.tail < bq.head {
		return bq.head - bq.tail
	}

	if bq.numOfEntries > 0 {
		return 0
	} else {
		return bq.capacity
	}
}
