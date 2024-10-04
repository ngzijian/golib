package main

import (
	"fmt"
	"github.com/ngzijian/golib/collection"
)

func main() {
	arr := []string{"1", "2", "1", "2", "3"}
	list := collection.List{
		Val: arr,
	}
	fmt.Println(list.Val)
	list.Distinct()
	fmt.Println(list.Val)
}
