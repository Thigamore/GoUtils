package main

import (
	"fmt"

	dt "github.com/Thigamore/GoUtils/datatypes"
)

func main() {
	list := dt.CreateListFromArr(&[]string{"I", "am", "cool"})
	fmt.Println(list)
}
