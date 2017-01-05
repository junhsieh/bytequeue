package bytequeue

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

func (bq *ByteQueue) debugInitByteArr() {
	for k, _ := range bq.byteArr {
		bq.byteArr[k] = 'X'
	}
}

func (bq *ByteQueue) debugCountX() int {
	count := 0

	for _, v := range bq.byteArr {
		if v == 'X' {
			count++
		}
	}

	return count
}

func (bq *ByteQueue) debugHighlightByteArr(data []byte) string {
	str := "["

	for k, v := range data {
		if k == bq.head {
			str += ColorBegin + "31m" + fmt.Sprintf("%02v", v) + ColorEnd + " "
		} else if k == bq.tail {
			str += ColorBegin + "35m" + fmt.Sprintf("%02v", v) + ColorEnd + " "
		} else if v == 'X' {
			str += ColorBegin + "32m" + fmt.Sprintf("%02v", v) + ColorEnd + " "
		} else {
			str += fmt.Sprintf("%02v", v) + " "
		}
	}

	return str + "]"
}

func (bq *ByteQueue) debugGenByte() string {
	str := "("

	for i := 0; i < bq.capacity; i++ {
		str += fmt.Sprintf("%02d ", i)
	}
	return str + ")"
}

func (bq *ByteQueue) debugRandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (bq *ByteQueue) debugRandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
