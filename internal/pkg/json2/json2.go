package json2

import (
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.Config{
	EscapeHTML:             true,
	SortMapKeys:            true,
	ValidateJsonRawMessage: true,
	TagKey:                 "json2",
}.Froze()

// Marshal with tag `json2`
func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal with tag `json2`
func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// Copy copy from src to dist by tagkey: json2
func Copy(dist, src interface{}) error {
	b, err := Marshal(src)
	if err != nil {
		return err
	}
	return Unmarshal(b, dist)
}
