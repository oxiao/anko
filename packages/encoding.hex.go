package packages

import (
	"encoding/hex"
	"github.com/mattn/anko/env"
	"reflect"
)

func init() {
	env.Packages["hex"] = map[string]reflect.Value{
		"Decode":         reflect.ValueOf(hex.Decode),
		"DecodeString":   reflect.ValueOf(hex.DecodeString),
		"DecodedLen":     reflect.ValueOf(hex.DecodedLen),
		"Dump":           reflect.ValueOf(hex.Dump),
		"Dumper":         reflect.ValueOf(hex.Dumper),
		"Encode":         reflect.ValueOf(hex.Encode),
		"EncodeToString": reflect.ValueOf(hex.EncodeToString),
		"EncodedLen":     reflect.ValueOf(hex.EncodedLen),
		"ErrLength":      reflect.ValueOf(hex.ErrLength),
		"NewDecoder":     reflect.ValueOf(hex.NewDecoder),
		"NewEncoder":     reflect.ValueOf(hex.NewEncoder),
	}
}
