package circularqueue

import (
	"fmt"
	"math/rand"
)

const (
	ColorBegin = "\033["
	ColorEnd   = "\033[0m"
)

// without X
//const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWYZ"
const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWYZ"

func (cq *CircularQueue) debugInitByteArr() {
	for k, _ := range cq.byteArr {
		cq.byteArr[k] = 'X'
	}
}

func (cq *CircularQueue) debugCountX() int {
	count := 0

	for _, v := range cq.byteArr {
		if v == 'X' {
			count++
		}
	}

	return count
}

func (cq *CircularQueue) debugHighlightByteArr(data []byte) string {
	str := "["

	for k, v := range data {
		if k == cq.head {
			str += ColorBegin + "31m" + fmt.Sprintf("%02v", v) + ColorEnd + " "
		} else if k == cq.tail {
			str += ColorBegin + "35m" + fmt.Sprintf("%02v", v) + ColorEnd + " "
		} else if v == 'X' {
			str += ColorBegin + "32m" + fmt.Sprintf("%02v", v) + ColorEnd + " "
		} else {
			str += fmt.Sprintf("%02v", v) + " "
		}
	}

	return str + "]"
}

func (cq *CircularQueue) debugGenByte() string {
	str := "("

	for i := 0; i < cq.capacity; i++ {
		str += fmt.Sprintf("%02d ", i)
	}
	return str + ")"
}

func (cq *CircularQueue) debugRandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (cq *CircularQueue) debugRandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
