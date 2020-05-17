package ctxutils_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/GreenHedgehog/ctxutils"
)

func TestDumpValues(t *testing.T) {
	tests := []struct {
		desc   string
		ctx    context.Context
		values ctxutils.Values
	}{
		{
			desc:   "Context backgpound",
			ctx:    context.Background(),
			values: ctxutils.Values{},
		},
		{
			desc:   "Context with cancel",
			ctx:    withCancel(context.Background()),
			values: ctxutils.Values{},
		},
		{
			desc:   "Context with timeout",
			ctx:    withTimeout(context.Background()),
			values: ctxutils.Values{},
		},
		{
			desc:   "Context with cancel",
			ctx:    withDeadline(context.Background()),
			values: ctxutils.Values{},
		},
		{
			desc: "One key-value pair",
			ctx:  context.WithValue(context.Background(), "key1", "value1"),
			values: ctxutils.Values{
				"key1": "value1",
			},
		},
		{
			desc: "Multiple key-value pairs",
			ctx: context.WithValue(
				context.WithValue(
					context.WithValue(
						context.Background(), "key1", "value1"),
					"key2", "value2"),
				"key3", "value3"),
			values: ctxutils.Values{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
			},
		},
		{
			desc: "Layered context types",
			ctx: withDeadline(context.WithValue(
				withTimeout(context.WithValue(
					withCancel(context.WithValue(
						context.Background(), "key1", "value1"),
					), "key2", "value2"),
				), "key3", "value3"),
			),
			values: ctxutils.Values{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
			},
		},
	}

	for number, test := range tests {
		values := ctxutils.DumpValues(test.ctx)
		if !reflect.DeepEqual(values, test.values) {
			t.Fatalf("Test #%d (%s): expected - %+v, got - %+v", number+1, test.desc, test.values, values)
		}
	}
}

var (
	withCancel = func(ctx context.Context) context.Context {
		ctx, _ = context.WithCancel(ctx)
		return ctx
	}
	withTimeout = func(ctx context.Context) context.Context {
		ctx, _ = context.WithTimeout(ctx, time.Second)
		return ctx
	}
	withDeadline = func(ctx context.Context) context.Context {
		ctx, _ = context.WithDeadline(ctx, time.Now())
		return ctx
	}
)
