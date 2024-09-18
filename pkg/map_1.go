package pkg

import (
	"fmt"
)


func MapTest() {
	m1 := make(map[string]string)
	m1["hello"] = "world"
	m1["foo"] = "bar"
	PrintMap(m1)
	fmt.Println(m1["213"])
}

func PrintMap(m map[string]string) int {
	fmt.Printf("map: %v, len: %d\n", m, len(m))
	return len(m)
}