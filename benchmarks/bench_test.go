package benchmarks_test

import (
	m "github.com/MacroPower/marshal-bench-go/benchmarks"
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

	StructExpected1000 m.StructModelSl

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

	MapExpected1000 m.MapModelSl
)

func init() {
	for i := 0; i < 1000; i++ {
		StructExpected1000 = append(StructExpected1000, StructExpected)
		MapExpected1000 = append(MapExpected1000, MapExpected)
	}
}
