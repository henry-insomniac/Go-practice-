/**
 * @Author: mm
 * @Date: 2025/4/26 12:12
 * @Description: setName
 */
package updateName

type Person struct {
	Name string
}

// SetName 指针
func SetName(person *Person, newName string) Person {
	person.Name = newName
	return Person{Name: newName}
}
