package uuidshortener

type internalSlice struct {
	uuid string
	key  string
}

type sortedInternalSlice []*internalSlice

func (c sortedInternalSlice) Len() int           { return len(c) }
func (c sortedInternalSlice) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c sortedInternalSlice) Less(i, j int) bool { return c[i].key < c[j].key }

// InsertAt is for optimization purpose, should be use with Index
func (s *sortedInternalSlice) InsertAt(c *internalSlice, index int) {
	lenS := len(*s)
	if cap(*s) == lenS {
		tmp := make(sortedInternalSlice, lenS+1, lenS*2+1)
		copy(tmp, (*s)[:index])
		tmp[index] = c
		copy(tmp[index+1:], (*s)[index:])
		*s = tmp
	} else {
		*s = append(*s, c)
		copy((*s)[index+1:], (*s)[index:])
		(*s)[index] = c
	}
}

func (s *sortedInternalSlice) Insert(c *internalSlice) {
	lenS := len(*s)
	i, j := 0, lenS
	for i < j {
		h := i + (j-i)/2 // avoid overflow when computing h
		// i ≤ h < j
		if (*s)[h].key > c.key {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// We are not optimistic here.
	if cap(*s) == lenS {
		tmp := make(sortedInternalSlice, lenS+1, lenS*2+1)
		copy(tmp, (*s)[:i])
		tmp[i] = c
		copy(tmp[i+1:], (*s)[i:])
		*s = tmp
	} else {
		*s = append(*s, c)
		copy((*s)[i+1:], (*s)[i:])
		(*s)[i] = c
	}

}

func (s *sortedInternalSlice) Get(c string) *internalSlice {
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, len(*s)
	var cur *internalSlice
	for i < j {
		h := i + (j-i)/2 // avoid overflow when computing h
		cur = (*s)[h]
		// i ≤ h < j
		if cur.key > c {
			i = h + 1 // preserves f(i-1) == false
		} else if cur.key == c {
			return cur
		} else {
			j = h // preserves f(j) == true
		}
	}
	return nil
}

//Index return the insertion index and if it was found or not
func (s *sortedInternalSlice) Index(c string) (int, bool) {
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, len(*s)
	var cur *internalSlice
	for i < j {
		h := i + (j-i)/2 // avoid overflow when computing h
		cur = (*s)[h]
		// i ≤ h < j
		if cur.key > c {
			i = h + 1 // preserves f(i-1) == false
		} else if cur.key == c {
			return h, true
		} else {
			j = h // preserves f(j) == true
		}
	}
	return i, false
}

type sliceUUIDShortener struct {
	s   sortedInternalSlice
	cur []byte
}

// NewSlice returns a UUIDShortener that works with not constant uuid length and use a slice
// slice size will be numberOfUUIDReserved * 2
func NewSliceShortener(numberOfUUIDReserved int) UUIDShortener {
	return &sliceUUIDShortener{
		s:   make([]*internalSlice, 0, numberOfUUIDReserved*2),
		cur: []byte{0},
	}
}

func NewSliceExtender(numberOfUUIDReserved int) UUIDExtender {
	return &sliceUUIDShortener{
		s:   make([]*internalSlice, 0, numberOfUUIDReserved*2),
		cur: []byte{0},
	}
}

func (sus *sliceUUIDShortener) getNextCur() string {
	ret := string(sus.cur)
	for i := 0; i < len(sus.cur); i++ {
		sus.cur[i]++
		if sus.cur[i] != 0 {
			break
		}
	}
	if sus.cur[len(sus.cur)-1] == 0 {
		sus.cur = append(sus.cur, byte(1))
	}
	return ret
}

func (sus *sliceUUIDShortener) Extend(shortenedUUID string, prefixToAdd string) string {
	index, found := sus.s.Index(shortenedUUID)
	if found {
		return prefixToAdd + sus.s[index].uuid
	}
	cur := sus.getNextCur()
	sus.s.Insert(&internalSlice{
		uuid: shortenedUUID,
		key:  cur,
	})
	return prefixToAdd + shortenedUUID
}

func (sus *sliceUUIDShortener) Short(uuid string, prefixLengthToRemove int) string {
	uuid = uuid[prefixLengthToRemove:]

	index, found := sus.s.Index(uuid)
	if found {
		return sus.s[index].uuid
	}
	cur := sus.getNextCur()
	sus.s.InsertAt(&internalSlice{
		uuid: cur,
		key:  uuid,
	}, index)
	return uuid
}
