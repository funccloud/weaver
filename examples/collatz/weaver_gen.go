// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package main

import (
	"context"
	"errors"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:  "github.com/ServiceWeaver/weaver/examples/collatz/Even",
		Iface: reflect.TypeOf((*Even)(nil)).Elem(),
		Impl:  reflect.TypeOf(even{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return even_local_stub{impl: impl.(Even), tracer: tracer, doMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/collatz/Even", Method: "Do", Remote: false})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return even_client_stub{stub: stub, doMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/collatz/Even", Method: "Do", Remote: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return even_server_stub{impl: impl.(Even), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return even_reflect_stub{caller: caller}
		},
		RefData: "",
	})
	codegen.Register(codegen.Registration{
		Name:      "github.com/ServiceWeaver/weaver/Main",
		Iface:     reflect.TypeOf((*weaver.Main)(nil)).Elem(),
		Impl:      reflect.TypeOf(server{}),
		Listeners: []string{"collatz"},
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return main_local_stub{impl: impl.(weaver.Main), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any { return main_client_stub{stub: stub} },
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return main_server_stub{impl: impl.(weaver.Main), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return main_reflect_stub{caller: caller}
		},
		RefData: "⟦f95ad2dd:wEaVeReDgE:github.com/ServiceWeaver/weaver/Main→github.com/ServiceWeaver/weaver/examples/collatz/Odd⟧\n⟦987c175b:wEaVeReDgE:github.com/ServiceWeaver/weaver/Main→github.com/ServiceWeaver/weaver/examples/collatz/Even⟧\n⟦f3b62957:wEaVeRlIsTeNeRs:github.com/ServiceWeaver/weaver/Main→collatz⟧\n",
	})
	codegen.Register(codegen.Registration{
		Name:  "github.com/ServiceWeaver/weaver/examples/collatz/Odd",
		Iface: reflect.TypeOf((*Odd)(nil)).Elem(),
		Impl:  reflect.TypeOf(odd{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return odd_local_stub{impl: impl.(Odd), tracer: tracer, doMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/collatz/Odd", Method: "Do", Remote: false})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return odd_client_stub{stub: stub, doMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/collatz/Odd", Method: "Do", Remote: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return odd_server_stub{impl: impl.(Odd), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return odd_reflect_stub{caller: caller}
		},
		RefData: "",
	})
}

// weaver.InstanceOf checks.
var _ weaver.InstanceOf[Even] = (*even)(nil)
var _ weaver.InstanceOf[weaver.Main] = (*server)(nil)
var _ weaver.InstanceOf[Odd] = (*odd)(nil)

// weaver.Router checks.
var _ weaver.Unrouted = (*even)(nil)
var _ weaver.Unrouted = (*server)(nil)
var _ weaver.Unrouted = (*odd)(nil)

// Local stub implementations.

type even_local_stub struct {
	impl      Even
	tracer    trace.Tracer
	doMetrics *codegen.MethodMetrics
}

// Check that even_local_stub implements the Even interface.
var _ Even = (*even_local_stub)(nil)

func (s even_local_stub) Do(ctx context.Context, a0 int) (r0 int, err error) {
	// Update metrics.
	begin := s.doMetrics.Begin()
	defer func() { s.doMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.Even.Do", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Do(ctx, a0)
}

type main_local_stub struct {
	impl   weaver.Main
	tracer trace.Tracer
}

// Check that main_local_stub implements the weaver.Main interface.
var _ weaver.Main = (*main_local_stub)(nil)

type odd_local_stub struct {
	impl      Odd
	tracer    trace.Tracer
	doMetrics *codegen.MethodMetrics
}

// Check that odd_local_stub implements the Odd interface.
var _ Odd = (*odd_local_stub)(nil)

func (s odd_local_stub) Do(ctx context.Context, a0 int) (r0 int, err error) {
	// Update metrics.
	begin := s.doMetrics.Begin()
	defer func() { s.doMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.Odd.Do", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Do(ctx, a0)
}

// Client stub implementations.

type even_client_stub struct {
	stub      codegen.Stub
	doMetrics *codegen.MethodMetrics
}

// Check that even_client_stub implements the Even interface.
var _ Even = (*even_client_stub)(nil)

func (s even_client_stub) Do(ctx context.Context, a0 int) (r0 int, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.doMetrics.Begin()
	defer func() { s.doMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.Even.Do", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += 8
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.Int(a0)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.Int()
	err = dec.Error()
	return
}

type main_client_stub struct {
	stub codegen.Stub
}

// Check that main_client_stub implements the weaver.Main interface.
var _ weaver.Main = (*main_client_stub)(nil)

type odd_client_stub struct {
	stub      codegen.Stub
	doMetrics *codegen.MethodMetrics
}

// Check that odd_client_stub implements the Odd interface.
var _ Odd = (*odd_client_stub)(nil)

func (s odd_client_stub) Do(ctx context.Context, a0 int) (r0 int, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.doMetrics.Begin()
	defer func() { s.doMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.Odd.Do", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += 8
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.Int(a0)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.Int()
	err = dec.Error()
	return
}

// Note that "weaver generate" will always generate the error message below.
// Everything is okay. The error message is only relevant if you see it when
// you run "go build" or "go run".
var _ codegen.LatestVersion = codegen.Version[[0][20]struct{}](`

ERROR: You generated this file with 'weaver generate' v0.20.0 (codegen
version v0.20.0). The generated code is incompatible with the version of the
github.com/ServiceWeaver/weaver module that you're using. The weaver module
version can be found in your go.mod file or by running the following command.

    go list -m github.com/ServiceWeaver/weaver

We recommend updating the weaver module and the 'weaver generate' command by
running the following.

    go get github.com/ServiceWeaver/weaver@latest
    go install github.com/ServiceWeaver/weaver/cmd/weaver@latest

Then, re-run 'weaver generate' and re-build your code. If the problem persists,
please file an issue at https://github.com/ServiceWeaver/weaver/issues.

`)

// Server stub implementations.

type even_server_stub struct {
	impl    Even
	addLoad func(key uint64, load float64)
}

// Check that even_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*even_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s even_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Do":
		return s.do
	default:
		return nil
	}
}

func (s even_server_stub) do(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 int
	a0 = dec.Int()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Do(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Int(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

type main_server_stub struct {
	impl    weaver.Main
	addLoad func(key uint64, load float64)
}

// Check that main_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*main_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s main_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	default:
		return nil
	}
}

type odd_server_stub struct {
	impl    Odd
	addLoad func(key uint64, load float64)
}

// Check that odd_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*odd_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s odd_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Do":
		return s.do
	default:
		return nil
	}
}

func (s odd_server_stub) do(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 int
	a0 = dec.Int()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Do(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Int(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

// Reflect stub implementations.

type even_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that even_reflect_stub implements the Even interface.
var _ Even = (*even_reflect_stub)(nil)

func (s even_reflect_stub) Do(ctx context.Context, a0 int) (r0 int, err error) {
	err = s.caller("Do", ctx, []any{a0}, []any{&r0})
	return
}

type main_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that main_reflect_stub implements the weaver.Main interface.
var _ weaver.Main = (*main_reflect_stub)(nil)

type odd_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that odd_reflect_stub implements the Odd interface.
var _ Odd = (*odd_reflect_stub)(nil)

func (s odd_reflect_stub) Do(ctx context.Context, a0 int) (r0 int, err error) {
	err = s.caller("Do", ctx, []any{a0}, []any{&r0})
	return
}
