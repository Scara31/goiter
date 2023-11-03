package iter

import (
	"strconv"
	"testing"
)

func TestIterator(t *testing.T) {
	slice1 := []string{"a", "b", "c"}
	slice2 := []string{"d"}
	slice3 := []string{"e", "f"}

	totalLen := len(slice1) + len(slice2) + len(slice3)
	mergedSlice := make([]string, 0, totalLen)

	mergedSlice = append(mergedSlice, slice1...)
	mergedSlice = append(mergedSlice, slice2...)
	mergedSlice = append(mergedSlice, slice3...)

	iter := NewIterator(slice1, slice2, slice3)

	if iter.Len != totalLen {
		t.Errorf("iterator and slice length mismatch, %d != %d", iter.Len, totalLen)
	}

	i := 0
	for {
		got, ok := iter.Take()
		if !ok {
			break
		}

		want := mergedSlice[i]
		i++

		if got != want {
			t.Errorf("iterator and slice mismatch, %s != %s", got, want)
		}
	}
}

const sliceLen1 = 10 * 1000
const sliceLen2 = 30 * 1000
const sliceLen3 = 5 * 1000

var slice1 = genSlice(sliceLen1)
var slice2 = genSlice(sliceLen2)
var slice3 = genSlice(sliceLen3)

func BenchmarkStdMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newSlice := make([]string, 0, sliceLen1+sliceLen2+sliceLen3)
		newSlice = append(newSlice, slice1...)
		newSlice = append(newSlice, slice2...)
		newSlice = append(newSlice, slice3...)

		for _, item := range newSlice {
			doNothing(item)
		}
	}
}

func BenchmarkIter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iter := NewIterator(slice1, slice2, slice3)

		for {
			v, ok := iter.Take()
			if !ok {
				break
			}
			doNothing(v)
		}
	}
}

func genSlice(sliceLen int) []string {
	s := make([]string, sliceLen)
	for i := 0; i < sliceLen; i++ {
		s[i] = strconv.Itoa(i)
	}
	return s
}

func doNothing(v any) {}
