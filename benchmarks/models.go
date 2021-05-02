package benchmarks

import _ "github.com/mailru/easyjson/gen"

//go:generate easyjson -all models.go

type StructModel struct {
	A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T string
}

type StructModelPtr struct {
	A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T *string
}

type StructModelIntern struct {
	A string `json:"A,intern"`
	B string `json:"B,intern"`
	C string `json:"C,intern"`
	D string `json:"D,intern"`
	E string `json:"E,intern"`
	F string `json:"F,intern"`
	G string `json:"G,intern"`
	H string `json:"H,intern"`
	I string `json:"I,intern"`
	J string `json:"J,intern"`
	K string `json:"K,intern"`
	L string `json:"L,intern"`
	M string `json:"M,intern"`
	N string `json:"N,intern"`
	O string `json:"O,intern"`
	P string `json:"P,intern"`
	Q string `json:"Q,intern"`
	R string `json:"R,intern"`
	S string `json:"S,intern"`
	T string `json:"T,intern"`
}

//easyjson:json
type MapModel map[string]string

var (
	StructExpected = StructModel{
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

	MapExpected = MapModel{
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
)
