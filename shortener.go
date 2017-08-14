package uuidshortener

import "bytes"

type UUIDShortener interface {
	Short(uuid string, prefixLengthToRemove int) string
}

type UUIDExtender interface {
	Extend(shortenedUUID string, prefixToAdd string) string
}

type uuidShortener struct {
	prefixDictionnary
}

// New returns a UUIDShortener that works with not constant uuid length
// uuid length must be between 1 and 256 included after removing the prefix
func NewShortener() UUIDShortener {
	return &uuidShortener{}
}

func NewExtender() UUIDExtender {
	return &uuidShortener{}
}

func (us *uuidShortener) Extend(shortenedUUID string, prefixToAdd string) string {
	keyBuffer := bytes.Buffer{}
	keyBuffer.WriteString(prefixToAdd)
	firstByte := shortenedUUID[0]
	index, found := us.next.Index(firstByte)
	var nextDict *prefixDictionnary
	if found {
		nextDict = us.next[index]
	} else {
		nextDict = &prefixDictionnary{
			curChar: firstByte,
		}
		us.next.InsertAt(nextDict, index)
	}
	return nextDict.get(shortenedUUID[1:], &keyBuffer)
}

func (us *uuidShortener) Short(uuid string, prefixLengthToRemove int) string {
	keyBuffer := bytes.Buffer{}
	uuid = uuid[prefixLengthToRemove:]
	firstByte := byte(len(uuid))
	keyBuffer.WriteByte(firstByte)
	index, found := us.next.Index(firstByte)
	var nextDict *prefixDictionnary
	if found {
		nextDict = us.next[index]
	} else {
		nextDict = &prefixDictionnary{
			curChar: firstByte,
		}
		us.next.InsertAt(nextDict, index)
	}

	return nextDict.get(uuid, &keyBuffer)
}
