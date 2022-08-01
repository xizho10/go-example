// go run args.go -max=100
package chapter13

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func Chapter13main() {
	rand.Seed(time.Now().UnixNano()) //获得随机结果
	// define flags
	maxp := flag.Int("max", 6, "the max value")
	// parse
	flag.Parse()
	// generate a number between 0 and max
	fmt.Println("maxp: ", *maxp)
	fmt.Println(rand.Intn(*maxp))
	// Any additional non-flag arguments can be retrieved with flag.Args() which returns a []string.
	//fmt.Println(flag.Args())
}
