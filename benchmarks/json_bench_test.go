package benchmarks_test

import (
	"encoding/json"
	"testing"

	m "github.com/MacroPower/marshal-bench-go/benchmarks"
	"github.com/mailru/easyjson"
)

func BenchmarkJsonUnmarshalStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := m.StructModel{}
		json.Unmarshal(JsonData, &s)
	}
}

func BenchmarkJsonUnmarshalStructPtr(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := m.StructModelPtr{}
		json.Unmarshal(JsonData, &s)
	}
}

func BenchmarkJsonUnmarshalMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mm := m.MapModel{}
		json.Unmarshal(JsonData, &mm)
	}
}

func BenchmarkJsonMarshalStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		json.Marshal(StructExpected)
	}
}

func BenchmarkJsonMarshalMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		json.Marshal(MapExpected)
	}
}

func BenchmarkEasyjsonUnmarshalStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := m.StructModel{}
		easyjson.Unmarshal(JsonData, &s)
	}
}

func BenchmarkEasyjsonUnmarshalStructPtr(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := m.StructModelPtr{}
		easyjson.Unmarshal(JsonData, &s)
	}
}

func BenchmarkEasyjsonUnmarshalStructIntern(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := m.StructModelIntern{}
		easyjson.Unmarshal(JsonData, &s)
	}
}

func BenchmarkEasyjsonUnmarshalMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mm := m.MapModel{}
		easyjson.Unmarshal(JsonData, &mm)
	}
}

func BenchmarkEasyjsonMarshalStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		easyjson.Marshal(StructExpected)
	}
}

func BenchmarkEasyjsonMarshalMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		easyjson.Marshal(MapExpected)
	}
}
