package circularqueue

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

// CircularQueue is a non-thread safe queue.
type CircularQueue struct {
	byteArr      []byte
	head         int
	tail         int
	numOfEntries int // number of entries
	capacity     int
	headerBuffer []byte

	enableClearByte          bool // DEBUG: for testing purpose.
	enableByteArrDetail      bool // DEBUG: for testing purpose.
	enableNumOfPopBytesTrack bool // DEBUG: for testing purpose.
	enablePopWithoutData     bool // DEBUG: for testing purpose.
	numOfPopBytes            int  // DEBUG: for testing purpose.
	numOfAvailableBytes      int  // DEBUG: for testing purpose.
}

// NewCircularQueue initializes new CircularQueue.
// capacity: in MB.
func NewCircularQueue(capacityMB int) *CircularQueue {
	return &CircularQueue{
		byteArr:      make([]byte, capacityMB),
		capacity:     capacityMB,
		headerBuffer: make([]byte, headerEntrySize),

		numOfAvailableBytes: capacityMB, // DEBUG: for testing purpose.
	}
}

func (cq *CircularQueue) Pop() ([]byte, error) {
	if cq.numOfEntries == 0 {
		return nil, errors.New("Empty queue")
	}

	// get the header of the oldest entry.
	for i := 0; i < headerEntrySize; i++ {
		cq.headerBuffer[i] = cq.byteArr[cq.head]

		// DEBUG
		if cq.enableClearByte == true {
			cq.byteArr[cq.head] = 'X' // reset. Can be removed?
		}

		cq.head++

		if cq.head == cq.capacity {
			cq.head = 0
		}

		// DEBUG
		if cq.enableNumOfPopBytesTrack == true {
			cq.numOfPopBytes++       // DEBUG: for testing purpose.
			cq.numOfAvailableBytes++ // DEBUG: for testing purpose.
		}
	}

	// reset data bytes of the oldest entry and move head to the next entry.
	dataLen := int(binary.BigEndian.Uint32(cq.headerBuffer))
	data := make([]byte, headerEntrySize+dataLen)

	for i := 0; i < dataLen; i++ {
		data[i] = cq.byteArr[cq.head]

		// DEBUG
		if cq.enableClearByte == true {
			cq.byteArr[cq.head] = 'X' // reset. Can be removed?
		}

		cq.head++

		if cq.head == cq.capacity {
			cq.head = 0
		}

		// DEBUG
		if cq.enableNumOfPopBytesTrack == true {
			cq.numOfPopBytes++       // DEBUG: for testing purpose
			cq.numOfAvailableBytes++ // DEBUG: for testing purpose
		}
	}

	//
	cq.numOfEntries--

	return data, nil
}

func (cq *CircularQueue) PopWithoutData() {
	if cq.numOfEntries == 0 {
		return
	}

	// get the header of the oldest entry.
	for i := 0; i < headerEntrySize; i++ {
		cq.headerBuffer[i] = cq.byteArr[cq.head]
		cq.head++

		if cq.head == cq.capacity {
			cq.head = 0
		}

		// DEBUG
		if cq.enableNumOfPopBytesTrack == true {
			cq.numOfPopBytes++       // DEBUG: for testing purpose
			cq.numOfAvailableBytes++ // DEBUG: for testing purpose
		}
	}

	// reset data bytes of the oldest entry and move head to the next entry.
	dataLen := int(binary.BigEndian.Uint32(cq.headerBuffer))
	cq.head = cq.head + dataLen

	if cq.head >= cq.capacity {
		cq.head = cq.head - cq.capacity
	}

	// DEBUG
	if cq.enableNumOfPopBytesTrack == true {
		cq.numOfPopBytes += dataLen       // DEBUG: for testing purpose
		cq.numOfAvailableBytes += dataLen // DEBUG: for testing purpose
	}

	//
	cq.numOfEntries--
}

// Push ...
// return the index of the pushed data
func (cq *CircularQueue) Push(data []byte) (int, error) {
	// DEBUG
	if cq.enableByteArrDetail == true {
		fmt.Printf(strings.Repeat("=", 130) + "\n")
	}

	// DEBUG: for testing purpose.
	if cq.enableNumOfPopBytesTrack == true {
		cq.numOfPopBytes = 0
	}

	dataLen := len(data)
	entryLen := headerEntrySize + dataLen

	if entryLen > cq.capacity {
		return -1, errors.New("Entry size is bigger than capacity.")
	}

	// Save index for later use before pushing
	index := cq.tail

	for {
		if entryLen > cq.availableSpaceAfterTail() {
			if cq.enablePopWithoutData == false {
				if _, err := cq.Pop(); err != nil {
					return -1, err
				}
			} else {
				cq.PopWithoutData()
			}

			// DEBUG
			if cq.enableByteArrDetail == true {
				fmt.Printf("info    (after pop):\t\tentryLen: %d\t\thead: %d\t\ttail: %d\t\tnumOfEntries: %d\t\tavailable: %d\n",
					entryLen, cq.head, cq.tail, cq.numOfEntries, cq.availableSpaceAfterTail())
				fmt.Printf("                   : %s\n", cq.debugGenByte())
				fmt.Printf("byteArr (after pop): %02v\n", cq.debugHighlightByteArr(cq.byteArr))
				fmt.Printf("\n")
			}
		} else {
			break
		}
	}

	// copy header
	binary.BigEndian.PutUint32(cq.headerBuffer, uint32(dataLen))

	cq.setByteArr(cq.headerBuffer)

	// copy data
	cq.setByteArr(data)

	//
	cq.numOfEntries++

	// DEBUG
	if cq.enableByteArrDetail == true {
		fmt.Printf("byteArr (afte push): %02v\n", cq.debugHighlightByteArr(cq.byteArr))
	}

	return index, nil
}

func (cq *CircularQueue) setByteArr(data []byte) {
	for _, v := range data {
		cq.byteArr[cq.tail] = v
		cq.tail++

		if cq.tail == cq.capacity {
			cq.tail = 0
		}

		// DEBUG
		if cq.enableNumOfPopBytesTrack == true {
			cq.numOfAvailableBytes-- // DEBUG: for testing purpose
		}
	}
}

func (cq *CircularQueue) availableSpaceAfterTail() int {
	if cq.tail > cq.head {
		//return cq.capacity - (cq.tail - cq.head)
		return cq.capacity - cq.tail + cq.head
	} else if cq.tail < cq.head {
		return cq.head - cq.tail
	}

	if cq.numOfEntries > 0 {
		return 0
	} else {
		return cq.capacity
	}
}
