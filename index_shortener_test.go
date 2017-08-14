package uuidshortener

import "testing"

func indexShortener(uuids []string, t *testing.T) {
	sender := NewIndexShortener(len(uuids))
	receiver := NewIndexExtender(len(uuids))

	// First time
	for _, uuid := range uuids {
		tosend := sender.Short(uuid, 2)
		if tosend[0] != indexString && tosend[1:] != uuid[2:] {
			t.Errorf("first-time uuid is not sent correctly: %s,%s, %c", uuid, tosend, tosend[0])
		}
		received := receiver.Extend(tosend, "u-")
		if uuid != received {
			t.Errorf("Received and sent uuid differ first time: %s,%s", uuid, received)
		}
	}
	for _, uuid := range uuids {
		tosend := sender.Short(uuid, 2)
		if tosend[0] != indexNumber || tosend[1:] == uuid[2:] {
			t.Errorf("not-first-time uuid is not shortened: %s,%s", uuid, tosend)
		}
		received := receiver.Extend(tosend, "u-")

		if uuid != received {
			t.Errorf("Received and sent uuid differ: %s,%s", uuid, received)
		}
	}
}

func TestIndexShortener(t *testing.T) {
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
	indexShortener(uuids, t)
	indexShortener(generateUUIDs("u-", 2000), t)
}
