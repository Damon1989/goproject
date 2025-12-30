package main

import "github.com/damon/gogofly/cmd"

func main() {
	defer cmd.Clean()
	cmd.Start()
}
