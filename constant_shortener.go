package uuidshortener

import "bytes"

type constantUUIDShortener struct {
	prefixDictionnary
}

func NewConstantLengthShortener() UUIDShortener {
	return &constantUUIDShortener{}
}
func NewConstantLengthExtender() UUIDExtender {
	return &constantUUIDShortener{}
}

func (us *constantUUIDShortener) Extend(shortenedUUID string, prefixToAdd string) string {
	keyBuffer := bytes.Buffer{}
	keyBuffer.WriteString(prefixToAdd)
	return us.get(shortenedUUID, &keyBuffer)
}

func (us *constantUUIDShortener) Short(uuid string, prefixLengthToRemove int) string {
	keyBuffer := bytes.Buffer{}
	return us.get(uuid[prefixLengthToRemove:], &keyBuffer)
}
