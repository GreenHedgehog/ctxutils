package ctxutils

import (
	"context"
	"reflect"
	"time"
	"unsafe"
)

var (
	_ctxValue     = context.WithValue(context.Background(), "key", "value")
	_ctxValueDump = *(*valueCtx)((*iface)(unsafe.Pointer(&_ctxValue)).data)

	_ctxCancel, _  = context.WithCancel(context.Background())
	_ctxCancelDump = *(*parentCtx)((*iface)(unsafe.Pointer(&_ctxCancel)).data)

	_ctxTimeout, _  = context.WithTimeout(context.Background(), time.Second)
	_ctxTimeoutDump = *(*parentCtx)((*iface)(unsafe.Pointer(&_ctxTimeout)).data)

	_ctxDeadline, _  = context.WithDeadline(context.Background(), time.Now())
	_ctxDeadlineDump = *(*parentCtx)((*iface)(unsafe.Pointer(&_ctxDeadline)).data)
)

func init() {
	switch {
	case
		_ctxValueDump.key != "key",
		_ctxValueDump.value != "value",
		!reflect.DeepEqual(_ctxValueDump.Context, context.Background()),
		!reflect.DeepEqual(_ctxCancelDump.Context, context.Background()),
		!reflect.DeepEqual(_ctxTimeoutDump.Context, context.Background()),
		!reflect.DeepEqual(_ctxDeadlineDump.Context, context.Background()):
		panic("wrong memory table")
	}
}

type (
	// substitute for `src/runtime.(iface)``
	iface struct {
		_    unsafe.Pointer
		data unsafe.Pointer
	}

	// substitute for of `src/context.(valueCtx)`
	valueCtx struct {
		context.Context
		key, value interface{}
	}

	// substitute for `src/context.(*timerCtx)` and `src/context.(*cancelCtx)`
	parentCtx struct {
		context.Context
		_ struct{}
	}
)
