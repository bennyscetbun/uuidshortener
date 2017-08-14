package uuidshortener

const (
	indexNumber byte = iota
	indexString
)

var indexStringString = string([]byte{indexString})

type sortedString []string

func (c sortedString) Len() int           { return len(c) }
func (c sortedString) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c sortedString) Less(i, j int) bool { return c[i] < c[j] }

// InsertAt is for optimization purpose, should be use with Index
func (s *sortedString) InsertAt(c string, index int) {
	lenS := len(*s)
	if cap(*s) == lenS {
		tmp := make(sortedString, lenS+1, lenS*2+1)
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

func (s *sortedString) Insert(c string) {
	lenS := len(*s)
	i, j := 0, lenS
	for i < j {
		h := i + (j-i)/2 // avoid overflow when computing h
		// i ≤ h < j
		if (*s)[h] > c {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// We are not optimistic here.
	if cap(*s) == lenS {
		tmp := make(sortedString, lenS+1, lenS*2+1)
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

//Index return the insertion index and if it was found or not
func (s *sortedString) Index(c string) (int, bool) {
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, len(*s)
	var cur string
	for i < j {
		h := i + (j-i)/2 // avoid overflow when computing h
		cur = (*s)[h]
		// i ≤ h < j
		if cur > c {
			i = h + 1 // preserves f(i-1) == false
		} else if cur == c {
			return h, true
		} else {
			j = h // preserves f(j) == true
		}
	}
	return i, false
}

type indexUUIDShortener struct {
	s sortedString
	m map[string]string
	i uint32
}

// NewMap returns a UUIDShortener that works with not constant uuid length and use maps
func NewIndexShortener(numberOfUUIDReserved int) UUIDShortener {
	return &indexUUIDShortener{
		s: make(sortedString, 0, numberOfUUIDReserved),
		m: make(map[string]string),
	}
}

func NewIndexExtender(numberOfUUIDReserved int) UUIDExtender {
	return &indexUUIDShortener{
		s: make(sortedString, 0, numberOfUUIDReserved),
		m: make(map[string]string),
	}
}

func parseIndex(b string) uint32 {
	ret := uint32(0)
	mul := uint(0)
	for i := 1; i < len(b); i++ {
		ret |= uint32(b[i]) << mul
		mul += 8
	}
	return ret
}

func formatIndex(v uint32) string {
	b := make([]byte, 1, 5)
	b[0] = indexNumber
	for v > 0 {
		b = append(b, byte(v))
		v = v >> 8
	}
	return string(b)
}

func (ius *indexUUIDShortener) Extend(shortenedUUID string, prefixToAdd string) string {
	if shortenedUUID[0] == indexNumber {
		return prefixToAdd + ius.s[parseIndex(shortenedUUID)]
	}
	shortenedUUID = shortenedUUID[1:]
	ius.s = append(ius.s, shortenedUUID)
	return prefixToAdd + shortenedUUID
}

func (ius *indexUUIDShortener) Short(uuid string, prefixLengthToRemove int) string {
	uuid = uuid[prefixLengthToRemove:]
	if index, found := ius.m[uuid]; found {
		return index
	}
	ius.m[uuid] = formatIndex(ius.i)
	ius.i++
	return indexStringString + uuid
}
