package pack

import (
	"fmt"
	"github.com/chensheng0/repoA/item"
)

func PackItem(i item.Item) {
	fmt.Println("name ", i.Name)
	fmt.Println("age ", i.Age)
	fmt.Println("gender ", i.Gender)
	fmt.Println("class ", i.Class)
	fmt.Println("end")
}
