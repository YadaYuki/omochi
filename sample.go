package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	a := new([]int64)
	json.Unmarshal([]byte("{1,3,4,5}"), a)
	fmt.Println(*a)

}
