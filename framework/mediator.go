package framework

import (
	"errors"
	"reflect"
	"sync"
)

var (
	registeredHandlers sync.Map
)

func init() {
	registeredHandlers = sync.Map{}
}

type key[TRequest any, TResult any] struct {
}

// Register registers the provided request handler to be used for the corresponding requests
func Register[TRequest any, TResult any](handler RequestHandler[TRequest, TResult]) error {
	k := key[TRequest, TResult]{}
	_, existed := registeredHandlers.LoadOrStore(reflect.TypeOf(k), handler)
	if existed {
		return errors.New("the provided type is already registered to a handler")
	}
	return nil
}

// Send processes the provided request and returns the produced result
func Send[TRequest any, TResult any](r TRequest) (TResult, error) {
	var zeroRes TResult
	var k key[TRequest, TResult]
	handler, ok := registeredHandlers.Load(reflect.TypeOf(k))
	if !ok {
		return zeroRes, errors.New("could not find zeroRes handler for this function")
	}
	switch handler := handler.(type) {
	case RequestHandler[TRequest, TResult]:
		return handler.Handle(r)
	}
	return zeroRes, errors.New("Invalid handler")
}

// RequestHandler handles TRequest and returns TResult
type RequestHandler[TRequest any, TResult any] interface {
	Handle(request TRequest) (TResult, error)
}
