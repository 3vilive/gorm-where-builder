package builder

import "fmt"

type ConditionType int

const (
	ConditionTypeUnknown     ConditionType = 0  //
	ConditionTypeEq          ConditionType = 10 // =
	ConditionTypeNe          ConditionType = 11 // !=
	ConditionTypeGt          ConditionType = 12 // >
	ConditionTypeGte         ConditionType = 13 // >=
	ConditionTypeLt          ConditionType = 14 // <
	ConditionTypeLte         ConditionType = 15 // <=
	ConditionTypeExtactLike  ConditionType = 20 // like '?'
	ConditionTypePrefixLike  ConditionType = 21 // like '?%'
	ConditionTypeSuffixLike  ConditionType = 22 // like '%?'
	ConditionTypeContainLike ConditionType = 23 // like '%?%'
	ConditionTypeIn          ConditionType = 30 // in
	ConditionTypeIsNull      ConditionType = 40 // is null
	ConditionTypeIsNotNull   ConditionType = 41 // is not null
)

type Condition struct {
	Type ConditionType
	Val  interface{}
}

func (c *Condition) BuildConditionSQL(field string) string {
	switch c.Type {
	case ConditionTypeEq:
		return fmt.Sprintf("`%s` = ?", field)
	case ConditionTypeNe:
		return fmt.Sprintf("`%s` != ?", field)
	case ConditionTypeGt:
		return fmt.Sprintf("`%s` > ?", field)
	case ConditionTypeGte:
		return fmt.Sprintf("`%s` >= ?", field)
	case ConditionTypeLt:
		return fmt.Sprintf("`%s` < ?", field)
	case ConditionTypeLte:
		return fmt.Sprintf("`%s` <= ?", field)
	case ConditionTypeExtactLike, ConditionTypePrefixLike, ConditionTypeSuffixLike, ConditionTypeContainLike:
		return fmt.Sprintf("`%s` like ?", field)
	case ConditionTypeIn:
		return fmt.Sprintf("`%s` in ?", field)
	case ConditionTypeIsNull:
		return fmt.Sprintf("`%s` is null", field)
	case ConditionTypeIsNotNull:
		return fmt.Sprintf("`%s` is not null", field)
	}

	return ""
}

func buildConditionTypeFn(t ConditionType) func(interface{}) Condition {
	return func(i interface{}) Condition {
		return Condition{
			Type: t,
			Val:  i,
		}
	}
}

var Eq = buildConditionTypeFn(ConditionTypeEq)
var Ne = buildConditionTypeFn(ConditionTypeNe)
var Gt = buildConditionTypeFn(ConditionTypeGt)
var Gte = buildConditionTypeFn(ConditionTypeGte)
var Lt = buildConditionTypeFn(ConditionTypeLt)
var Lte = buildConditionTypeFn(ConditionTypeLte)
var ExtactLike = buildConditionTypeFn(ConditionTypeExtactLike)
var PrefixLike = buildConditionTypeFn(ConditionTypePrefixLike)
var SuffixLike = buildConditionTypeFn(ConditionTypeSuffixLike)
var ContainLike = buildConditionTypeFn(ConditionTypeContainLike)
var In = buildConditionTypeFn(ConditionTypeIn)
var IsNull = buildConditionTypeFn(ConditionTypeIsNull)(nil)
var IsNotNull = buildConditionTypeFn(ConditionTypeIsNotNull)(nil)
