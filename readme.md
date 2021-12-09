# gorm-where-builder

help you build where condition

## usage

install:

```sh
go get github.com/3vilive/gorm-where-builder
```

build some condition sql:

```go
builder := builder.NewBuilder(builder.Model{
    "age": builder.Eq(1),
})
where, args := builder.Where()
fmt.Printf("where: %s\nargs: %#v\n", where, args)
// where: `age` = ?
// args: []interface {}{1}

// you can easy use Where() with gorm like this
// db, _ := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
// db.Table("user").Where(builder.Where()).Rows()
```

build some condition sql from struct:

```go
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
b = builder.NewBuilder(filter)
where, args = b.Where()
fmt.Printf("where: %s\nargs: %#v\n", where, args)
// where: `name` = ? and `age` != ? and `category` like ? and `group` in ? and `wallet` is null and `bottle` is not null
// args: []interface {}{"3vilive", 18, "%Go%", []string{"User Growth", "Sales"}}
```

## tag 

use tag to describe struct behavior: `where:"field_name,condition_type"`

| condition_type | sql | 
| :-: | :-: |
| eq | = |
| ne | != |
| extactlike | like 'foo' |
| prefixlike | like 'foo%' |
| suffixlike | like '%foo' |
| containlike, like | '%foo%' |
| in | in |
| isnull | is null |
| isnotnull | is not null |




