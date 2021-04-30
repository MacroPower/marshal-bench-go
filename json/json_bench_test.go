package json_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	m "github.com/MacroPower/marshal-bench-go/json"
	"github.com/mailru/easyjson"
)

var (
	StructExpected = m.StructModel{
		A: "6BEJpen653OVhFmRgj3C",
		B: "jLQxPUdQl0SyS0cKqpgj",
		C: "21TaM8M4QSpTk9dfDoAE",
		D: "CtyreDDDw6Y19EG0ITH9",
		E: "P7DKaskrDDlb0u6fMojb",
		F: "SEqvT98zoRP7UchGhMhf",
		G: "hX7RlfwTqg4pxIxjtvft",
		H: "YRBQhg5fkYW9R8dQAKiS",
		I: "GdIm3jF2AO3yHxiYp994",
		J: "woGAX0x4LWUMGnV6IEeU",
		K: "uXRIn7VGNHZdZwJQRP0B",
		L: "liAdFAC5eX1ZHI9pNTXu",
		M: "vjq8eQqLRzCVLhzA8Z4s",
		N: "7mrVocwc6MHHiNQrfYBY",
		O: "G9jQSV6Xt19gQGTVU8at",
		P: "MVohu5lsD7qG4XjxkUXC",
		Q: "GzRrlMmadyim496JYuYm",
		R: "iLSkEqup5G2mmJvkslEe",
		S: "zfNmYU4FLzXtDan9w6y6",
		T: "oVdp5dC7nzxS2VMRFE6S",
	}

	MapExpected = m.MapModel{
		"A": "6BEJpen653OVhFmRgj3C",
		"B": "jLQxPUdQl0SyS0cKqpgj",
		"C": "21TaM8M4QSpTk9dfDoAE",
		"D": "CtyreDDDw6Y19EG0ITH9",
		"E": "P7DKaskrDDlb0u6fMojb",
		"F": "SEqvT98zoRP7UchGhMhf",
		"G": "hX7RlfwTqg4pxIxjtvft",
		"H": "YRBQhg5fkYW9R8dQAKiS",
		"I": "GdIm3jF2AO3yHxiYp994",
		"J": "woGAX0x4LWUMGnV6IEeU",
		"K": "uXRIn7VGNHZdZwJQRP0B",
		"L": "liAdFAC5eX1ZHI9pNTXu",
		"M": "vjq8eQqLRzCVLhzA8Z4s",
		"N": "7mrVocwc6MHHiNQrfYBY",
		"O": "G9jQSV6Xt19gQGTVU8at",
		"P": "MVohu5lsD7qG4XjxkUXC",
		"Q": "GzRrlMmadyim496JYuYm",
		"R": "iLSkEqup5G2mmJvkslEe",
		"S": "zfNmYU4FLzXtDan9w6y6",
		"T": "oVdp5dC7nzxS2VMRFE6S",
	}

	JsonData []byte
)

func init() {
	content, err := ioutil.ReadFile("testdata/strings_flat_small.json")
	if err != nil {
		panic(err)
	}
	JsonData = content
}

func BenchmarkJsonUnmarshalStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := m.StructModel{}
		json.Unmarshal(JsonData, &s)
		if s.T != StructExpected.T {
			panic("incorrect data")
		}
	}
}

func BenchmarkJsonUnmarshalStructPtr(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := m.StructModelPtr{}
		json.Unmarshal(JsonData, &s)
		if *s.T != StructExpected.T {
			panic("incorrect data")
		}
	}
}

func BenchmarkJsonUnmarshalMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mm := m.MapModel{}
		json.Unmarshal(JsonData, &mm)
		if mm["T"] != MapExpected["T"] {
			panic("incorrect data")
		}
	}
}

func BenchmarkJsonMarshalStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		bt, err := json.Marshal(StructExpected)
		if err != nil {
			panic(err)
		}
		js := string(bt)
		if len(js) != 541 {
			panic("incorrect data")
		}
	}
}

func BenchmarkJsonMarshalMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		bt, err := json.Marshal(MapExpected)
		if err != nil {
			panic(err)
		}
		js := string(bt)
		if len(js) != 541 {
			panic("incorrect data")
		}
	}
}

func BenchmarkEasyjsonUnmarshalStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := m.StructModel{}
		easyjson.Unmarshal(JsonData, &s)
		if s.T != StructExpected.T {
			panic("incorrect data")
		}
	}
}

func BenchmarkEasyjsonUnmarshalStructPtr(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := m.StructModelPtr{}
		easyjson.Unmarshal(JsonData, &s)
		if *s.T != StructExpected.T {
			panic("incorrect data")
		}
	}
}

func BenchmarkEasyjsonUnmarshalStructIntern(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := m.StructModelIntern{}
		easyjson.Unmarshal(JsonData, &s)
		if s.T != StructExpected.T {
			panic("incorrect data")
		}
	}
}

func BenchmarkEasyjsonUnmarshalMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mm := m.MapModel{}
		easyjson.Unmarshal(JsonData, &mm)
		if mm["T"] != MapExpected["T"] {
			panic("incorrect data")
		}
	}
}

func BenchmarkEasyjsonMarshalStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		bt, err := easyjson.Marshal(StructExpected)
		if err != nil {
			panic(err)
		}
		js := string(bt)
		if len(js) != 541 {
			panic("incorrect data")
		}
	}
}

func BenchmarkEasyjsonMarshalMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		bt, err := easyjson.Marshal(MapExpected)
		if err != nil {
			panic(err)
		}
		js := string(bt)
		if len(js) != 541 {
			panic("incorrect data")
		}
	}
}
