package json

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
