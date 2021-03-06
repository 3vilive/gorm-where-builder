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
		"money":    "`money` > ?",
		"money2":   "`money2` >= ?",
		"money3":   "`money3` < ?",
		"money4":   "`money4` <= ?",
	}
	m := Model{
		"name":     []Condition{Eq("3vilive")},
		"age":      []Condition{Ne(1)},
		"category": []Condition{ExtactLike("Go")},
		"group":    []Condition{In([]string{"User Growth", "Sales"})},
		"wallet":   []Condition{IsNull},
		"bottle":   []Condition{IsNotNull},
		"money":    []Condition{Gt(1)},
		"money2":   []Condition{Gte(1)},
		"money3":   []Condition{Lt(1)},
		"money4":   []Condition{Lte(1)},
	}

	for field, fieldCondisions := range m {
		expect := conditionValidationMap[field]
		for _, condition := range fieldCondisions {
			got := condition.BuildConditionSQL(field)
			if expect != got {
				t.Errorf("expect `%s` but got `%s` when test case `%s`", expect, got, field)
			}
		}

	}
}

func TestExtactLike(t *testing.T) {
	builder := NewBuilder(Model{
		"category": []Condition{ExtactLike("Go")},
	})
	got := builder.Build()[1].(string)
	expect := "Go"
	if got != expect {
		t.Errorf("expect `%s` but got `%s`", expect, got)
	}
}

func TestPrefixLike(t *testing.T) {
	builder := NewBuilder(Model{
		"category": []Condition{PrefixLike("Go")},
	})
	got := builder.Build()[1].(string)
	expect := "Go%"
	if got != expect {
		t.Errorf("expect `%s` but got `%s`", expect, got)
	}
}

func TestSuffixLike(t *testing.T) {
	builder := NewBuilder(Model{
		"category": []Condition{SuffixLike("Go")},
	})
	got := builder.Build()[1].(string)
	expect := "%Go"
	if got != expect {
		t.Errorf("expect `%s` but got `%s`", expect, got)
	}
}
func TestContainLike(t *testing.T) {
	builder := NewBuilder(Model{
		"category": []Condition{ContainLike("Go")},
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
		"name":     []Condition{Eq("3vilive")},
		"age":      []Condition{Ne(1)},
		"category": []Condition{ContainLike("Go")},
		"group":    []Condition{In([]string{"User Growth", "Sales"})},
		"wallet":   []Condition{IsNull},
		"bottle":   []Condition{IsNotNull},
	}

	m := NewModelFromStruct(filter)
	fmt.Printf("m: %#v\n", m)

	for field, conditions := range validationModel {
		checkConds, ok := m[field]
		if !ok {
			t.Errorf("validate field `%s` error: condition not found", field)
			continue
		}

		for i, condition := range conditions {
			if condition.Type != checkConds[i].Type {
				t.Errorf("validate field `%s` error: expect condition type %d but got %d", field, condition.Type, checkConds[i].Type)
			}
		}

	}

	builder := NewBuilder(m)
	whereWithArgs := builder.Build()
	where, args := whereWithArgs[0], whereWithArgs[1:]
	fmt.Printf("%#v", where)
	fmt.Printf("%#v", args)
}
