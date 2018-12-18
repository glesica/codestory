package models

type Struct struct {
	FieldCount int      `json:"fieldCount"`
	FieldNames []string `json:"fieldNames"`
	File       string   `json:"file"`
}
