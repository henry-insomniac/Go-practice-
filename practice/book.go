/**
 * @Author: mm
 * @Date: 2025/4/26 13:02
 * @Description: book.go
 */
package practice

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	BookName string
	Author   string
}

func NewBook(author string, bookName string) *Book {
	return &Book{
		BookName: bookName,
		Author:   author,
	}
}

func (b *Book) GetName() string {
	return b.BookName
}

func (b *Book) GetAuthor() string {
	return b.Author
}

func ChangeBookInfo(b *Book, newBookName string, newAuthor string) {
	b.BookName = newBookName
	b.Author = newAuthor
}

func SaveBookInLib(b *Book, lib map[string]*Book) map[string]*Book {
	if _, exist := lib[b.BookName]; exist {
		fmt.Println("已有同名书籍！无法存储:", b.BookName)
		return nil
	}

	lib[b.BookName] = b
	fmt.Println("保存成功")
	return lib
}

func ListBooks(lib map[string]*Book) string {
	// 1. 把 map 转成切片
	var books []*Book
	for _, book := range lib {
		books = append(books, book)
	}

	// 2. 序列化成 JSON
	jsonData, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		fmt.Println("序列化失败:", err)
		return ""
	}

	// 3. 返回 JSON 字符串
	return string(jsonData)
}
