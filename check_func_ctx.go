// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"sync"
)

// CheckFuncCtx is a context-aware CheckFunc: like CheckFunc, but also
// receives a context.Context, for checks that need cancellation, a
// deadline, or a request-scoped value (a DB session, a tenant claim) that a
// struct tag alone has no way to carry.
type CheckFuncCtx[T any] func(ctx context.Context, value T) (T, error)

// MakeCheckFuncCtx is a function that returns a context-aware check
// function using the given params, the checkersCtx-tag counterpart of
// MakeCheckFunc.
type MakeCheckFuncCtx func(params string) CheckFuncCtx[reflect.Value]

// ctxTag is the name of the field tag used for context-aware checks
// registered through RegisterCtxMaker. It's independent of checkerTag: a
// checkersCtx entry only ever runs through CheckStructWithContext, since
// CheckStruct and CheckWithConfig have no context.Context to give it and
// leave the tag untouched.
const ctxTag = "checkersCtx"

// ctxMakersMu guards ctxMakers, for the same reason makersMu guards makers:
// RegisterCtxMaker can run concurrently with CheckStructWithContext calls.
var ctxMakersMu sync.RWMutex

// ctxMakers provides a mapping of context-aware maker functions keyed by
// check name. There are no built-in context-aware checkers -- every
// built-in checker is already expressible without one -- so this only ever
// holds makers added through RegisterCtxMaker.
var ctxMakers = map[string]MakeCheckFuncCtx{}

// RegisterCtxMaker registers a new context-aware maker function with the
// given name, for use in a field's checkersCtx tag.
func RegisterCtxMaker(name string, maker MakeCheckFuncCtx) {
	ctxMakersMu.Lock()
	defer ctxMakersMu.Unlock()

	ctxMakers[name] = maker
}

// RegisteredCtxMakerNames returns the name of every currently registered
// context-aware checker maker added via RegisterCtxMaker. There are no
// built-ins, so the result is empty until at least one is registered. The
// order is not significant.
func RegisteredCtxMakerNames() []string {
	ctxMakersMu.RLock()
	defer ctxMakersMu.RUnlock()

	names := make([]string, 0, len(ctxMakers))
	for name := range ctxMakers {
		names = append(names, name)
	}

	return names
}

// CheckWithContext applies the given context-aware check functions to a
// value sequentially, threading ctx through each one and stopping at the
// first error, exactly like Check.
func CheckWithContext[T any](ctx context.Context, value T, checks ...CheckFuncCtx[T]) (T, error) {
	var err error

	for _, check := range checks {
		value, err = check(ctx, value)
		if err != nil {
			break
		}
	}

	return value, err
}

// runCtxChecks runs value through config's space-separated checkersCtx
// tokens against ctx, in order, stopping at the first error. Panics if a
// token names a checker with no registered ctx maker, the same way an
// unknown name in the checkers tag panics.
func runCtxChecks(ctx context.Context, value reflect.Value, config string) (reflect.Value, error) {
	var err error

	for _, token := range strings.Fields(config) {
		name, params, _ := strings.Cut(token, ":")

		ctxMakersMu.RLock()
		maker, ok := ctxMakers[name]
		ctxMakersMu.RUnlock()

		if !ok {
			panic(fmt.Sprintf("unknown ctx checker %q", name))
		}

		value, err = maker(params)(ctx, value)
		if err != nil {
			break
		}
	}

	return value, err
}
