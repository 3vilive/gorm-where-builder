package builder

import (
	"reflect"
	"strings"
)

const (
	_TagEq               = "eq"
	_TagNe               = "ne"
	_TagExtactLike       = "extactlike"
	_TagPrefixLike       = "prefixlike"
	_TagSuffixLike       = "suffixlike"
	_TagContainLike      = "containlike"
	_TagContainLikeShort = "like"
	_TagIn               = "in"
	_TagIsNull           = "isnull"
	_TagIsNotNull        = "isnotnull"
)

func NewModelFromStruct(s interface{}) Model {
	model := make(Model)

	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return model
	}

	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		fv := v.Field(i)
		if fv.IsNil() {
			continue
		}

		ft := t.Field(i)
		filterTag := ft.Tag.Get("where")
		if filterTag == "" {
			continue
		}

		tags := strings.Split(filterTag, ",")
		if len(tags) != 2 {
			continue
		}

		field, conditionTag := tags[0], tags[1]
		switch conditionTag {
		case _TagEq:
			model[field] = Eq(fv.Elem().Interface())
		case _TagNe:
			model[field] = Ne(fv.Elem().Interface())
		case _TagContainLike, _TagContainLikeShort:
			model[field] = ContainLike(fv.Elem().Interface())
		case _TagPrefixLike:
			model[field] = PrefixLike(fv.Elem().Interface())
		case _TagSuffixLike:
			model[field] = SuffixLike(fv.Elem().Interface())
		case _TagExtactLike:
			model[field] = ExtactLike(fv.Elem().Interface())
		case _TagIn:
			if fv.Kind() == reflect.Slice || fv.Kind() == reflect.Array {
				model[field] = In(fv.Interface())
			}
		case _TagIsNull:
			model[field] = IsNull
		case _TagIsNotNull:
			model[field] = IsNotNull
		}
	}

	return model
}
