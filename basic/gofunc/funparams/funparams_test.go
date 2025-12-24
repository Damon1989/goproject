package funparams

import (
	"fmt"
	"log"
	"testing"
)

func TestMethodMain(t *testing.T) {
	methodMain()
}

func TestDoMain(t *testing.T) {
	doMain()
}

func TestFn1(t *testing.T) {
	fn1(10)
}

func TestFn2(t *testing.T) {
	log.Println(fn2(100))
}

func TestFn3(t *testing.T) {
	log.Println(fn3(5))
}

func TestName(t *testing.T) {
	log.Println(0)
	defer log.Println(1)
	defer log.Println(2)
	defer log.Println(3)
	i := 0
	j := 100
	r := j / i
	fmt.Println(r)
	log.Println(4)

}
