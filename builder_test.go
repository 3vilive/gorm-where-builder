package builder

import (
	"fmt"
	"testing"
)

func TestConditions(t *testing.T) {
	conditionValidationMap := map[string]string{
		"name":     "`name` = ?",
		"age":      "`age` != ?",
		"category": "`category` like ?",
		"group":    "`group` in ?",
		"wallet":   "`wallet` is null",
		"bottle":   "`bottle` is not null",
	}
	m := Model{
		"name":     Eq("3vilive"),
		"age":      Ne(1),
		"category": ExtactLike("Go"),
		"group":    In([]string{"User Growth", "Sales"}),
		"wallet":   IsNull,
		"bottle":   IsNotNull,
	}

	for field, condition := range m {
		expect := conditionValidationMap[field]
		got := condition.BuildConditionSQL(field)

		if expect != got {
			t.Errorf("expect `%s` but got `%s` when test case `%s`", expect, got, field)
		}
	}
}

func TestExtactLike(t *testing.T) {
	builder := NewBuilder(Model{
		"category": ExtactLike("Go"),
	})
	got := builder.Build()[1].(string)
	expect := "Go"
	if got != expect {
		t.Errorf("expect `%s` but got `%s`", expect, got)
	}
}

func TestPrefixLike(t *testing.T) {
	builder := NewBuilder(Model{
		"category": PrefixLike("Go"),
	})
	got := builder.Build()[1].(string)
	expect := "Go%"
	if got != expect {
		t.Errorf("expect `%s` but got `%s`", expect, got)
	}
}

func TestSuffixLike(t *testing.T) {
	builder := NewBuilder(Model{
		"category": SuffixLike("Go"),
	})
	got := builder.Build()[1].(string)
	expect := "%Go"
	if got != expect {
		t.Errorf("expect `%s` but got `%s`", expect, got)
	}
}
func TestContainLike(t *testing.T) {
	builder := NewBuilder(Model{
		"category": ContainLike("Go"),
	})
	got := builder.Build()[1].(string)
	expect := "%Go%"
	if got != expect {
		t.Errorf("expect `%s` but got `%s`", expect, got)
	}
}

func TestFromStruct(t *testing.T) {
	type UserFilter struct {
		Name     *string     `where:"name,eq"`
		Age      *int        `where:"age,ne"`
		Category *string     `where:"category,like"`
		Group    []string    `where:"group,in"`
		Wallet   interface{} `where:"wallet,isnull"`
		Bottle   interface{} `where:"bottle,isnotnull"`
	}

	name := "3vilive"
	age := 18
	category := "Go"
	groups := []string{"User Growth", "Sales"}
	wallet := struct{}{}
	bottle := struct{}{}

	filter := UserFilter{
		Name:     &name,
		Age:      &age,
		Category: &category,
		Group:    groups,
		Wallet:   wallet,
		Bottle:   bottle,
	}

	validationModel := Model{
		"name":     Eq("3vilive"),
		"age":      Ne(1),
		"category": ContainLike("Go"),
		"group":    In([]string{"User Growth", "Sales"}),
		"wallet":   IsNull,
		"bottle":   IsNotNull,
	}

	m := NewModelFromStruct(filter)
	fmt.Printf("m: %#v\n", m)

	for field, condition := range validationModel {
		checkCond, ok := m[field]
		if !ok {
			t.Errorf("validate field `%s` error: condition not found", field)
			continue
		}

		if condition.Type != checkCond.Type {
			t.Errorf("validate field `%s` error: expect condition type %d but got %d", field, condition.Type, checkCond.Type)
		}
	}

	builder := NewBuilder(m)
	whereWithArgs := builder.Build()
	where, args := whereWithArgs[0], whereWithArgs[1:]
	fmt.Printf("%#v", where)
	fmt.Printf("%#v", args)
}
