package iter

type Iterator struct {
	Len int

	data [][]string
	s    int // slice index cursor
	i    int // slice element index cursor
}

func NewIterator(slices ...[]string) *Iterator {
	totalLen := 0
	filledSlices := make([][]string, 0, len(slices))

	for _, slice := range slices {
		l := len(slice)
		if l != 0 { // nil slice also has len = 0
			totalLen += l
			filledSlices = append(filledSlices, slice)
		}
	}

	return &Iterator{
		Len:  totalLen,
		data: filledSlices,
		s:    0,
		i:    0,
	}
}

func (iter *Iterator) Take() (v string, ok bool) {
	// element is out of current slice bound
	if iter.i > len(iter.data[iter.s])-1 {
		iter.s++
		iter.i = 0

		// element is out of global bound
		if iter.s > len(iter.data)-1 {
			return "", false
		}
	}

	v, ok = iter.data[iter.s][iter.i], true
	iter.i++
	return v, ok
}
