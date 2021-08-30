package main

import (
	"fmt"

	"github.com/Prithvipal/proto-buffer/src/simple/simplepb"
)

func main() {
	dosimlpe()
}

func dosimlpe() {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SimpleList: []int32{1, 4, 7, 9},
	}

	fmt.Printf("%+v\n", sm)
	sm.Name = "I renamed you"
	fmt.Printf("%+v\n", sm)
}
