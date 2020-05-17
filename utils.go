package ctxutils

import (
	"context"
	"fmt"
	"unsafe"
)

// Values - represents values stored in context.
type Values = map[interface{}]interface{}

// AddValues - shortcut for adding multiple values to context.
func AddValues(ctx context.Context, values Values) context.Context {
	for key, value := range values {
		ctx = context.WithValue(ctx, key, value)
	}
	return ctx
}

// DumpValues - dumps all values from context.
func DumpValues(ctx context.Context) Values {
	values := Values{}
Loop:
	for ctx != nil {
		switch fmt.Sprintf("%T", ctx) {
		case "*context.valueCtx":
			v := *(*valueCtx)((*iface)(unsafe.Pointer(&ctx)).data)
			values[v.key] = v.value
			ctx = v.Context
		case "*context.timerCtx", "*context.cancelCtx":
			v := *(*parentCtx)((*iface)(unsafe.Pointer(&ctx)).data)
			ctx = v.Context
		default:
			break Loop
		}
	}
	return values
}
