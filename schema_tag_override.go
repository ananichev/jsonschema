package jsonschema

import (
	"fmt"
	"reflect"
)

// SchemaTagOverride is a mechanism to allow jsonschema tag overrides
type SchemaTagOverride interface {
	// Set should be given:
	// targetStruct - struct that contains the field to be overridden
	// targetField - name of the field that is to be overridden
	// tag - the provided jsonschema tag
	Set(targetStruct interface{}, targetField string, tag string) error
	// Get is used by this library to retrieve overrides
	Get(targetStructType reflect.Type, targetField string) string
	// SetStructType should be given:
	// jsonSchemaType - simple type of json schema
	// targetStruct - struct that contains the field to be overridden
	// tag - the provided jsonschema tag
	SetStructType(targetStruct interface{}, jsonSchemaType string, tag string) error
	// Get is used by this library to retrieve overrides struct types
	// Returns type and tag
	GetStructType(targetStructType reflect.Type) (string, string)
}

// GetSchemaTagOverride returns initialized SchemaTagOverride
func GetSchemaTagOverride() SchemaTagOverride {
	c := make(map[reflect.Type]map[string]string)
	structConfig := make(map[reflect.Type]string)
	structTypesConfig := make(map[reflect.Type]string)
	return &overrides{config: c, structConfig: structConfig, structTypesConfig: structTypesConfig}
}

type overrides struct {
	config map[reflect.Type]map[string]string
	structConfig map[reflect.Type]string
	structTypesConfig map[reflect.Type]string
}

// Set adds a jsonschema tag override to internal map
func (o *overrides) Set(targetStruct interface{}, targetField string, tag string) error {
	ts := reflect.TypeOf(targetStruct)

	if k := ts.Kind(); k != reflect.Struct {
		return fmt.Errorf("expecting struct, got %s instead", reflect.Kind(k))
	}

	if _, ok := ts.FieldByName(targetField); !ok {
		return fmt.Errorf("targetStruct %s does not have field %s", ts.Name(), targetField)
	}

	if o.config[ts] != nil {
		o.config[ts][targetField] = tag

		return nil
	}

	o.config[ts] = map[string]string{targetField: tag}

	return nil
}

// Get retrieves tags from internal map
func (o *overrides) Get(targetStructType reflect.Type, targetField string) string {
	if targetStructType.Kind() != reflect.Struct {
		return ""
	}

	if o.config[targetStructType] == nil {
		return ""
	}

	return o.config[targetStructType][targetField]
}


// Set adds a jsonschema tag override struct usage
func (o *overrides) SetStructType(targetStruct interface{}, structType string, tag string) error {
	if jsonTypeErr := validateJsonType(structType); jsonTypeErr != nil {
		return jsonTypeErr
	}
	ts := reflect.TypeOf(targetStruct)
	o.structTypesConfig[ts] = structType
	o.structConfig[ts] = tag
	return nil
}

// Get is used by this library to retrieve overrides struct types where they are used
func (o *overrides) GetStructType(targetStructType reflect.Type) (string, string) {
	return o.structTypesConfig[targetStructType], o.structConfig[targetStructType]
}