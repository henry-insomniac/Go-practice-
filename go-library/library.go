/**
 * @Author: mm
 * @Date: 2025/4/26 19:59
 * @Description: library
 */
package go_library

import (
	"encoding/json"
	"fmt"
)

// BookLib - 定义全局图书馆库
var BookLib = make(map[string]*Book)

// CreateBook - 添加一本书
func CreateBook(book *Book) *Book {
	if _, exist := BookLib[book.Title]; exist {
		return fmt.Errorf("书籍 '%s' 已存在", book.Title)
	}
	BookLib[book.Title] = book
	return nil
}

// Read - 查看图书馆所有书
func Read() (string, error) {
	var books []*Book
	for _, book := range BookLib {
		books = append(books, book)
	}
	data, err := json.Marshal(books)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Update - 修改书籍信息
func Update(oldTitle string, newBook *Book) error {
	if _, exist := BookLib[newBook.Title]; !exist {
		fmt.Errorf("要修改的书 '%s' 不存在", oldTitle)
		return nil
	}
	delete(BookLib, oldTitle)
	BookLib[newBook.Title] = newBook
	return nil
}

func Delete(title string) error {
	if _, exist := BookLib[title]; !exist {
		fmt.Errorf("要修改的书 '%s' 不存在", title)
	}
	delete(BookLib, title)
	fmt.Println("删除成功")
	return nil
}
