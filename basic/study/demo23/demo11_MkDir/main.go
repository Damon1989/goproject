package main

import (
	"fmt"
	"os"
)

func main() {
	//err := os.Mkdir("./go", 0666)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	err := os.MkdirAll("./dir1/dir2/dir3", 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

}
