package benchmarks_test

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"

	m "github.com/MacroPower/marshal-bench-go/benchmarks"
	"github.com/kelindar/binary"
)

var (
	binStruct []byte
	binMap    []byte

	gobStruct []byte
	gobMap    []byte
)

func init() {
	binStruct, _ = binary.Marshal(StructExpected1000)
	fmt.Printf("Binary struct size: %d\n", len(binStruct))

	binMap, _ = binary.Marshal(MapExpected1000)
	fmt.Printf("Binary map size: %d\n", len(binMap))

	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	enc.Encode(StructExpected1000)
	gobStruct = buf.Bytes()
	fmt.Printf("Gob struct size: %d\n", len(gobStruct))

	buf2 := new(bytes.Buffer)
	enc2 := gob.NewEncoder(buf2)
	enc2.Encode(MapExpected1000)
	gobMap = buf2.Bytes()
	fmt.Printf("Gob map size: %d\n", len(gobMap))
}

func BenchmarkKelindarBinaryUnmarshalStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := m.StructModelSl{}
		binary.Unmarshal(binStruct, &s)
	}
}

func BenchmarkKelindarBinaryUnmarshalMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := m.MapModelSl{}
		binary.Unmarshal(binMap, &s)
	}
}

func BenchmarkKelindarBinaryMarshalStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		binary.Marshal(StructExpected1000)
	}
}

func BenchmarkKelindarBinaryMarshalMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		binary.Marshal(MapExpected1000)
	}
}

func BenchmarkGobMarshalStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := new(bytes.Buffer)
		gob.NewEncoder(buf).Encode(StructExpected1000)
	}
}

func BenchmarkGobMarshalMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := new(bytes.Buffer)
		gob.NewEncoder(buf).Encode(MapExpected1000)
	}
}

func BenchmarkGobUnmarshalStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := m.StructModelSl{}
		gob.NewDecoder(bytes.NewBuffer(gobStruct)).Decode(&s)
	}
}

func BenchmarkGobUnmarshalMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := m.MapModelSl{}
		gob.NewDecoder(bytes.NewBuffer(gobMap)).Decode(&s)
	}
}
