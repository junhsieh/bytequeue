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
	//queue.IsDebug = true
	queue.debugInitByteArr()

	//var index int
	var err error

	//str := "AAA"
	str := "PZrbdBGRzbiBlWKaSuqqgjBYrq"

	for i := 0; i < 1; i++ {
		if _, err = queue.Push([]byte(str)); err != nil {
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
	//queue.IsDebug = true
	queue.debugInitByteArr()

	var strSize int
	var str string
	spaceLeft := queue.capacity

	for i := 0; i < 100; i++ {
		strSize = queue.debugRandInt(0, queue.capacity-headerEntrySize+1)
		str = queue.debugRandStringBytes(strSize)

		if _, err := queue.Push([]byte(str)); err != nil {
			t.Errorf("queue.Push([]byte(%d %s)): %v", strSize, str, err)
		}

		spaceLeft = spaceLeft - headerEntrySize - strSize + queue.popBytes

		if queue.availableSpaceAfterTail() != queue.debugCountX() {
			t.Errorf("queue.debugCountX() %d: %v vs %v; head: %d; tail: %d; count: %d; strSize: %d", i, queue.availableSpaceAfterTail(), queue.debugCountX(), queue.head, queue.tail, queue.count, strSize)
		}

		if queue.availableSpaceAfterTail() != spaceLeft {
			t.Errorf("spaceLeft %d: %v vs %v; head: %d; tail: %d; count: %d; strSize: %d", i, queue.availableSpaceAfterTail(), queue.debugCountX(), queue.head, queue.tail, queue.count, strSize)
		}
	}
}

func BenchmarkPush(b *testing.B) {
	queueSize := 30
	queue := NewByteQueue(queueSize)

	var strSize int
	var str string

	for i := 0; i < b.N; i++ {
		strSize = queue.debugRandInt(0, queueSize-headerEntrySize+1)
		str = queue.debugRandStringBytes(strSize)

		queue.Push([]byte(str))
	}
}
