package bytequeue

import (
	"fmt"
	"testing"
)

func TestDebug(t *testing.T) {
	queue := NewByteQueue(30)

	var index int
	var err error

	for i := 0; i < 3; i++ {
		if index, err = queue.Push([]byte("AAA")); err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}

	fmt.Printf("index: %v\n", index)
	fmt.Printf("byteArr: %v\n", queue.GetByteArr())
	//t.Errorf("util.JSONDeepEqual(%s, %s) = %v", o.EncodeString(), s1, ok)
}
