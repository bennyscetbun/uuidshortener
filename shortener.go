package uuidshortener

import "bytes"

//We will use a map until i found a better struct
type prefixDictionnary struct {
	next    sortedPrefixDictionnary
	curChar byte
	val     string
}

type UUIDShortener interface {
	Extend(shortenedUUID string, prefixToAdd string) string
	Short(uuid string, prefixLengthToRemove int) string
}

type uuidShortener struct {
	prefixDictionnary
}

func New() UUIDShortener {
	return &uuidShortener{}
}

func (pd *prefixDictionnary) get(uuid string, curKey *bytes.Buffer) string {
	if len(uuid) == 0 {
		curKey.WriteString(pd.val)
		return curKey.String()
	}
	firstByte := uuid[0]
	uuid = uuid[1:]
	index, found := pd.next.Index(firstByte)
	if found {
		curKey.WriteByte(firstByte)
		if pd.next[index].val == uuid {
			return curKey.String()
		}
		return pd.next[index].get(uuid, curKey)
	}
	next := &prefixDictionnary{
		curChar: firstByte,
		val:     uuid,
	}
	pd.next.InsertAt(next, index)
	curKey.WriteByte(firstByte)
	curKey.WriteString(uuid)
	return curKey.String()
}

func (us *uuidShortener) Extend(shortenedUUID string, prefixToAdd string) string {
	keyBuffer := bytes.Buffer{}
	keyBuffer.WriteString(prefixToAdd)
	return us.get(shortenedUUID, &keyBuffer)
}

func (us *uuidShortener) Short(uuid string, prefixLengthToRemove int) string {
	keyBuffer := bytes.Buffer{}
	return us.get(uuid[prefixLengthToRemove:], &keyBuffer)
}
