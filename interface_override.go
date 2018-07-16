package jsonschema

import "reflect"

type InterfaceOverride interface {
	// Sets AndOneOf interface override
	// target - struct to be overridden
	SetInterfaceAndOneOf(target interface{}, andOneOf AndOneOf) error
	// Returns interface AndOneOf to be overridden
	GetInterfaceAndOneOf(targetType reflect.Type) AndOneOf

	// Sets OneOf interface override
	// target - struct to be overridden
	SetInterfaceOneOf(target interface{}, oneOf OneOf) error
	// Returns interface OneOf to be overridden
	GetInterfaceOneOf(targetType reflect.Type) OneOf

	// Sets IfThenElse interface override
	// target - struct to be overridden
	SetInterfaceIfThenElse(target interface{}, ifThenElse IfThenElse) error
	// Returns interface IfThenElse to be overridden
	GetInterfaceIfThenElse(targetType reflect.Type) IfThenElse

	// Sets SchemaCase interface override
	// target - struct to be overridden
	SetInterfaceSchemaCase(target interface{}, schemaCase SchemaCase) error
	// Returns interface SchemaCase to be overridden
	GetInterfaceSchemaCase(targetType reflect.Type) SchemaCase
}

type interfaceOverrides struct {
	andOneOf   map[reflect.Type]AndOneOf
	oneOf      map[reflect.Type]OneOf
	ifThenElse map[reflect.Type]IfThenElse
	schemaCase map[reflect.Type]SchemaCase
}

// Returns initialized InterfaceOverride interface
func GetInterfaceOverride() InterfaceOverride {
	return &interfaceOverrides{
		andOneOf:   make(map[reflect.Type]AndOneOf),
		oneOf:      make(map[reflect.Type]OneOf),
		ifThenElse: make(map[reflect.Type]IfThenElse),
		schemaCase: make(map[reflect.Type]SchemaCase),
	}
}

func (o *interfaceOverrides) SetInterfaceAndOneOf(target interface{}, andOneOf AndOneOf) error {
	ts := reflect.TypeOf(target)
	o.andOneOf[ts] = andOneOf
	return nil
}

func (o *interfaceOverrides) GetInterfaceAndOneOf(targetType reflect.Type) AndOneOf {
	return o.andOneOf[targetType]
}

func (o *interfaceOverrides) SetInterfaceOneOf(target interface{}, oneOf OneOf) error {
	ts := reflect.TypeOf(target)
	o.oneOf[ts] = oneOf
	return nil
}

func (o *interfaceOverrides) GetInterfaceOneOf(targetType reflect.Type) OneOf {
	return o.oneOf[targetType]
}

func (o *interfaceOverrides) SetInterfaceIfThenElse(target interface{}, ifThenElse IfThenElse) error {
	ts := reflect.TypeOf(target)
	o.ifThenElse[ts] = ifThenElse
	return nil
}

func (o *interfaceOverrides) GetInterfaceIfThenElse(targetType reflect.Type) IfThenElse {
	return o.ifThenElse[targetType]
}

func (o *interfaceOverrides) SetInterfaceSchemaCase(target interface{}, schemaCase SchemaCase) error {
	ts := reflect.TypeOf(target)
	o.schemaCase[ts] = schemaCase
	return nil
}

func (o *interfaceOverrides) GetInterfaceSchemaCase(targetType reflect.Type) SchemaCase {
	return o.schemaCase[targetType]
}
