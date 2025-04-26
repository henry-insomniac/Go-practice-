// 第一个 Go
package main

import (
	"fmt"
	"github.com/henry-insomniac/Go-practice/updateName"
)

func main() {

	p := updateName.Person{
		Name: "test",
	}
	fmt.Println("oldName ->", p)
	newPerson := updateName.SetName(&p, "test1")
	fmt.Println("newName ->", newPerson)
}
