package bytequeue

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestDebug(t *testing.T) {
	queue := NewByteQueue(30)
	queue.IsDebug = true
	queue.debugInitByteArr()

	var index int
	var err error

	//str := "AAA"
	str := "PZrbdBGRzbiBlWKaSuqqgjBYrqPc"

	for i := 0; i < 1; i++ {
		fmt.Printf("========================================")
		fmt.Printf("========================================")
		fmt.Printf("========================================\n")

		if index, err = queue.Push([]byte(str)); err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}

	fmt.Printf("\n")
	fmt.Printf("head: %v\n", queue.head)
	fmt.Printf("tail: %v\n", queue.tail)
	fmt.Printf("index: %v\n", index)
	//fmt.Printf("byteArr (afte push): %v\n", queue.byteArr)
	//t.Errorf("util.JSONDeepEqual(%s, %s) = %v", o.EncodeString(), s1, ok)
}

func TestAvailableSpace(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	queueSize := 30
	queue := NewByteQueue(queueSize)
	queue.debugInitByteArr()

	var strSize int
	var str string

	for i := 0; i < 70; i++ {
		strSize = queue.debugRandInt(0, queueSize-headerEntrySize)
		str = queue.debugRandStringBytes(strSize)

		if _, err := queue.Push([]byte(str)); err != nil {
			t.Errorf("queue.Push([]byte(%d %s)): %v", strSize, str, err)
		}

		if queue.debugCountX() != queue.availableSpaceAfterTail() {
			t.Errorf("availableSpaceAfterTail: %v vs %v; head: %d; tail: %d", queue.debugCountX(), queue.availableSpaceAfterTail(), queue.head, queue.tail)
		}
	}

}

func BenchmarkPush(b *testing.B) {
	queue := NewByteQueue(30)

	for i := 0; i < b.N; i++ {
		queue.Push([]byte("AAA"))
	}
}
