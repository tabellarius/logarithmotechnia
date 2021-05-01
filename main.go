package main

import "fmt"

func main() {
	map1 := map[string]string{"one": "1", "two": "2"}
	map2 := map1
	fmt.Println(map1, map2)

	map1["one"] = "100"
	fmt.Println(map1, map2)

}
