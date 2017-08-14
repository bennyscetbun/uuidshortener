package uuidshortener

type mapsUUIDShortener struct {
	m   map[string]string
	cur []byte
}

// NewMap returns a UUIDShortener that works with not constant uuid length and use maps
func NewMapShortener() UUIDShortener {
	return &mapsUUIDShortener{
		m:   make(map[string]string),
		cur: []byte{0},
	}
}

func NewMapExtender() UUIDExtender {
	return &mapsUUIDShortener{
		m:   make(map[string]string),
		cur: []byte{0},
	}
}

func (mus *mapsUUIDShortener) getNextCur() string {
	ret := string(mus.cur)
	for i := 0; i < len(mus.cur); i++ {
		mus.cur[i]++
		if mus.cur[i] != 0 {
			break
		}
	}
	if mus.cur[len(mus.cur)-1] == 0 {
		mus.cur = append(mus.cur, byte(1))
	}
	return ret
}

func (mus *mapsUUIDShortener) Extend(shortenedUUID string, prefixToAdd string) string {
	if val, found := mus.m[shortenedUUID]; found {
		return prefixToAdd + val
	}
	cur := mus.getNextCur()
	mus.m[cur] = shortenedUUID
	return prefixToAdd + shortenedUUID
}

func (mus *mapsUUIDShortener) Short(uuid string, prefixLengthToRemove int) string {
	uuid = uuid[prefixLengthToRemove:]
	if val, found := mus.m[uuid]; found {
		return val
	}
	cur := mus.getNextCur()
	mus.m[uuid] = cur
	return uuid
}
