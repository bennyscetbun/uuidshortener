package uuidshortener

import "testing"

func test(b *testing.B, sender UUIDShortener, receiver UUIDExtender, numberOfSend, numberToGenerate int) {
	b.StopTimer()
	uuids := generateUUIDs("", numberToGenerate)
	l := len(uuids)
	b.StartTimer()
	var uuid string
	for n := 0; n < b.N; n++ {
		uuid = uuids[n%l]
		for i := 0; i < numberOfSend; i++ {
			if final := receiver.Extend(sender.Short(uuid, 0), ""); final != uuid {
				b.Fatal("WRONG RESULT first", final, uuid)
			}
		}
	}

}

const nbToGenerate = 2000

func BenchmarkMaps10(b *testing.B) {
	test(b, NewMapShortener(), NewMapExtender(), 10, nbToGenerate)
}

func BenchmarkConstant10(b *testing.B) {
	test(b, NewConstantLengthShortener(), NewConstantLengthExtender(), 10, nbToGenerate)
}

func BenchmarkSlice10(b *testing.B) {
	test(b, NewSliceShortener(1000), NewSliceExtender(1000), 10, nbToGenerate)
}

func BenchmarkIndex10(b *testing.B) {
	test(b, NewIndexShortener(1000), NewIndexExtender(1000), 10, nbToGenerate)
}

func BenchmarkBasic10(b *testing.B) {
	test(b, NewShortener(), NewExtender(), 10, nbToGenerate)
}

func BenchmarkMaps100(b *testing.B) {
	test(b, NewMapShortener(), NewMapExtender(), 100, nbToGenerate)
}

func BenchmarkConstant100(b *testing.B) {
	test(b, NewConstantLengthShortener(), NewConstantLengthExtender(), 100, nbToGenerate)
}

func BenchmarkSlice100(b *testing.B) {
	test(b, NewSliceShortener(nbToGenerate), NewSliceExtender(nbToGenerate), 100, nbToGenerate)
}

func BenchmarkIndex100(b *testing.B) {
	test(b, NewIndexShortener(1000), NewIndexExtender(1000), 100, nbToGenerate)
}

func BenchmarkBasic100(b *testing.B) {
	test(b, NewShortener(), NewExtender(), 100, nbToGenerate)
}

func BenchmarkMaps10000(b *testing.B) {
	test(b, NewMapShortener(), NewMapExtender(), 10000, nbToGenerate)
}

func BenchmarkConstant10000(b *testing.B) {
	test(b, NewConstantLengthShortener(), NewConstantLengthExtender(), 10000, nbToGenerate)
}

func BenchmarkSlice10000(b *testing.B) {
	test(b, NewSliceShortener(nbToGenerate), NewSliceExtender(nbToGenerate), 10000, nbToGenerate)
}

func BenchmarkIndex10000(b *testing.B) {
	test(b, NewIndexShortener(1000), NewIndexExtender(1000), 10000, nbToGenerate)
}

func BenchmarkBasic10000(b *testing.B) {
	test(b, NewShortener(), NewExtender(), 10000, nbToGenerate)
}
