package main

type Function struct {
	Arity      int    `json:"arity"`
	Complexity int    `json:"complexity"`
	Lines      int    `json:"lines"`
	Name       string `json:"name"`
	Panics     int    `json:"panics"`
}
