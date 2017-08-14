package uuidshortener

/*
func TestConstantShortener(t *testing.T) {
	var uuids = []string{
		"u-aaaaaaaaaaaa",
		"u-baaaaaaaaaaa",
		"u-abaaaaaaaaaa",
		"u-aaaaaaaaaaab",
	}
	sender := NewConstantLengthShortener()
	receiver := NewConstantLengthExtender()

	// First time
	for _, uuid := range uuids {
		tosend := sender.Short(uuid, 2)
		if tosend != uuid[2:] {
			t.Errorf("first-time uuid is not sent correctly: %s,%s", uuid, tosend)
		}
		received := receiver.Extend(tosend, "u-")
		if uuid != received {
			t.Errorf("Received and sent uuid differ: %s,%s", uuid, received)
		}
	}
	for _, uuid := range uuids {
		tosend := sender.Short(uuid, 2)
		if tosend == uuid[2:] {
			t.Errorf("not-first-time uuid is not shortened: %s,%s", uuid, tosend)
		}
		received := receiver.Extend(tosend, "u-")
		if uuid != received {
			t.Errorf("Received and sent uuid differ: %s,%s", uuid, received)
		}
	}
}
*/
