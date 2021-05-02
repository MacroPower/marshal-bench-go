package benchmarks_test

import (
	"testing"

	m "github.com/MacroPower/marshal-bench-go/benchmarks"
	"github.com/kelindar/binary"
)

var (
	bin, _    = binary.Marshal(StructExpected)
	binMap, _ = binary.Marshal(MapExpected)
)

func BenchmarkKelindarBinaryUnmarshalStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := m.StructModel{}
		binary.Unmarshal(bin, &s)
	}
}

func BenchmarkKelindarBinaryUnmarshalStructPtr(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := m.StructModelPtr{}
		binary.Unmarshal(bin, &s)
	}
}

func BenchmarkKelindarBinaryUnmarshalMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := m.MapModel{}
		binary.Unmarshal(binMap, &s)
	}
}

func BenchmarkKelindarBinaryMarshalStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		binary.Marshal(StructExpected)
	}
}

func BenchmarkKelindarBinaryMarshalMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		binary.Marshal(MapExpected)
	}
}
