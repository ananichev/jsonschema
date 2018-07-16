package jsonschema

import "reflect"

type andOneOfFunc struct {
	handler func() []reflect.StructField
}

func (a *andOneOfFunc) AndOneOf() []reflect.StructField {
	return a.handler()
}

func NewAndOneOfFunc(handler func() []reflect.StructField) AndOneOf {
	return &andOneOfFunc{
		handler: handler,
	}
}

type oneOfFunc struct {
	handler func() []reflect.StructField
}

func NewOneOfFunc(handler func() []reflect.StructField) OneOf {
	return &oneOfFunc{
		handler: handler,
	}
}

func (a *oneOfFunc) OneOf() []reflect.StructField {
	return a.handler()
}

type ifThenElseFunc struct {
	handler func() SchemaCondition
}

func NewIfThenElseFunc(handler func() SchemaCondition) IfThenElse {
	return &ifThenElseFunc{
		handler: handler,
	}
}

func (a *ifThenElseFunc) IfThenElse() SchemaCondition {
	return a.handler()
}

type schemaCaseFunc struct {
	handler func() SchemaSwitch
}

func NewSchemaCaseFunc(handler func() SchemaSwitch) SchemaCase {
	return &schemaCaseFunc{handler: handler}
}

func (a *schemaCaseFunc) Case() SchemaSwitch {
	return a.handler()
}
