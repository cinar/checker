// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"strings"
	"sync"
)

// messageTag is the name of the field tag used to override the message of
// one or more of a field's own checks/normalizers, without touching
// locales or registering a custom checker. Its value is a semicolon-
// separated list of "name=message" pairs, where name is the bare checker
// name as written in the checkers/validate tag (no ":params"). For a
// slice/map field, a container-level override is prefixed with "@", the
// same convention the checkers tag itself uses.
const messageTag = "checkersMsg"

// parseMessages parses a messageTag value into a map of checker name to its
// overriding message. Empty entries (from a trailing/leading/doubled ";")
// are skipped; an entry with no "=" maps to an empty message.
func parseMessages(raw string) map[string]string {
	messages := make(map[string]string)

	for _, entry := range strings.Split(raw, ";") {
		entry = strings.TrimSpace(entry)
		if entry == "" {
			continue
		}

		name, message, _ := strings.Cut(entry, "=")
		messages[strings.TrimSpace(name)] = strings.TrimSpace(message)
	}

	return messages
}

// messagesCache caches the parseMessages result by its exact source
// messageTag string, for the same reason configCache does: a given tag
// value is only ever parsed once, no matter how many struct instances or
// fields share that exact string.
var messagesCache sync.Map

// getMessages returns the parsed messages for raw, parsing and caching it
// on first use.
func getMessages(raw string) map[string]string {
	if cached, ok := messagesCache.Load(raw); ok {
		return cached.(map[string]string)
	}

	messages := parseMessages(raw)

	actual, _ := messagesCache.LoadOrStore(raw, messages)

	return actual.(map[string]string)
}

// sliceMessageSplit is the parsed container/item split of a slice or map
// field's messageTag value.
type sliceMessageSplit struct {
	sliceMessages string
	itemMessages  string
}

// splitSliceMessageConfig splits a messageTag value into slice/map-level and
// item-level "name=message" entries, using the same "@" prefix convention
// splitSliceConfig uses for the checkers tag.
func splitSliceMessageConfig(raw string) (string, string) {
	sliceEntries := make([]string, 0)
	itemEntries := make([]string, 0)

	for _, entry := range strings.Split(raw, ";") {
		trimmed := strings.TrimSpace(entry)
		if trimmed == "" {
			continue
		}

		if strings.HasPrefix(trimmed, sliceConfigPrefix) {
			sliceEntries = append(sliceEntries, strings.TrimPrefix(trimmed, sliceConfigPrefix))
		} else {
			itemEntries = append(itemEntries, trimmed)
		}
	}

	return strings.Join(sliceEntries, ";"), strings.Join(itemEntries, ";")
}

// sliceMessageSplitCache caches sliceMessageSplit by its exact source
// messageTag string, mirroring sliceSplitCache.
var sliceMessageSplitCache sync.Map

// getSliceMessageSplit returns the sliceMessageSplit for raw, computing and
// caching it on first use.
func getSliceMessageSplit(raw string) sliceMessageSplit {
	if cached, ok := sliceMessageSplitCache.Load(raw); ok {
		return cached.(sliceMessageSplit)
	}

	sliceMessages, itemMessages := splitSliceMessageConfig(raw)
	split := sliceMessageSplit{sliceMessages: sliceMessages, itemMessages: itemMessages}

	actual, _ := sliceMessageSplitCache.LoadOrStore(raw, split)

	return actual.(sliceMessageSplit)
}
