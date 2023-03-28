package decorators_test

import (
	"context"
	"fmt"
	"testing"
	"yeyu/decorators"
)

func TestDecorateFunction(t *testing.T) {
	var randomFunc = func(s string, ctx context.Context) {
		fmt.Printf("Got string: %s\n", s)
		isFunctionDecoratorValue := ctx.Value(decorators.FunctionDecorator)

		isFunctionDecorator, castableToBoolean := isFunctionDecoratorValue.(bool)

		if castableToBoolean && isFunctionDecorator {
			fmt.Println("Function has been decorated with a function decorator")
		} else {
			fmt.Println("Weird, not decorated...")
		}

	}

	decoratedFn, _ := decorators.DecorateFn(randomFunc)
	decoratedFn("hello", context.Background())
}
