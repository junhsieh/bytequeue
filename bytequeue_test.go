package bytequeue

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func TestDebug(t *testing.T) {
	queueSize := 30
	queue := NewByteQueue(queueSize)
	//queue.enableByteArrDetail = true
	queue.debugInitByteArr()

	//var index int
	var err error

	//data := "AAA"
	data := "PZrbdBGRzbiBlWKaSuqqgjBYrq"

	for i := 0; i < 1; i++ {
		if _, err = queue.Push([]byte(data)); err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}

	//fmt.Printf("\n")
	//fmt.Printf("head: %v\n", queue.head)
	//fmt.Printf("tail: %v\n", queue.tail)
	//fmt.Printf("index: %v\n", index)
	//fmt.Printf("byteArr (afte push): %v\n", queue.byteArr)
	//t.Errorf("util.JSONDeepEqual(%s, %s) = %v", o.EncodeString(), s1, ok)
}

func TestAvailableSpace(t *testing.T) {
	//rand.Seed(time.Now().UTC().UnixNano())

	queueSize := 30
	queue := NewByteQueue(queueSize)
	queue.enableClearByte = true
	//queue.enableByteArrDetail = true
	queue.enableNumOfPopBytesTrack = true
	queue.debugInitByteArr()

	var dataLen int
	var data string
	//var index int

	//checkHead := 0
	checkTail := 0
	checkSpaceLeft := queue.capacity

	for i := 0; i < 100000; i++ {
		dataLen = queue.debugRandInt(0, queue.capacity-headerEntrySize+1)
		data = queue.debugRandStringBytes(dataLen)

		if _, err := queue.Push([]byte(data)); err != nil {
			t.Errorf("queue.Push([]byte(%d %s)): %v", dataLen, data, err)
		}

		// check head

		// check tail
		checkTail = checkTail + headerEntrySize + dataLen

		if checkTail >= queue.capacity {
			checkTail = checkTail - queue.capacity
		}

		if queue.tail != checkTail {
			t.Errorf("checkTail %d: %v vs %v; head: %d; tail: %d; count: %d; dataLen: %d", i, queue.availableSpaceAfterTail(), queue.debugCountX(), queue.head, queue.tail, queue.count, dataLen)
		}

		// check available space
		checkSpaceLeft = checkSpaceLeft - headerEntrySize - dataLen + queue.numOfPopBytes

		if queue.availableSpaceAfterTail() != queue.debugCountX() {
			t.Errorf("queue.debugCountX() %d: %v vs %v; head: %d; tail: %d; count: %d; dataLen: %d", i, queue.availableSpaceAfterTail(), queue.debugCountX(), queue.head, queue.tail, queue.count, dataLen)
		}

		if queue.availableSpaceAfterTail() != checkSpaceLeft {
			t.Errorf("checkSpaceLeft %d: %v vs %v; head: %d; tail: %d; count: %d; dataLen: %d", i, queue.availableSpaceAfterTail(), queue.debugCountX(), queue.head, queue.tail, queue.count, dataLen)
		}

		if queue.availableSpaceAfterTail() != queue.numOfAvailableBytes {
			t.Errorf("queue.numOfAvailableBytes %d: %v vs %v; head: %d; tail: %d; count: %d; dataLen: %d", i, queue.availableSpaceAfterTail(), queue.debugCountX(), queue.head, queue.tail, queue.count, dataLen)
		}
	}
}

func BenchmarkPush(b *testing.B) {
	queueSize := 30
	queue := NewByteQueue(queueSize)

	var dataLen int
	var data string

	for i := 0; i < b.N; i++ {
		dataLen = queue.debugRandInt(0, queueSize-headerEntrySize+1)
		data = queue.debugRandStringBytes(dataLen)

		queue.Push([]byte(data))
	}
}
