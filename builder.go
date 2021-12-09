package builder

import (
	"fmt"
	"strings"
)

type Model map[string]Condition

type Builder struct {
	model Model
}

func NewBuilderFromModel(model Model) *Builder {
	return &Builder{
		model: model,
	}
}

func NewBuilderFromStruct(s interface{}) *Builder {
	return NewBuilderFromModel(NewModelFromStruct(s))
}

func NewBuilder(any interface{}) *Builder {
	switch i := any.(type) {
	case Model:
		return NewBuilderFromModel(i)
	case *Model:
		return NewBuilderFromModel(*i)
	}

	return NewBuilderFromStruct(any)
}

func (b *Builder) Build() []interface{} {
	if b == nil || len(b.model) == 0 {
		return make([]interface{}, 0)
	}

	where := make([]interface{}, 0, len(b.model)+1)
	where = append(where, "")
	conditions := make([]string, 0, len(b.model))
	for field, condition := range b.model {
		conditions = append(conditions, condition.BuildConditionSQL(field))

		// 过滤不需要 Val 的类型
		switch condition.Type {
		case ConditionTypeIsNull, ConditionTypeIsNotNull:
			continue

		case ConditionTypePrefixLike:
			where = append(where, fmt.Sprintf("%s%%", condition.Val))
		case ConditionTypeSuffixLike:
			where = append(where, fmt.Sprintf("%%%s", condition.Val))
		case ConditionTypeContainLike:
			where = append(where, fmt.Sprintf("%%%s%%", condition.Val))

		default:
			where = append(where, condition.Val)
		}

	}
	where[0] = strings.Join(conditions, " and ")

	return where
}

func (b *Builder) Where() (string, []interface{}) {
	whereWithArgs := b.Build()
	if len(whereWithArgs) == 0 {
		return "", make([]interface{}, 0)
	}
	if len(whereWithArgs) == 1 {
		return whereWithArgs[0].(string), make([]interface{}, 0)
	}

	return whereWithArgs[0].(string), whereWithArgs[1:]
}
