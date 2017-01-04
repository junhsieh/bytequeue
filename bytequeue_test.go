package bytequeue

import (
	"fmt"
	"testing"
)

func TestDebug(t *testing.T) {
	queue := NewByteQueue(30)

	var index int
	var err error

	for i := 0; i < 6; i++ {
		fmt.Printf("==================================\n")

		if index, err = queue.Push([]byte("AAA")); err != nil {
			fmt.Printf("err: %v\n", err)
		} else {
			fmt.Printf("byteArr (afte push): %02v\n", queue.GetByteArr())
		}
	}

	fmt.Printf("head: %v\n", queue.GetHead())
	fmt.Printf("tail: %v\n", queue.GetTail())
	fmt.Printf("index: %v\n", index)
	//fmt.Printf("byteArr (afte push): %v\n", queue.GetByteArr())
	//t.Errorf("util.JSONDeepEqual(%s, %s) = %v", o.EncodeString(), s1, ok)
}
