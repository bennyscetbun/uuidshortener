package uuidshortener

import "testing"

var uuids = []string{
	"u-aaaaaaaaaaaa",
	"u-baaaaaaaaaaa",
	"u-abaaaaaaaaaa",
	"u-aaaaaaaaaaab",
}

func TestShortener(t *testing.T) {
	sender := New()
	receiver := New()

	// First time
	for _, uuid := range uuids {
		tosend := sender.Short(uuid, 2)
		if len(tosend)+2 != len(uuid) {
			t.Errorf("first-time uuid is not sent correctly: %s,%s", uuid, tosend)
		}
		received := receiver.Extend(tosend, "u-")
		if uuid != received {
			t.Errorf("Received and sent uuid differ: %s,%s", uuid, received)
		}
	}
	for _, uuid := range uuids {
		tosend := sender.Short(uuid, 2)
		if len(tosend) >= len(uuid)-2 {
			t.Errorf("not-first-time uuid is not shortened: %s,%s", uuid, tosend)
		}
		received := receiver.Extend(tosend, "u-")
		if uuid != received {
			t.Errorf("Received and sent uuid differ: %s,%s", uuid, received)
		}
	}
}
