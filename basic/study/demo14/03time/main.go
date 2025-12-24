package main

import (
	"fmt"
	"time"
)

func main() {
	var unixTime = 1740045904
	timeObj := time.Unix(int64(unixTime), 0)
	fmt.Println(timeObj.Format("2006-01-02 15:04:05"))

}
