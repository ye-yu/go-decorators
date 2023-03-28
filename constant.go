package decorators

type DecoratorType int

const (
	FunctionDecorator  DecoratorType = iota
	ParameterDecorator DecoratorType = iota
)
