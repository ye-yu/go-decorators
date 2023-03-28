package decorators

import (
	"context"
	"fmt"
	"reflect"
)

var typeofContext = reflect.TypeOf((*context.Context)(nil)).Elem()

func DecorateFn[T any](fn T) (decoratedFn T, e error) {
	reflectedValue := reflect.ValueOf(fn)
	typeofFn := reflectedValue.Type()
	if typeofFn.Kind() != reflect.Func {
		e = fmt.Errorf("%w: %v", ErrNotFunc, fn)
		return fn, e
	}
	argTypes := make([]reflect.Type, typeofFn.NumIn())
	argTypeContext := make([]int, 0)
	for i := 0; i < typeofFn.NumIn(); i++ {
		inArg := typeofFn.In(i)
		argTypes[i] = inArg
		if inArg.Implements(typeofContext) {
			argTypeContext = append(argTypeContext, i)
		}
	}

	outTypes := make([]reflect.Type, typeofFn.NumOut())
	for i := 0; i < typeofFn.NumOut(); i++ {
		outTypes[i] = typeofFn.Out(i)
	}

	newReflectedFnType := reflect.FuncOf(argTypes, outTypes, typeofFn.IsVariadic())
	newReflectedFn := reflect.MakeFunc(newReflectedFnType, func(args []reflect.Value) (results []reflect.Value) {
		for _, indexOfArgTypeContext := range argTypeContext {
			reflectedValueOfTypeContext := args[indexOfArgTypeContext]
			reflectValueInterface := reflectedValueOfTypeContext.Interface()
			newContext := context.WithValue(reflectValueInterface.(context.Context), FunctionDecorator, true)
			newReflectedContext := reflect.New(reflectedValueOfTypeContext.Type())
			newReflectedContextElem := newReflectedContext.Elem()
			newReflectedContextElem.Set(reflect.ValueOf(newContext))
			args[indexOfArgTypeContext] = newReflectedContextElem
		}
		return reflectedValue.Call(args)
	})

	reflectedDecorated := reflect.New(typeofFn)
	reflectedDecoratedFn := reflectedDecorated.Elem()
	if !reflectedDecoratedFn.CanSet() {
		return fn, ErrCannotSetFunc
	}

	reflectedDecoratedFn.Set(newReflectedFn)
	reflectedDecoratedFnAsInterface := reflectedDecorated.Interface()
	decoratedFnPtr, castable := reflectedDecoratedFnAsInterface.(*T)

	if !castable {
		return fn, ErrCannotCastFunc
	}

	decoratedFn = *decoratedFnPtr
	return decoratedFn, nil
}
