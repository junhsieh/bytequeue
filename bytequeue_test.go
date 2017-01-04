package bytequeue

import (
	"fmt"
	"testing"
)

func TestDebug(t *testing.T) {
	queue := NewByteQueue(30)
	queue.IsDebug = true
	queue.DebugInitByteArr()

	var index int
	var err error

	for i := 0; i < 7; i++ {
		fmt.Printf("========================================")
		fmt.Printf("========================================")
		fmt.Printf("========================================\n")

		if index, err = queue.Push([]byte("AAA")); err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}

	fmt.Printf("\n")
	fmt.Printf("head: %v\n", queue.GetHead())
	fmt.Printf("tail: %v\n", queue.GetTail())
	fmt.Printf("index: %v\n", index)
	//fmt.Printf("byteArr (afte push): %v\n", queue.GetByteArr())
	//t.Errorf("util.JSONDeepEqual(%s, %s) = %v", o.EncodeString(), s1, ok)
}

func TestAvailableSpace(t *testing.T) {
	queue := NewByteQueue(30)
	queue.DebugInitByteArr()

	for i := 0; i < 70000; i++ {
		if _, err := queue.Push([]byte("AAA")); err != nil {
			t.Errorf("ERR: queue.Push: %v", err)
		}

		if queue.DebugCountX() != queue.availableSpaceAfterTail() {
			t.Errorf("ERR: availableSpaceAfterTail: %v %v", queue.DebugCountX(), queue.availableSpaceAfterTail())
		}
	}
}

func BenchmarkPush(b *testing.B) {
	queue := NewByteQueue(30)

	for i := 0; i < b.N; i++ {
		queue.Push([]byte("AAA"))
	}
}
