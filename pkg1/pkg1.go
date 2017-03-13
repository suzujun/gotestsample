package pkg1

import "fmt"

func Sum(a int, b int) int {
	c := a + b
	fmt.Println(a, b, c)
	return c
}
func GetFuga() string {
	return "fuga"
}
