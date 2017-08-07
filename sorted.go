package uuidshortener

type sortedPrefixDictionnary []*prefixDictionnary

func (c sortedPrefixDictionnary) Len() int           { return len(c) }
func (c sortedPrefixDictionnary) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c sortedPrefixDictionnary) Less(i, j int) bool { return c[i].curChar < c[j].curChar }

// InsertAt is for optimization purpose, should be use with Index
func (s *sortedPrefixDictionnary) InsertAt(c *prefixDictionnary, index int) {
	lenS := len(*s)
	if cap(*s) == lenS {
		tmp := make(sortedPrefixDictionnary, lenS+1, lenS*2+1)
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

func (s *sortedPrefixDictionnary) Insert(c *prefixDictionnary) {
	lenS := len(*s)
	i, j := 0, lenS
	for i < j {
		h := i + (j-i)/2 // avoid overflow when computing h
		// i ≤ h < j
		if (*s)[h].curChar > c.curChar {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// We are not optimistic here.
	if cap(*s) == lenS {
		tmp := make(sortedPrefixDictionnary, lenS+1, lenS*2+1)
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

func (s *sortedPrefixDictionnary) Get(c byte) *prefixDictionnary {
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, len(*s)
	var cur *prefixDictionnary
	for i < j {
		h := i + (j-i)/2 // avoid overflow when computing h
		cur = (*s)[h]
		// i ≤ h < j
		if cur.curChar > c {
			i = h + 1 // preserves f(i-1) == false
		} else if cur.curChar == c {
			return cur
		} else {
			j = h // preserves f(j) == true
		}
	}
	return nil
}

//Index return the insertion index and if it was found or not
func (s *sortedPrefixDictionnary) Index(c byte) (int, bool) {
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, len(*s)
	var cur *prefixDictionnary
	for i < j {
		h := i + (j-i)/2 // avoid overflow when computing h
		cur = (*s)[h]
		// i ≤ h < j
		if cur.curChar > c {
			i = h + 1 // preserves f(i-1) == false
		} else if cur.curChar == c {
			return h, true
		} else {
			j = h // preserves f(j) == true
		}
	}
	return i, false
}
