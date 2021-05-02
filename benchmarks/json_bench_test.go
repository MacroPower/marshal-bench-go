package benchmarks_test

import (
	"bytes"
	"compress/flate"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	m "github.com/MacroPower/marshal-bench-go/benchmarks"
	"github.com/mailru/easyjson"
)

var (
	flateStructStr string
	flateMapStr    string

	jsonStruct []byte
	jsonMap    []byte
)

func init() {
	buf := new(bytes.Buffer)
	w, _ := flate.NewWriter(buf, -2)
	easyjson.MarshalToWriter(StructExpected1000, w)
	w.Close()
	flateStructStr = buf.String()
	fmt.Printf("Flate struct size: %d\n", len(flateStructStr))

	buf2 := new(bytes.Buffer)
	w2, _ := flate.NewWriter(buf2, -2)
	easyjson.MarshalToWriter(MapExpected1000, w2)
	w2.Close()
	flateMapStr = buf2.String()
	fmt.Printf("Flate map size: %d\n", len(flateMapStr))

	jsonStruct, _ = json.Marshal(StructExpected1000)
	fmt.Printf("JSON struct size: %d\n", len(jsonStruct))

	jsonMap, _ = json.Marshal(MapExpected1000)
	fmt.Printf("JSON map size: %d\n", len(jsonMap))
}

func BenchmarkJsonUnmarshalStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := m.StructModelSl{}
		json.Unmarshal(jsonStruct, &s)
	}
}

func BenchmarkJsonUnmarshalStructPtr(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := m.StructModelPtrSl{}
		json.Unmarshal(jsonStruct, &s)
	}
}

func BenchmarkJsonUnmarshalMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mm := m.MapModelSl{}
		json.Unmarshal(jsonMap, &mm)
	}
}

func BenchmarkJsonMarshalStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		json.Marshal(StructExpected1000)
	}
}

func BenchmarkJsonMarshalMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		json.Marshal(MapExpected1000)
	}
}

func BenchmarkEasyjsonUnmarshalStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := m.StructModelSl{}
		easyjson.Unmarshal(jsonStruct, &s)
	}
}

func BenchmarkEasyjsonUnmarshalStructPtr(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := m.StructModelPtrSl{}
		easyjson.Unmarshal(jsonStruct, &s)
	}
}

func BenchmarkEasyjsonUnmarshalStructIntern(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := m.StructModelInternSl{}
		easyjson.Unmarshal(jsonStruct, &s)
	}
}

func BenchmarkEasyjsonUnmarshalStructReaderFlate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := m.StructModelSl{}
		rc := flate.NewReader(strings.NewReader(flateStructStr))
		easyjson.UnmarshalFromReader(rc, &s)
		rc.Close()
	}
}

func BenchmarkEasyjsonUnmarshalStructReaderFlateIntern(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := m.StructModelInternSl{}
		rc := flate.NewReader(strings.NewReader(flateStructStr))
		easyjson.UnmarshalFromReader(rc, &s)
		rc.Close()
	}
}

func BenchmarkEasyjsonUnmarshalMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mm := m.MapModelSl{}
		easyjson.Unmarshal(jsonMap, &mm)
	}
}

func BenchmarkEasyjsonUnmarshalMapReaderFlate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := m.MapModelSl{}
		rc := flate.NewReader(strings.NewReader(flateMapStr))
		easyjson.UnmarshalFromReader(rc, &s)
		rc.Close()
	}
}

func BenchmarkEasyjsonMarshalStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		easyjson.Marshal(StructExpected1000)
	}
}

func BenchmarkEasyjsonMarshalStructWriter(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := new(bytes.Buffer)
		easyjson.MarshalToWriter(StructExpected1000, buf)
	}
}

func BenchmarkEasyjsonMarshalStructWriterFlate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := new(bytes.Buffer)
		w, _ := flate.NewWriter(buf, -2)
		easyjson.MarshalToWriter(StructExpected1000, w)
	}
}

func BenchmarkEasyjsonMarshalMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		easyjson.Marshal(MapExpected1000)
	}
}
