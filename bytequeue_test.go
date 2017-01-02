package bytequeue

import (
	"fmt"
	"testing"
)

func TestDebug(t *testing.T) {
	queue := NewByteQueue(1)

	if index, err := queue.Push([]byte("How")); err != nil {
	} else {
		fmt.Printf("index: %d\n", index)
	}
	//t.Errorf("util.JSONDeepEqual(%s, %s) = %v", o.EncodeString(), s1, ok)
}
