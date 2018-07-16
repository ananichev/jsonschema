package jsonschema

import (
	"errors"
	"strings"
)

var jsonSchemaSimpleTypes = []string{"string", "number", "object", "array", "boolean", "null"}

func validateJsonType(needle string) error {
	for _, jsonSchemaType := range jsonSchemaSimpleTypes {
		if jsonSchemaType == needle {
			return nil
		}
	}
	return errors.New("Unsupported struct type " + needle + " Supported types: " + strings.Join(jsonSchemaSimpleTypes, ","))
}
