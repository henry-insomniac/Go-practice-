/**
 * @Author: mm
 * @Date: 2025/4/26 19:59
 * @Description: book
 */
package go_library

import "fmt"

type Book struct {
	// 手写字母大写对外可见
	Title  string
	Author string
}

func (b *Book) Introduce() {
	fmt.Println("Book Introduce")
	
}
