// 第一个 Go
package main

import (
	"fmt"
	"github.com/henry-insomniac/Go-practice/practice"
	"github.com/henry-insomniac/Go-practice/updateName"
)

func main() {

	p := updateName.Person{
		Name: "test",
	}
	fmt.Println("oldName ->", p)
	newPerson := updateName.SetName(&p, "test1")
	fmt.Println("newName ->", newPerson)

	bookLib := make(map[string]*practice.Book)

	book1 := &practice.Book{
		BookName: "相对论",
		Author:   "爱因斯坦",
	}

	book2 := &practice.Book{
		BookName: "认知觉醒",
		Author:   "周岭",
	}

	book3 := &practice.Book{
		BookName: "认知觉醒",
		Author:   "周岭",
	}
	practice.SaveBookInLib(book1, bookLib)
	practice.SaveBookInLib(book2, bookLib)
	practice.SaveBookInLib(book3, bookLib)

	jsonStr := practice.ListBooks(bookLib)
	fmt.Println(jsonStr)
}
