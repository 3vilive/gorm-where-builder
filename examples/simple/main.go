package main

import (
	"fmt"

	builder "github.com/3vilive/gorm-where-builder"
)

func main() {
	b := builder.NewBuilder(builder.Model{
		"age": []builder.Condition{builder.Eq(1)},
	})
	where, args := b.Where()
	fmt.Printf("where: %s\nargs: %#v\n", where, args)

	// db, _ := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	// db.Table("user").Where(builder.Where()).Rows()

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
}
