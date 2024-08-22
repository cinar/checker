package checker

import (
	"fmt"
	"reflect"
	"strings"
)

type Rule struct {
	FieldIndex []int
	Checkers   []CheckFunc
}

// rulesRegistry stores a collection of rules associated with a specific type.
var rulesRegistry = map[reflect.Type][]*Rule{}

// checkersCache stores checkers associated with a specific config.
var checkersCache = map[string]CheckFunc{}

// RegisterRulesFromConfig registers rules based on the provided struct type and config map.
func RegisterRulesFromConfig(structType reflect.Type, nameToConfig map[string]string) error {
	rules := make([]*Rule, len(nameToConfig))

	for name, config := range nameToConfig {
		field, ok := structType.FieldByName(name)
		if !ok {
			return fmt.Errorf("field %s not found", name)
		}

		checkers, err := initCheckersFromConfig(config)
		if err != nil {
			return fmt.Errorf("field %s has errors: %w", name, err)
		}

		rules = append(rules, &Rule{
			FieldIndex: field.Index,
			Checkers:   checkers,
		})
	}

	rulesRegistry[structType] = rules

	return nil
}

// RegisterRulesFromTag registers rules based on the provided struct type and field tags.
func RegisterRulesFromTag(structType reflect.Type) error {
	structTypes := []reflect.Type{structType}

	for len(structTypes) > 0 {
		structType := structTypes[0]
		structTypes = structTypes[1:]

		// Skip already registed structs
		_, found := rulesRegistry[structType]
		if found {
			continue
		}

		var rules []*Rule

		for i := range structType.NumField() {
			field := structType.Field(i)

			// Queue the nested structs
			if field.Type.Kind() == reflect.Struct {
				structTypes = append(structTypes, field.Type)
				continue
			}

			config := field.Tag.Get("checkers")
			if config == "" {
				continue
			}

			checkers, err := initCheckersFromConfig(config)
			if err != nil {
				return fmt.Errorf("field %s has errors: %w", field.Name, err)
			}

			rules = append(rules, &Rule{
				FieldIndex: field.Index,
				Checkers:   checkers,
			})
		}

		rulesRegistry[structType] = rules
	}

	return nil
}

// initCheckersFromConfig parses the given config string and initializes checker instances.
func initCheckersFromConfig(config string) ([]CheckFunc, error) {
	items := strings.Fields(config)
	checkers := make([]CheckFunc, len(items))

	for i, item := range items {
		checker, ok := checkersCache[item]
		if ok {
			checkers[i] = checker
			continue
		}

		name, params, _ := strings.Cut(item, ":")

		maker, ok := makers[name]
		if !ok {
			return nil, fmt.Errorf("checker %s not found", name)
		}

		checkers[i] = maker(params)
		checkersCache[item] = checkers[i]
	}

	return checkers, nil
}
