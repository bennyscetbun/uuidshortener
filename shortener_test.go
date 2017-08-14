package uuidshortener

import (
	"math/rand"
	"testing"
)

const (
	uuidSpace = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func generateUUIDs(prefix string, nb int) []string {
	ret := make([]string, nb)
	r := rand.New(rand.NewSource(42))
	for i := 0; i < len(ret); i++ {
		data := make([]byte, 32+len(prefix))
		for j := 0; j < len(prefix); j++ {
			data[j] = prefix[j]
		}
		r.Read(data[len(prefix):])
		for i, b := range data[2:] {
			data[i+2] = byte(uuidSpace[int(b)%len(uuidSpace)])
		}
		ret[i] = string(data)
	}
	return ret
}

func shortener(uuids []string, t *testing.T) {
	sender := NewShortener()
	receiver := NewExtender()

	// First time
	for _, uuid := range uuids {
		tosend := sender.Short(uuid, 2)
		if tosend[0] != byte(len(uuid[2:])) || tosend[1:] != uuid[2:] {
			t.Errorf("first-time uuid is not sent correctly: %s,%s, %c", uuid, tosend, tosend[0])
		}
		received := receiver.Extend(tosend, "u-")
		if uuid != received {
			t.Errorf("Received and sent uuid differ: %s,%s", uuid, received)
		}
	}
	for _, uuid := range uuids {
		tosend := sender.Short(uuid, 2)
		if len(uuid[2:]) != int(tosend[0]) || tosend[1:] == uuid[2:] {
			t.Errorf("not-first-time uuid is not shortened: %s,%s, %d", uuid, tosend, int(tosend[0]))
		}
		received := receiver.Extend(tosend, "u-")

		if uuid != received {
			t.Errorf("Received and sent uuid differ: %s,%s", uuid, received)
		}
	}
}

func TestShortener(t *testing.T) {
	var uuids = []string{
		"u-aaaaaaaaaaaa",
		"u-baaaaaaaaaaa",
		"u-abaaaaaaaaaa",
		"u-aaaaaaaaaaab",
		"u-aaa",
		"u-aaaaa",
		"u-baa",
		"u-aa",
		"u-ba",
	}
	shortener(uuids, t)
	shortener(generateUUIDs("u-", 2000), t)
}
