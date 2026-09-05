// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import "context"

// PipelineStep is a single unit of work in a Pipeline: it validates (and
// may normalize) part of value, given ctx, and returns the name to report
// an error under -- matching CheckErrors' field-name keys -- and the
// error itself, if any. Build one with Field or Rule, or write a
// PipelineStep literal directly for full control.
type PipelineStep[T any] func(ctx context.Context, value *T) (name string, err error)

// Pipeline is a generic, programmatic, context-aware alternative to
// CheckStruct's struct-tag validation, for domain rules that need
// request-scoped state (a database uniqueness check, a tenant boundary,
// an auth claim) or that don't cleanly map to a single field's tag. It
// reuses the same CheckFunc primitives as struct-tag validation, via the
// Field and Rule step constructors, and is fully opt-in: it doesn't
// affect CheckStruct or the checkers/validate tag path at all, and the
// two can be combined freely on the same type.
type Pipeline[T any] struct {
	steps []PipelineStep[T]
}

// NewPipeline creates an empty Pipeline for type T.
func NewPipeline[T any]() *Pipeline[T] {
	return &Pipeline[T]{}
}

// Step appends one or more steps to the pipeline, in the order given, and
// returns the pipeline so calls can be chained.
func (p *Pipeline[T]) Step(steps ...PipelineStep[T]) *Pipeline[T] {
	p.steps = append(p.steps, steps...)
	return p
}

// Validate runs every step against value in order, collecting each
// failing step's error under its name, and reports whether all of them
// passed. Unlike a single field's own checker chain (which stops at the
// first failing checker), Validate always runs every step, mirroring
// CheckStruct's field-independent error collection: an API caller sees
// every problem in one response, not just the first one encountered.
func (p *Pipeline[T]) Validate(ctx context.Context, value *T) (CheckErrors, bool) {
	errs := make(CheckErrors)

	for _, step := range p.steps {
		if name, err := step(ctx, value); err != nil {
			errs[name] = err
		}
	}

	return errs, len(errs) == 0
}

// Field builds a PipelineStep that runs checks against a single field of
// T, addressed by accessor, normalizing it in place exactly like a
// struct-tag field's checkers chain does: checks run in order and stop at
// the first error, and any normalizer among them (Trim, Lower, ...)
// writes its result back through accessor before the next check runs.
func Field[T, V any](name string, accessor func(*T) *V, checks ...CheckFunc[V]) PipelineStep[T] {
	return func(_ context.Context, value *T) (string, error) {
		fieldValue := accessor(value)

		newValue, err := Check(*fieldValue, checks...)
		*fieldValue = newValue

		return name, err
	}
}

// Rule builds a PipelineStep from a whole-value, context-aware domain
// rule: a function that inspects value and ctx, returning an error if the
// rule is violated. Use it for checks that don't fit a single field's
// tag, or that need request-scoped state a struct tag has no way to
// receive, such as a database lookup or a tenant/auth claim carried on
// ctx.
func Rule[T any](name string, rule func(ctx context.Context, value *T) error) PipelineStep[T] {
	return func(ctx context.Context, value *T) (string, error) {
		return name, rule(ctx, value)
	}
}
